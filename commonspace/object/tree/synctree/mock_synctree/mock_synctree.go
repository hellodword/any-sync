// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anytypeio/any-sync/commonspace/object/tree/synctree (interfaces: SyncClient,SyncTree,ReceiveQueue,HeadNotifiable)

// Package mock_synctree is a generated GoMock package.
package mock_synctree

import (
	context "context"
	reflect "reflect"

	objecttree "github.com/anytypeio/any-sync/commonspace/object/tree/objecttree"
	treechangeproto "github.com/anytypeio/any-sync/commonspace/object/tree/treechangeproto"
	treestorage "github.com/anytypeio/any-sync/commonspace/object/tree/treestorage"
	spacesyncproto "github.com/anytypeio/any-sync/commonspace/spacesyncproto"
	gomock "github.com/golang/mock/gomock"
)

// MockSyncClient is a mock of SyncClient interface.
type MockSyncClient struct {
	ctrl     *gomock.Controller
	recorder *MockSyncClientMockRecorder
}

// MockSyncClientMockRecorder is the mock recorder for MockSyncClient.
type MockSyncClientMockRecorder struct {
	mock *MockSyncClient
}

// NewMockSyncClient creates a new mock instance.
func NewMockSyncClient(ctrl *gomock.Controller) *MockSyncClient {
	mock := &MockSyncClient{ctrl: ctrl}
	mock.recorder = &MockSyncClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSyncClient) EXPECT() *MockSyncClientMockRecorder {
	return m.recorder
}

// BroadcastAsync mocks base method.
func (m *MockSyncClient) BroadcastAsync(arg0 *treechangeproto.TreeSyncMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BroadcastAsync", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// BroadcastAsync indicates an expected call of BroadcastAsync.
func (mr *MockSyncClientMockRecorder) BroadcastAsync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastAsync", reflect.TypeOf((*MockSyncClient)(nil).BroadcastAsync), arg0)
}

// BroadcastAsyncOrSendResponsible mocks base method.
func (m *MockSyncClient) BroadcastAsyncOrSendResponsible(arg0 *treechangeproto.TreeSyncMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BroadcastAsyncOrSendResponsible", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// BroadcastAsyncOrSendResponsible indicates an expected call of BroadcastAsyncOrSendResponsible.
func (mr *MockSyncClientMockRecorder) BroadcastAsyncOrSendResponsible(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastAsyncOrSendResponsible", reflect.TypeOf((*MockSyncClient)(nil).BroadcastAsyncOrSendResponsible), arg0)
}

// CreateFullSyncRequest mocks base method.
func (m *MockSyncClient) CreateFullSyncRequest(arg0 objecttree.ObjectTree, arg1, arg2 []string) (*treechangeproto.TreeSyncMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFullSyncRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(*treechangeproto.TreeSyncMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFullSyncRequest indicates an expected call of CreateFullSyncRequest.
func (mr *MockSyncClientMockRecorder) CreateFullSyncRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFullSyncRequest", reflect.TypeOf((*MockSyncClient)(nil).CreateFullSyncRequest), arg0, arg1, arg2)
}

// CreateFullSyncResponse mocks base method.
func (m *MockSyncClient) CreateFullSyncResponse(arg0 objecttree.ObjectTree, arg1, arg2 []string) (*treechangeproto.TreeSyncMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFullSyncResponse", arg0, arg1, arg2)
	ret0, _ := ret[0].(*treechangeproto.TreeSyncMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFullSyncResponse indicates an expected call of CreateFullSyncResponse.
func (mr *MockSyncClientMockRecorder) CreateFullSyncResponse(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFullSyncResponse", reflect.TypeOf((*MockSyncClient)(nil).CreateFullSyncResponse), arg0, arg1, arg2)
}

// CreateHeadUpdate mocks base method.
func (m *MockSyncClient) CreateHeadUpdate(arg0 objecttree.ObjectTree, arg1 []*treechangeproto.RawTreeChangeWithId) *treechangeproto.TreeSyncMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHeadUpdate", arg0, arg1)
	ret0, _ := ret[0].(*treechangeproto.TreeSyncMessage)
	return ret0
}

// CreateHeadUpdate indicates an expected call of CreateHeadUpdate.
func (mr *MockSyncClientMockRecorder) CreateHeadUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHeadUpdate", reflect.TypeOf((*MockSyncClient)(nil).CreateHeadUpdate), arg0, arg1)
}

// CreateNewTreeRequest mocks base method.
func (m *MockSyncClient) CreateNewTreeRequest() *treechangeproto.TreeSyncMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewTreeRequest")
	ret0, _ := ret[0].(*treechangeproto.TreeSyncMessage)
	return ret0
}

// CreateNewTreeRequest indicates an expected call of CreateNewTreeRequest.
func (mr *MockSyncClientMockRecorder) CreateNewTreeRequest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewTreeRequest", reflect.TypeOf((*MockSyncClient)(nil).CreateNewTreeRequest))
}

