// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/digitalocean/client (interfaces: FloatingIpsService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

// MockFloatingIpsService is a mock of FloatingIpsService interface.
type MockFloatingIpsService struct {
	ctrl     *gomock.Controller
	recorder *MockFloatingIpsServiceMockRecorder
}

// MockFloatingIpsServiceMockRecorder is the mock recorder for MockFloatingIpsService.
type MockFloatingIpsServiceMockRecorder struct {
	mock *MockFloatingIpsService
}

// NewMockFloatingIpsService creates a new mock instance.
func NewMockFloatingIpsService(ctrl *gomock.Controller) *MockFloatingIpsService {
	mock := &MockFloatingIpsService{ctrl: ctrl}
	mock.recorder = &MockFloatingIpsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFloatingIpsService) EXPECT() *MockFloatingIpsServiceMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockFloatingIpsService) List(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.FloatingIP, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]godo.FloatingIP)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockFloatingIpsServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFloatingIpsService)(nil).List), arg0, arg1)
}
