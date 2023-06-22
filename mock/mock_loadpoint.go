// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/evcc-io/evcc/core (interfaces: Updater)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUpdater is a mock of Updater interface.
type MockUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockUpdaterMockRecorder
}

// MockUpdaterMockRecorder is the mock recorder for MockUpdater.
type MockUpdaterMockRecorder struct {
	mock *MockUpdater
}

// NewMockUpdater creates a new mock instance.
func NewMockUpdater(ctrl *gomock.Controller) *MockUpdater {
	mock := &MockUpdater{ctrl: ctrl}
	mock.recorder = &MockUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdater) EXPECT() *MockUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockUpdater) Update(arg0 float64, arg1, arg2 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Update", arg0, arg1, arg2)
}

// Update indicates an expected call of Update.
func (mr *MockUpdaterMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUpdater)(nil).Update), arg0, arg1, arg2)
}
