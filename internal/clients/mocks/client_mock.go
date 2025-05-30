// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/clients/clients.go
//
// Generated by this command:
//
//	mockgen -source=./internal/clients/clients.go -destination=./internal/clients/mocks/client_mock.go -package=clientmocks
//

// Package clientmocks is a generated GoMock package.
package clientmocks

import (
	clients "planning/internal/clients"
	models "planning/internal/models"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIClients is a mock of IClients interface.
type MockIClients struct {
	ctrl     *gomock.Controller
	recorder *MockIClientsMockRecorder
	isgomock struct{}
}

// MockIClientsMockRecorder is the mock recorder for MockIClients.
type MockIClientsMockRecorder struct {
	mock *MockIClients
}

// NewMockIClients creates a new mock instance.
func NewMockIClients(ctrl *gomock.Controller) *MockIClients {
	mock := &MockIClients{ctrl: ctrl}
	mock.recorder = &MockIClientsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIClients) EXPECT() *MockIClientsMockRecorder {
	return m.recorder
}

// GetProvider1 mocks base method.
func (m *MockIClients) GetProvider1() clients.IProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProvider1")
	ret0, _ := ret[0].(clients.IProvider)
	return ret0
}

// GetProvider1 indicates an expected call of GetProvider1.
func (mr *MockIClientsMockRecorder) GetProvider1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProvider1", reflect.TypeOf((*MockIClients)(nil).GetProvider1))
}

// GetProvider2 mocks base method.
func (m *MockIClients) GetProvider2() clients.IProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProvider2")
	ret0, _ := ret[0].(clients.IProvider)
	return ret0
}

// GetProvider2 indicates an expected call of GetProvider2.
func (mr *MockIClientsMockRecorder) GetProvider2() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProvider2", reflect.TypeOf((*MockIClients)(nil).GetProvider2))
}

// MockIProvider is a mock of IProvider interface.
type MockIProvider struct {
	ctrl     *gomock.Controller
	recorder *MockIProviderMockRecorder
	isgomock struct{}
}

// MockIProviderMockRecorder is the mock recorder for MockIProvider.
type MockIProviderMockRecorder struct {
	mock *MockIProvider
}

// NewMockIProvider creates a new mock instance.
func NewMockIProvider(ctrl *gomock.Controller) *MockIProvider {
	mock := &MockIProvider{ctrl: ctrl}
	mock.recorder = &MockIProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProvider) EXPECT() *MockIProviderMockRecorder {
	return m.recorder
}

// FetchTasks mocks base method.
func (m *MockIProvider) FetchTasks() ([]models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchTasks")
	ret0, _ := ret[0].([]models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchTasks indicates an expected call of FetchTasks.
func (mr *MockIProviderMockRecorder) FetchTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchTasks", reflect.TypeOf((*MockIProvider)(nil).FetchTasks))
}
