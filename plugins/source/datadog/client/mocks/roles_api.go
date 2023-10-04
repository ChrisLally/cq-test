// Code generated by MockGen. DO NOT EDIT.
// Source: roles_api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	reflect "reflect"

	datadogV2 "github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	gomock "github.com/golang/mock/gomock"
)

// MockRolesAPIClient is a mock of RolesAPIClient interface.
type MockRolesAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockRolesAPIClientMockRecorder
}

// MockRolesAPIClientMockRecorder is the mock recorder for MockRolesAPIClient.
type MockRolesAPIClientMockRecorder struct {
	mock *MockRolesAPIClient
}

// NewMockRolesAPIClient creates a new mock instance.
func NewMockRolesAPIClient(ctrl *gomock.Controller) *MockRolesAPIClient {
	mock := &MockRolesAPIClient{ctrl: ctrl}
	mock.recorder = &MockRolesAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRolesAPIClient) EXPECT() *MockRolesAPIClientMockRecorder {
	return m.recorder
}

// ListPermissions mocks base method.
func (m *MockRolesAPIClient) ListPermissions(arg0 context.Context) (datadogV2.PermissionsResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPermissions", arg0)
	ret0, _ := ret[0].(datadogV2.PermissionsResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListPermissions indicates an expected call of ListPermissions.
func (mr *MockRolesAPIClientMockRecorder) ListPermissions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPermissions", reflect.TypeOf((*MockRolesAPIClient)(nil).ListPermissions), arg0)
}

// ListRolePermissions mocks base method.
func (m *MockRolesAPIClient) ListRolePermissions(arg0 context.Context, arg1 string) (datadogV2.PermissionsResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRolePermissions", arg0, arg1)
	ret0, _ := ret[0].(datadogV2.PermissionsResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListRolePermissions indicates an expected call of ListRolePermissions.
func (mr *MockRolesAPIClientMockRecorder) ListRolePermissions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRolePermissions", reflect.TypeOf((*MockRolesAPIClient)(nil).ListRolePermissions), arg0, arg1)
}

// ListRoleUsers mocks base method.
func (m *MockRolesAPIClient) ListRoleUsers(arg0 context.Context, arg1 string, arg2 ...datadogV2.ListRoleUsersOptionalParameters) (datadogV2.UsersResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRoleUsers", varargs...)
	ret0, _ := ret[0].(datadogV2.UsersResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListRoleUsers indicates an expected call of ListRoleUsers.
func (mr *MockRolesAPIClientMockRecorder) ListRoleUsers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoleUsers", reflect.TypeOf((*MockRolesAPIClient)(nil).ListRoleUsers), varargs...)
}

// ListRoles mocks base method.
func (m *MockRolesAPIClient) ListRoles(arg0 context.Context, arg1 ...datadogV2.ListRolesOptionalParameters) (datadogV2.RolesResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRoles", varargs...)
	ret0, _ := ret[0].(datadogV2.RolesResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListRoles indicates an expected call of ListRoles.
func (mr *MockRolesAPIClientMockRecorder) ListRoles(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoles", reflect.TypeOf((*MockRolesAPIClient)(nil).ListRoles), varargs...)
}
