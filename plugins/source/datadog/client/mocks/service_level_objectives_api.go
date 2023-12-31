// Code generated by MockGen. DO NOT EDIT.
// Source: service_level_objectives_api.go

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

// MockServiceLevelObjectivesAPIClient is a mock of ServiceLevelObjectivesAPIClient interface.
type MockServiceLevelObjectivesAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockServiceLevelObjectivesAPIClientMockRecorder
}

// MockServiceLevelObjectivesAPIClientMockRecorder is the mock recorder for MockServiceLevelObjectivesAPIClient.
type MockServiceLevelObjectivesAPIClientMockRecorder struct {
	mock *MockServiceLevelObjectivesAPIClient
}

// NewMockServiceLevelObjectivesAPIClient creates a new mock instance.
func NewMockServiceLevelObjectivesAPIClient(ctrl *gomock.Controller) *MockServiceLevelObjectivesAPIClient {
	mock := &MockServiceLevelObjectivesAPIClient{ctrl: ctrl}
	mock.recorder = &MockServiceLevelObjectivesAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceLevelObjectivesAPIClient) EXPECT() *MockServiceLevelObjectivesAPIClientMockRecorder {
	return m.recorder
}

// ListSLOs mocks base method.
func (m *MockServiceLevelObjectivesAPIClient) ListSLOs(arg0 context.Context, arg1 ...datadogV1.ListSLOsOptionalParameters) (datadogV1.SLOListResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSLOs", varargs...)
	ret0, _ := ret[0].(datadogV1.SLOListResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListSLOs indicates an expected call of ListSLOs.
func (mr *MockServiceLevelObjectivesAPIClientMockRecorder) ListSLOs(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSLOs", reflect.TypeOf((*MockServiceLevelObjectivesAPIClient)(nil).ListSLOs), varargs...)
}

// ListSLOsWithPagination mocks base method.
func (m *MockServiceLevelObjectivesAPIClient) ListSLOsWithPagination(arg0 context.Context, arg1 ...datadogV1.ListSLOsOptionalParameters) (<-chan datadog.PaginationResult[datadogV1.ServiceLevelObjective], func()) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSLOsWithPagination", varargs...)
	ret0, _ := ret[0].(<-chan datadog.PaginationResult[datadogV1.ServiceLevelObjective])
	ret1, _ := ret[1].(func())
	return ret0, ret1
}

// ListSLOsWithPagination indicates an expected call of ListSLOsWithPagination.
func (mr *MockServiceLevelObjectivesAPIClientMockRecorder) ListSLOsWithPagination(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSLOsWithPagination", reflect.TypeOf((*MockServiceLevelObjectivesAPIClient)(nil).ListSLOsWithPagination), varargs...)
}
