// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/apps/chat/wshub/ws_interface.go

// Package mock_wshub is a generated GoMock package.

package mock_wshub

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/user/2019_1_newTeam2/models"
	http "net/http"
	reflect "reflect"
)

// MockIWSCommunicator is a mock of IWSCommunicator interface
type MockIWSCommunicator struct {
	ctrl     *gomock.Controller
	recorder *MockIWSCommunicatorMockRecorder
}

// MockIWSCommunicatorMockRecorder is the mock recorder for MockIWSCommunicator
type MockIWSCommunicatorMockRecorder struct {
	mock *MockIWSCommunicator
}

// NewMockIWSCommunicator creates a new mock instance
func NewMockIWSCommunicator(ctrl *gomock.Controller) *MockIWSCommunicator {
	mock := &MockIWSCommunicator{ctrl: ctrl}
	mock.recorder = &MockIWSCommunicatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIWSCommunicator) EXPECT() *MockIWSCommunicatorMockRecorder {
	return m.recorder
}

// AddClient mocks base method
func (m *MockIWSCommunicator) AddClient(w http.ResponseWriter, r *http.Request, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddClient", w, r, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddClient indicates an expected call of AddClient
func (mr *MockIWSCommunicatorMockRecorder) AddClient(w, r, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddClient", reflect.TypeOf((*MockIWSCommunicator)(nil).AddClient), w, r, id)
}

// SendToClient mocks base method
func (m *MockIWSCommunicator) SendToClient(mes *models.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendToClient", mes)
}

// SendToClient indicates an expected call of SendToClient
func (mr *MockIWSCommunicatorMockRecorder) SendToClient(mes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendToClient", reflect.TypeOf((*MockIWSCommunicator)(nil).SendToClient), mes)
}

// DeleteClient mocks base method
func (m *MockIWSCommunicator) DeleteClient(ID int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteClient", ID)
}

// DeleteClient indicates an expected call of DeleteClient
func (mr *MockIWSCommunicatorMockRecorder) DeleteClient(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClient", reflect.TypeOf((*MockIWSCommunicator)(nil).DeleteClient), ID)
}
