// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anytypeio/any-sync/commonspace/credentialprovider (interfaces: CredentialProvider)

// Package mock_credentialprovider is a generated GoMock package.
package mock_credentialprovider

import (
	context "context"
	reflect "reflect"

	spacesyncproto "github.com/anytypeio/any-sync/commonspace/spacesyncproto"
	gomock "github.com/golang/mock/gomock"
)

// MockCredentialProvider is a mock of CredentialProvider interface.
type MockCredentialProvider struct {
	ctrl     *gomock.Controller
	recorder *MockCredentialProviderMockRecorder
}

// MockCredentialProviderMockRecorder is the mock recorder for MockCredentialProvider.
type MockCredentialProviderMockRecorder struct {
	mock *MockCredentialProvider
}

// NewMockCredentialProvider creates a new mock instance.
func NewMockCredentialProvider(ctrl *gomock.Controller) *MockCredentialProvider {
	mock := &MockCredentialProvider{ctrl: ctrl}
	mock.recorder = &MockCredentialProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCredentialProvider) EXPECT() *MockCredentialProviderMockRecorder {
	return m.recorder
}

// GetCredential mocks base method.
func (m *MockCredentialProvider) GetCredential(arg0 context.Context, arg1 *spacesyncproto.RawSpaceHeaderWithId) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredential", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredential indicates an expected call of GetCredential.
func (mr *MockCredentialProviderMockRecorder) GetCredential(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredential", reflect.TypeOf((*MockCredentialProvider)(nil).GetCredential), arg0, arg1)
}
