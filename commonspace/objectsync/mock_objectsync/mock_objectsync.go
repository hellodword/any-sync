// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/any-sync/commonspace/objectsync (interfaces: ObjectSync)
//
// Generated by this command:
//
//	mockgen -destination mock_objectsync/mock_objectsync.go github.com/anyproto/any-sync/commonspace/objectsync ObjectSync
//
// Package mock_objectsync is a generated GoMock package.
package mock_objectsync

import (
	context "context"
	reflect "reflect"
	time "time"

	app "github.com/anyproto/any-sync/app"
	objectsync "github.com/anyproto/any-sync/commonspace/objectsync"
	spacesyncproto "github.com/anyproto/any-sync/commonspace/spacesyncproto"
	gomock "go.uber.org/mock/gomock"
)

// MockObjectSync is a mock of ObjectSync interface.
type MockObjectSync struct {
	ctrl     *gomock.Controller
	recorder *MockObjectSyncMockRecorder
}

// MockObjectSyncMockRecorder is the mock recorder for MockObjectSync.
type MockObjectSyncMockRecorder struct {
	mock *MockObjectSync
}

// NewMockObjectSync creates a new mock instance.
func NewMockObjectSync(ctrl *gomock.Controller) *MockObjectSync {
	mock := &MockObjectSync{ctrl: ctrl}
	mock.recorder = &MockObjectSyncMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectSync) EXPECT() *MockObjectSyncMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockObjectSync) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockObjectSyncMockRecorder) Close(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockObjectSync)(nil).Close), arg0)
}

// CloseThread mocks base method.
func (m *MockObjectSync) CloseThread(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseThread", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseThread indicates an expected call of CloseThread.
func (mr *MockObjectSyncMockRecorder) CloseThread(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseThread", reflect.TypeOf((*MockObjectSync)(nil).CloseThread), arg0)
}

// HandleMessage mocks base method.
func (m *MockObjectSync) HandleMessage(arg0 context.Context, arg1 objectsync.HandleMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleMessage indicates an expected call of HandleMessage.
func (mr *MockObjectSyncMockRecorder) HandleMessage(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMessage", reflect.TypeOf((*MockObjectSync)(nil).HandleMessage), arg0, arg1)
}

// HandleRequest mocks base method.
func (m *MockObjectSync) HandleRequest(arg0 context.Context, arg1 *spacesyncproto.ObjectSyncMessage) (*spacesyncproto.ObjectSyncMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleRequest", arg0, arg1)
	ret0, _ := ret[0].(*spacesyncproto.ObjectSyncMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleRequest indicates an expected call of HandleRequest.
func (mr *MockObjectSyncMockRecorder) HandleRequest(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleRequest", reflect.TypeOf((*MockObjectSync)(nil).HandleRequest), arg0, arg1)
}

// Init mocks base method.
func (m *MockObjectSync) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockObjectSyncMockRecorder) Init(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockObjectSync)(nil).Init), arg0)
}

// LastUsage mocks base method.
func (m *MockObjectSync) LastUsage() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastUsage")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// LastUsage indicates an expected call of LastUsage.
func (mr *MockObjectSyncMockRecorder) LastUsage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastUsage", reflect.TypeOf((*MockObjectSync)(nil).LastUsage))
}

// Name mocks base method.
func (m *MockObjectSync) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockObjectSyncMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockObjectSync)(nil).Name))
}

// Run mocks base method.
func (m *MockObjectSync) Run(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockObjectSyncMockRecorder) Run(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockObjectSync)(nil).Run), arg0)
}
