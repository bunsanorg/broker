// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/bunsanorg/broker/go/bunsan/broker/worker/driver/driverutil (interfaces: StdoutRunner)

package mock_driverutil

import (
	driverutil "github.com/bunsanorg/broker/go/bunsan/broker/worker/driver/driverutil"
	gomock "github.com/golang/mock/gomock"
	exec "os/exec"
)

// Mock of StdoutRunner interface
type MockStdoutRunner struct {
	ctrl     *gomock.Controller
	recorder *_MockStdoutRunnerRecorder
}

// Recorder for MockStdoutRunner (not exported)
type _MockStdoutRunnerRecorder struct {
	mock *MockStdoutRunner
}

func NewMockStdoutRunner(ctrl *gomock.Controller) *MockStdoutRunner {
	mock := &MockStdoutRunner{ctrl: ctrl}
	mock.recorder = &_MockStdoutRunnerRecorder{mock}
	return mock
}

func (_m *MockStdoutRunner) EXPECT() *_MockStdoutRunnerRecorder {
	return _m.recorder
}

func (_m *MockStdoutRunner) Run(_param0 *exec.Cmd, _param1 driverutil.StdoutParser) error {
	ret := _m.ctrl.Call(_m, "Run", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStdoutRunnerRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Run", arg0, arg1)
}