// SendAsync mocks base method.
func (m *MockSyncClient) SendAsync(arg0 string, arg1 *treechangeproto.TreeSyncMessage, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAsync", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAsync indicates an expected call of SendAsync.
func (mr *MockSyncClientMockRecorder) SendAsync(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAsync", reflect.TypeOf((*MockSyncClient)(nil).SendAsync), arg0, arg1, arg2)
}

// MockSyncTree is a mock of SyncTree interface.
type MockSyncTree struct {
	ctrl     *gomock.Controller
	recorder *MockSyncTreeMockRecorder
}

// MockSyncTreeMockRecorder is the mock recorder for MockSyncTree.
type MockSyncTreeMockRecorder struct {
	mock *MockSyncTree
}

// NewMockSyncTree creates a new mock instance.
func NewMockSyncTree(ctrl *gomock.Controller) *MockSyncTree {
	mock := &MockSyncTree{ctrl: ctrl}
	mock.recorder = &MockSyncTreeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSyncTree) EXPECT() *MockSyncTreeMockRecorder {
	return m.recorder
}

// AddContent mocks base method.
func (m *MockSyncTree) AddContent(arg0 context.Context, arg1 objecttree.SignableChangeContent) (objecttree.AddResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddContent", arg0, arg1)
	ret0, _ := ret[0].(objecttree.AddResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddContent indicates an expected call of AddContent.
func (mr *MockSyncTreeMockRecorder) AddContent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddContent", reflect.TypeOf((*MockSyncTree)(nil).AddContent), arg0, arg1)
}

// AddRawChanges mocks base method.
func (m *MockSyncTree) AddRawChanges(arg0 context.Context, arg1 objecttree.RawChangesPayload) (objecttree.AddResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRawChanges", arg0, arg1)
	ret0, _ := ret[0].(objecttree.AddResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRawChanges indicates an expected call of AddRawChanges.
func (mr *MockSyncTreeMockRecorder) AddRawChanges(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRawChanges", reflect.TypeOf((*MockSyncTree)(nil).AddRawChanges), arg0, arg1)
}

// ChangesAfterCommonSnapshot mocks base method.
func (m *MockSyncTree) ChangesAfterCommonSnapshot(arg0, arg1 []string) ([]*treechangeproto.RawTreeChangeWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangesAfterCommonSnapshot", arg0, arg1)
	ret0, _ := ret[0].([]*treechangeproto.RawTreeChangeWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangesAfterCommonSnapshot indicates an expected call of ChangesAfterCommonSnapshot.
func (mr *MockSyncTreeMockRecorder) ChangesAfterCommonSnapshot(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangesAfterCommonSnapshot", reflect.TypeOf((*MockSyncTree)(nil).ChangesAfterCommonSnapshot), arg0, arg1)
}

// Close mocks base method.
func (m *MockSyncTree) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSyncTreeMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSyncTree)(nil).Close))
}

// DebugDump mocks base method.
func (m *MockSyncTree) DebugDump(arg0 objecttree.DescriptionParser) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DebugDump", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DebugDump indicates an expected call of DebugDump.
func (mr *MockSyncTreeMockRecorder) DebugDump(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DebugDump", reflect.TypeOf((*MockSyncTree)(nil).DebugDump), arg0)
}

// Delete mocks base method.
func (m *MockSyncTree) Delete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSyncTreeMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSyncTree)(nil).Delete))
}

