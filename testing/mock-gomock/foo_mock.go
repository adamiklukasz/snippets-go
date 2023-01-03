// Code generated by MockGen. DO NOT EDIT.
// Source: foo.go

// Package mocking is a generated GoMock package.
package mocking

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockFoo is a mock of Foo interface
type MockFoo struct {
	ctrl     *gomock.Controller
	recorder *MockFooMockRecorder
}

// MockFooMockRecorder is the mock recorder for MockFoo
type MockFooMockRecorder struct {
	mock *MockFoo
}

// NewMockFoo creates a new mock instance
func NewMockFoo(ctrl *gomock.Controller) *MockFoo {
	mock := &MockFoo{ctrl: ctrl}
	mock.recorder = &MockFooMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFoo) EXPECT() *MockFooMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockFoo) Do(arg0 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// Do indicates an expected call of Do
func (mr *MockFooMockRecorder) Do(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockFoo)(nil).Do), arg0)
}
