// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vitelabs/go-vite/consensus (interfaces: DposReader)

// Package consensus is a generated GoMock package.
package consensus

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	types "github.com/vitelabs/go-vite/common/types"
	core "github.com/vitelabs/go-vite/consensus/core"
)

// MockDposReader is a mock of DposReader interface
type MockDposReader struct {
	ctrl     *gomock.Controller
	recorder *MockDposReaderMockRecorder
}

// MockDposReaderMockRecorder is the mock recorder for MockDposReader
type MockDposReaderMockRecorder struct {
	mock *MockDposReader
}

// NewMockDposReader creates a new mock instance
func NewMockDposReader(ctrl *gomock.Controller) *MockDposReader {
	mock := &MockDposReader{ctrl: ctrl}
	mock.recorder = &MockDposReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDposReader) EXPECT() *MockDposReaderMockRecorder {
	return m.recorder
}

// ElectionIndex mocks base method
func (m *MockDposReader) ElectionIndex(arg0 uint64) (*electionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ElectionIndex", arg0)
	ret0, _ := ret[0].(*electionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ElectionIndex indicates an expected call of ElectionIndex
func (mr *MockDposReaderMockRecorder) ElectionIndex(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ElectionIndex", reflect.TypeOf((*MockDposReader)(nil).ElectionIndex), arg0)
}

// GenProofTime mocks base method
func (m *MockDposReader) GenVoteTime(arg0 uint64) time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenProofTime", arg0)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GenProofTime indicates an expected call of GenProofTime
func (mr *MockDposReaderMockRecorder) GenVoteTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenProofTime", reflect.TypeOf((*MockDposReader)(nil).GenVoteTime), arg0)
}

// GetInfo mocks base method
func (m *MockDposReader) GetInfo() *core.GroupInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo")
	ret0, _ := ret[0].(*core.GroupInfo)
	return ret0
}

// GetInfo indicates an expected call of GetInfo
func (mr *MockDposReaderMockRecorder) GetInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockDposReader)(nil).GetInfo))
}

// Index2Time mocks base method
func (m *MockDposReader) Index2Time(arg0 uint64) (time.Time, time.Time) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Index2Time", arg0)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(time.Time)
	return ret0, ret1
}

// Index2Time indicates an expected call of Index2Time
func (mr *MockDposReaderMockRecorder) Index2Time(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index2Time", reflect.TypeOf((*MockDposReader)(nil).Index2Time), arg0)
}

// Time2Index mocks base method
func (m *MockDposReader) Time2Index(arg0 time.Time) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Time2Index", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Time2Index indicates an expected call of Time2Index
func (mr *MockDposReaderMockRecorder) Time2Index(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Time2Index", reflect.TypeOf((*MockDposReader)(nil).Time2Index), arg0)
}

// VerifyProducer mocks base method
func (m *MockDposReader) VerifyProducer(arg0 types.Address, arg1 time.Time) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyProducer", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyProducer indicates an expected call of VerifyProducer
func (mr *MockDposReaderMockRecorder) VerifyProducer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyProducer", reflect.TypeOf((*MockDposReader)(nil).VerifyProducer), arg0, arg1)
}