// HandleMessage mocks base method.
func (m *MockSyncTree) HandleMessage(arg0 context.Context, arg1 string, arg2 *spacesyncproto.ObjectSyncMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleMessage", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleMessage indicates an expected call of HandleMessage.
func (mr *MockSyncTreeMockRecorder) HandleMessage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMessage", reflect.TypeOf((*MockSyncTree)(nil).HandleMessage), arg0, arg1, arg2)
}

// HasChanges mocks base method.
func (m *MockSyncTree) HasChanges(arg0 ...string) bool {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HasChanges", varargs...)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasChanges indicates an expected call of HasChanges.
func (mr *MockSyncTreeMockRecorder) HasChanges(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasChanges", reflect.TypeOf((*MockSyncTree)(nil).HasChanges), arg0...)
}

// Header mocks base method.
func (m *MockSyncTree) Header() *treechangeproto.RawTreeChangeWithId {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(*treechangeproto.RawTreeChangeWithId)
	return ret0
}

// Header indicates an expected call of Header.
func (mr *MockSyncTreeMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockSyncTree)(nil).Header))
}

// Heads mocks base method.
func (m *MockSyncTree) Heads() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Heads")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Heads indicates an expected call of Heads.
func (mr *MockSyncTreeMockRecorder) Heads() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Heads", reflect.TypeOf((*MockSyncTree)(nil).Heads))
}

// Id mocks base method.
func (m *MockSyncTree) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockSyncTreeMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockSyncTree)(nil).Id))
}

// IterateFrom mocks base method.
func (m *MockSyncTree) IterateFrom(arg0 string, arg1 func([]byte) (interface{}, error), arg2 func(*objecttree.Change) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IterateFrom", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// IterateFrom indicates an expected call of IterateFrom.
func (mr *MockSyncTreeMockRecorder) IterateFrom(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateFrom", reflect.TypeOf((*MockSyncTree)(nil).IterateFrom), arg0, arg1, arg2)
}

// IterateRoot mocks base method.
func (m *MockSyncTree) IterateRoot(arg0 func([]byte) (interface{}, error), arg1 func(*objecttree.Change) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IterateRoot", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IterateRoot indicates an expected call of IterateRoot.
func (mr *MockSyncTreeMockRecorder) IterateRoot(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateRoot", reflect.TypeOf((*MockSyncTree)(nil).IterateRoot), arg0, arg1)
}

// Lock mocks base method.
func (m *MockSyncTree) Lock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Lock")
}

// Lock indicates an expected call of Lock.
func (mr *MockSyncTreeMockRecorder) Lock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockSyncTree)(nil).Lock))
}

// Ping mocks base method.
func (m *MockSyncTree) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockSyncTreeMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockSyncTree)(nil).Ping))
}

// RLock mocks base method.
func (m *MockSyncTree) RLock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RLock")
}

// RLock indicates an expected call of RLock.
func (mr *MockSyncTreeMockRecorder) RLock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RLock", reflect.TypeOf((*MockSyncTree)(nil).RLock))
}

// RUnlock mocks base method.
func (m *MockSyncTree) RUnlock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RUnlock")
}

// RUnlock indicates an expected call of RUnlock.
func (mr *MockSyncTreeMockRecorder) RUnlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RUnlock", reflect.TypeOf((*MockSyncTree)(nil).RUnlock))
}

// Root mocks base method.
func (m *MockSyncTree) Root() *objecttree.Change {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Root")
	ret0, _ := ret[0].(*objecttree.Change)
	return ret0
}

// Root indicates an expected call of Root.
func (mr *MockSyncTreeMockRecorder) Root() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Root", reflect.TypeOf((*MockSyncTree)(nil).Root))
}

// SnapshotPath mocks base method.
func (m *MockSyncTree) SnapshotPath() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotPath")
	ret0, _ := ret[0].([]string)
	return ret0
}

// SnapshotPath indicates an expected call of SnapshotPath.
func (mr *MockSyncTreeMockRecorder) SnapshotPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotPath", reflect.TypeOf((*MockSyncTree)(nil).SnapshotPath))
}

// Storage mocks base method.
func (m *MockSyncTree) Storage() treestorage.TreeStorage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Storage")
	ret0, _ := ret[0].(treestorage.TreeStorage)
	return ret0
}

