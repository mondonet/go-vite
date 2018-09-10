package worker

import (
	"container/heap"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/log15"
	"github.com/vitelabs/go-vite/unconfirmed/model"
	"sync"
)
type RightEvent struct {
	Gid            *types.Gid
	Address        *types.Address
	StartTs        uint64
	EndTs          uint64
	SnapshotHash   *types.Hash
	SnapshotHeight int
}

type ContractWorker struct {
	uAccess *model.UAccess

	gid                 *types.Gid
	addresses           *types.Address
	contractAddressList []*types.Address

	status                 int
	dispatcherSleep        bool
	dispatcherAlarm        chan struct{}
	breaker                chan struct{}
	stopDispatcherListener chan struct{}

	contractTasks   []*ContractTask
	priorityToQueue *model.PriorityToQueue
	blackList       map[string]bool // map[(toAddress+fromAddress).String]

	statusMutex sync.Mutex

	log log15.Logger
}

func NewContractWorker(uAccess *model.UAccess, gid *types.Gid, address *types.Address, addressList []*types.Address) *ContractWorker {
	return &ContractWorker{
		uAccess:                uAccess,
		gid:                    gid,
		addresses:              address,
		contractAddressList:    addressList,
		status:                 Create,
		dispatcherSleep:        false,
		dispatcherAlarm:        make(chan struct{}, 1),
		breaker:                make(chan struct{}, 1),
		stopDispatcherListener: make(chan struct{}, 1),
		contractTasks:          make([]*ContractTask, CONTRACT_TASK_SIZE),
		blackList:              make(map[string]bool),
		log:                    log15.New("ContractWorker addr", address.String(), "gid", gid),
	}
}

func (w *ContractWorker) Start(event *RightEvent) {
	w.log.Info("worker startWork is called")
	w.statusMutex.Lock()
	defer w.statusMutex.Unlock()

	w.uAccess.AddContractLis(w.gid, func() {
		w.NewUnconfirmedTxAlarm()
	})

	for _, v := range w.contractTasks {
		v.InitContractTask(w.uAccess, event)
		go v.Start(&w.blackList)
	}
	go w.DispatchTask()

	w.status = Start
}

func (w *ContractWorker) Stop() {
	w.statusMutex.Lock()
	defer w.statusMutex.Unlock()
	if w.status != Stop {

		w.breaker <- struct{}{}
		// todo: to clear tomap

		w.uAccess.RemoveContractLis(w.gid)
		w.dispatcherSleep = true
		close(w.dispatcherAlarm)

		<-w.stopDispatcherListener
		close(w.stopDispatcherListener)

		// todo 2. Stop all task
		for _, v := range w.contractTasks {
			v.Stop()
		}
		w.status = Stop
	}
}

func (w *ContractWorker) Close() error {
	w.Stop()
	return nil
}

func (w ContractWorker) Status() int {
	w.statusMutex.Lock()
	defer w.statusMutex.Unlock()
	return w.status
}

func (w *ContractWorker) NewUnconfirmedTxAlarm() {
	if w.dispatcherSleep {
		w.dispatcherAlarm <- struct{}{}
	}
}

func (w *ContractWorker) DispatchTask() {
	//todo add mutex
	w.FetchNew()
	for {
		for i := 0; i < w.priorityToQueue.Len(); i++ {
			tItem := heap.Pop(w.priorityToQueue).(*model.ToItem)
			priorityFromQueue := tItem.Value
			for j := 0; j < priorityFromQueue.Len(); j++ {
			FINDFREETASK:
				if w.Status() == Stop {
					// clear blackList
					w.blackList = nil
					// fixme: to clear priorityToQueue?
					goto END
				}

				freeTaskIndex := w.FindAFreeTask()
				if freeTaskIndex == -1 {
					goto FINDFREETASK
				}

				fItem := heap.Pop(priorityFromQueue).(*model.FromItem)
				w.contractTasks[freeTaskIndex].subQueue <- fItem
			}
		}

		w.dispatcherSleep = true

		select {
		case <-w.breaker:
			goto END
		case <-w.dispatcherAlarm:
			w.dispatcherSleep = false
			w.FetchNew()
		}
	}
END:
	w.log.Info("ContractWorker send stopDispatcherListener")
	w.stopDispatcherListener <- struct{}{}
	w.log.Info("ContractWorker DispatchTask end")
}

func (w *ContractWorker) FindAFreeTask() (index int) {
	for k, v := range w.contractTasks {
		if v.status == Idle {
			return k
		}
	}
	return -1
}

func (w *ContractWorker) FetchNew() {
	for i := 0; i < len(w.contractAddressList); i++ {
		blockList, err := w.uAccess.GetUnconfirmedBlocks(0, 1, CONTRACT_FETCH_SIZE, w.contractAddressList[i])
		if err != nil {
			w.log.Error("ContractWorker.FetchNew.GetUnconfirmedBlocks", "error", err)
			continue
		}
		for _, v := range blockList {
			// when a to-from pair  was added into blackList,
			// the other block which under the same to-from pair won't fetch any more during the same block-out period
			var blKey = (*v).ToAddress.String() + (*v).AccountAddress.String()
			if _, ok := w.blackList[blKey]; !ok {
				w.priorityToQueue.InsertNew(v)
			}
		}
	}
}
