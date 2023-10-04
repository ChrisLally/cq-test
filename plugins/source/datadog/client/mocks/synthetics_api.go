// Code generated by MockGen. DO NOT EDIT.
// Source: synthetics_api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	reflect "reflect"

	datadog "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	datadogV1 "github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	gomock "github.com/golang/mock/gomock"
)

// MockSyntheticsAPIClient is a mock of SyntheticsAPIClient interface.
type MockSyntheticsAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockSyntheticsAPIClientMockRecorder
}

// MockSyntheticsAPIClientMockRecorder is the mock recorder for MockSyntheticsAPIClient.
type MockSyntheticsAPIClientMockRecorder struct {
	mock *MockSyntheticsAPIClient
}

// NewMockSyntheticsAPIClient creates a new mock instance.
func NewMockSyntheticsAPIClient(ctrl *gomock.Controller) *MockSyntheticsAPIClient {
	mock := &MockSyntheticsAPIClient{ctrl: ctrl}
	mock.recorder = &MockSyntheticsAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSyntheticsAPIClient) EXPECT() *MockSyntheticsAPIClientMockRecorder {
	return m.recorder
}

// ListGlobalVariables mocks base method.
func (m *MockSyntheticsAPIClient) ListGlobalVariables(arg0 context.Context) (datadogV1.SyntheticsListGlobalVariablesResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGlobalVariables", arg0)
	ret0, _ := ret[0].(datadogV1.SyntheticsListGlobalVariablesResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListGlobalVariables indicates an expected call of ListGlobalVariables.
func (mr *MockSyntheticsAPIClientMockRecorder) ListGlobalVariables(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGlobalVariables", reflect.TypeOf((*MockSyntheticsAPIClient)(nil).ListGlobalVariables), arg0)
}

// ListLocations mocks base method.
func (m *MockSyntheticsAPIClient) ListLocations(arg0 context.Context) (datadogV1.SyntheticsLocations, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLocations", arg0)
	ret0, _ := ret[0].(datadogV1.SyntheticsLocations)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListLocations indicates an expected call of ListLocations.
func (mr *MockSyntheticsAPIClientMockRecorder) ListLocations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLocations", reflect.TypeOf((*MockSyntheticsAPIClient)(nil).ListLocations), arg0)
}

// ListTests mocks base method.
func (m *MockSyntheticsAPIClient) ListTests(arg0 context.Context, arg1 ...datadogV1.ListTestsOptionalParameters) (datadogV1.SyntheticsListTestsResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTests", varargs...)
	ret0, _ := ret[0].(datadogV1.SyntheticsListTestsResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListTests indicates an expected call of ListTests.
func (mr *MockSyntheticsAPIClientMockRecorder) ListTests(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTests", reflect.TypeOf((*MockSyntheticsAPIClient)(nil).ListTests), varargs...)
}

// ListTestsWithPagination mocks base method.
func (m *MockSyntheticsAPIClient) ListTestsWithPagination(arg0 context.Context, arg1 ...datadogV1.ListTestsOptionalParameters) (<-chan datadog.PaginationResult[datadogV1.SyntheticsTestDetails], func()) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTestsWithPagination", varargs...)
	ret0, _ := ret[0].(<-chan datadog.PaginationResult[datadogV1.SyntheticsTestDetails])
	ret1, _ := ret[1].(func())
	return ret0, ret1
}

// ListTestsWithPagination indicates an expected call of ListTestsWithPagination.
func (mr *MockSyntheticsAPIClientMockRecorder) ListTestsWithPagination(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTestsWithPagination", reflect.TypeOf((*MockSyntheticsAPIClient)(nil).ListTestsWithPagination), varargs...)
}