// Storage indicates an expected call of Storage.
func (mr *MockSyncTreeMockRecorder) Storage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Storage", reflect.TypeOf((*MockSyncTree)(nil).Storage))
}

// Unlock mocks base method.
func (m *MockSyncTree) Unlock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unlock")
}

// Unlock indicates an expected call of Unlock.
func (mr *MockSyncTreeMockRecorder) Unlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockSyncTree)(nil).Unlock))
}

// UnmarshalledHeader mocks base method.
func (m *MockSyncTree) UnmarshalledHeader() *objecttree.Change {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnmarshalledHeader")
	ret0, _ := ret[0].(*objecttree.Change)
	return ret0
}

// UnmarshalledHeader indicates an expected call of UnmarshalledHeader.
func (mr *MockSyncTreeMockRecorder) UnmarshalledHeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalledHeader", reflect.TypeOf((*MockSyncTree)(nil).UnmarshalledHeader))
}

// MockReceiveQueue is a mock of ReceiveQueue interface.
type MockReceiveQueue struct {
	ctrl     *gomock.Controller
	recorder *MockReceiveQueueMockRecorder
}

// MockReceiveQueueMockRecorder is the mock recorder for MockReceiveQueue.
type MockReceiveQueueMockRecorder struct {
	mock *MockReceiveQueue
}

// NewMockReceiveQueue creates a new mock instance.
func NewMockReceiveQueue(ctrl *gomock.Controller) *MockReceiveQueue {
	mock := &MockReceiveQueue{ctrl: ctrl}
	mock.recorder = &MockReceiveQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReceiveQueue) EXPECT() *MockReceiveQueueMockRecorder {
	return m.recorder
}

// AddMessage mocks base method.
func (m *MockReceiveQueue) AddMessage(arg0 string, arg1 *treechangeproto.TreeSyncMessage, arg2 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMessage", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AddMessage indicates an expected call of AddMessage.
func (mr *MockReceiveQueueMockRecorder) AddMessage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMessage", reflect.TypeOf((*MockReceiveQueue)(nil).AddMessage), arg0, arg1, arg2)
}

// ClearQueue mocks base method.
func (m *MockReceiveQueue) ClearQueue(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearQueue", arg0)
}

// ClearQueue indicates an expected call of ClearQueue.
func (mr *MockReceiveQueueMockRecorder) ClearQueue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearQueue", reflect.TypeOf((*MockReceiveQueue)(nil).ClearQueue), arg0)
}

// GetMessage mocks base method.
func (m *MockReceiveQueue) GetMessage(arg0 string) (*treechangeproto.TreeSyncMessage, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessage", arg0)
	ret0, _ := ret[0].(*treechangeproto.TreeSyncMessage)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMessage indicates an expected call of GetMessage.
func (mr *MockReceiveQueueMockRecorder) GetMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessage", reflect.TypeOf((*MockReceiveQueue)(nil).GetMessage), arg0)
}

// MockHeadNotifiable is a mock of HeadNotifiable interface.
type MockHeadNotifiable struct {
	ctrl     *gomock.Controller
	recorder *MockHeadNotifiableMockRecorder
}

// MockHeadNotifiableMockRecorder is the mock recorder for MockHeadNotifiable.
type MockHeadNotifiableMockRecorder struct {
	mock *MockHeadNotifiable
}

// NewMockHeadNotifiable creates a new mock instance.
func NewMockHeadNotifiable(ctrl *gomock.Controller) *MockHeadNotifiable {
	mock := &MockHeadNotifiable{ctrl: ctrl}
	mock.recorder = &MockHeadNotifiableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHeadNotifiable) EXPECT() *MockHeadNotifiableMockRecorder {
	return m.recorder
}

// UpdateHeads mocks base method.
func (m *MockHeadNotifiable) UpdateHeads(arg0 string, arg1 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateHeads", arg0, arg1)
}

// UpdateHeads indicates an expected call of UpdateHeads.
func (mr *MockHeadNotifiableMockRecorder) UpdateHeads(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHeads", reflect.TypeOf((*MockHeadNotifiable)(nil).UpdateHeads), arg0, arg1)
}
