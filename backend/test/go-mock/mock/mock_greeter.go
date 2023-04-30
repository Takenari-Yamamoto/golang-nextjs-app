// Code generated by MockGen. DO NOT EDIT.
// Source: greeter.go

// Package mock_main is a generated GoMock package.
package mock_main

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGreeter is a mock of Greeter interface.
type MockGreeter struct {
	ctrl     *gomock.Controller
	recorder *MockGreeterMockRecorder
}

// MockGreeterMockRecorder is the mock recorder for MockGreeter.
type MockGreeterMockRecorder struct {
	mock *MockGreeter
}

// NewMockGreeter creates a new mock instance.
func NewMockGreeter(ctrl *gomock.Controller) *MockGreeter {
	mock := &MockGreeter{ctrl: ctrl}
	mock.recorder = &MockGreeterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGreeter) EXPECT() *MockGreeterMockRecorder {
	return m.recorder
}

// Greet mocks base method.
func (m *MockGreeter) Greet(name string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Greet", name)
	ret0, _ := ret[0].(string)
	return ret0
}

// Greet indicates an expected call of Greet.
func (mr *MockGreeterMockRecorder) Greet(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Greet", reflect.TypeOf((*MockGreeter)(nil).Greet), name)
}
