// Code generated by MockGen. DO NOT EDIT.
// Source: checker.go

// Package mock_app is a generated GoMock package.
package app

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockChecker is a mock of Checker interface.
type MockChecker struct {
	ctrl     *gomock.Controller
	recorder *MockCheckerMockRecorder
}

// MockCheckerMockRecorder is the mock recorder for MockChecker.
type MockCheckerMockRecorder struct {
	mock *MockChecker
}

// NewMockChecker creates a new mock instance.
func NewMockChecker(ctrl *gomock.Controller) *MockChecker {
	mock := &MockChecker{ctrl: ctrl}
	mock.recorder = &MockCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChecker) EXPECT() *MockCheckerMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockChecker) Check() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Check indicates an expected call of Check.
func (mr *MockCheckerMockRecorder) Check() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockChecker)(nil).Check))
}
