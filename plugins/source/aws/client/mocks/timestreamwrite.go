// Code generated by MockGen. DO NOT EDIT.
// Source: timestreamwrite.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	timestreamwrite "github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	gomock "github.com/golang/mock/gomock"
)

// MockTimestreamwriteClient is a mock of TimestreamwriteClient interface.
type MockTimestreamwriteClient struct {
	ctrl     *gomock.Controller
	recorder *MockTimestreamwriteClientMockRecorder
}

// MockTimestreamwriteClientMockRecorder is the mock recorder for MockTimestreamwriteClient.
type MockTimestreamwriteClientMockRecorder struct {
	mock *MockTimestreamwriteClient
}

// NewMockTimestreamwriteClient creates a new mock instance.
func NewMockTimestreamwriteClient(ctrl *gomock.Controller) *MockTimestreamwriteClient {
	mock := &MockTimestreamwriteClient{ctrl: ctrl}
	mock.recorder = &MockTimestreamwriteClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTimestreamwriteClient) EXPECT() *MockTimestreamwriteClientMockRecorder {
	return m.recorder
}

// DescribeBatchLoadTask mocks base method.
func (m *MockTimestreamwriteClient) DescribeBatchLoadTask(arg0 context.Context, arg1 *timestreamwrite.DescribeBatchLoadTaskInput, arg2 ...func(*timestreamwrite.Options)) (*timestreamwrite.DescribeBatchLoadTaskOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &timestreamwrite.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeBatchLoadTask")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeBatchLoadTask", varargs...)
	ret0, _ := ret[0].(*timestreamwrite.DescribeBatchLoadTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeBatchLoadTask indicates an expected call of DescribeBatchLoadTask.
func (mr *MockTimestreamwriteClientMockRecorder) DescribeBatchLoadTask(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeBatchLoadTask", reflect.TypeOf((*MockTimestreamwriteClient)(nil).DescribeBatchLoadTask), varargs...)
}

// DescribeDatabase mocks base method.
func (m *MockTimestreamwriteClient) DescribeDatabase(arg0 context.Context, arg1 *timestreamwrite.DescribeDatabaseInput, arg2 ...func(*timestreamwrite.Options)) (*timestreamwrite.DescribeDatabaseOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &timestreamwrite.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeDatabase")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDatabase", varargs...)
	ret0, _ := ret[0].(*timestreamwrite.DescribeDatabaseOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDatabase indicates an expected call of DescribeDatabase.
func (mr *MockTimestreamwriteClientMockRecorder) DescribeDatabase(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDatabase", reflect.TypeOf((*MockTimestreamwriteClient)(nil).DescribeDatabase), varargs...)
}

// DescribeEndpoints mocks base method.
func (m *MockTimestreamwriteClient) DescribeEndpoints(arg0 context.Context, arg1 *timestreamwrite.DescribeEndpointsInput, arg2 ...func(*timestreamwrite.Options)) (*timestreamwrite.DescribeEndpointsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &timestreamwrite.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeEndpoints")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEndpoints", varargs...)
	ret0, _ := ret[0].(*timestreamwrite.DescribeEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEndpoints indicates an expected call of DescribeEndpoints.
func (mr *MockTimestreamwriteClientMockRecorder) DescribeEndpoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEndpoints", reflect.TypeOf((*MockTimestreamwriteClient)(nil).DescribeEndpoints), varargs...)
}

// DescribeTable mocks base method.
func (m *MockTimestreamwriteClient) DescribeTable(arg0 context.Context, arg1 *timestreamwrite.DescribeTableInput, arg2 ...func(*timestreamwrite.Options)) (*timestreamwrite.DescribeTableOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &timestreamwrite.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeTable")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTable", varargs...)
	ret0, _ := ret[0].(*timestreamwrite.DescribeTableOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTable indicates an expected call of DescribeTable.
func (mr *MockTimestreamwriteClientMockRecorder) DescribeTable(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTable", reflect.TypeOf((*MockTimestreamwriteClient)(nil).DescribeTable), varargs...)
}

// ListBatchLoadTasks mocks base method.
func (m *MockTimestreamwriteClient) ListBatchLoadTasks(arg0 context.Context, arg1 *timestreamwrite.ListBatchLoadTasksInput, arg2 ...func(*timestreamwrite.Options)) (*timestreamwrite.ListBatchLoadTasksOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &timestreamwrite.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListBatchLoadTasks")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBatchLoadTasks", varargs...)
	ret0, _ := ret[0].(*timestreamwrite.ListBatchLoadTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBatchLoadTasks indicates an expected call of ListBatchLoadTasks.
func (mr *MockTimestreamwriteClientMockRecorder) ListBatchLoadTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBatchLoadTasks", reflect.TypeOf((*MockTimestreamwriteClient)(nil).ListBatchLoadTasks), varargs...)
}

// ListDatabases mocks base method.
func (m *MockTimestreamwriteClient) ListDatabases(arg0 context.Context, arg1 *timestreamwrite.ListDatabasesInput, arg2 ...func(*timestreamwrite.Options)) (*timestreamwrite.ListDatabasesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &timestreamwrite.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListDatabases")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDatabases", varargs...)
	ret0, _ := ret[0].(*timestreamwrite.ListDatabasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDatabases indicates an expected call of ListDatabases.
func (mr *MockTimestreamwriteClientMockRecorder) ListDatabases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDatabases", reflect.TypeOf((*MockTimestreamwriteClient)(nil).ListDatabases), varargs...)
}

// ListTables mocks base method.
func (m *MockTimestreamwriteClient) ListTables(arg0 context.Context, arg1 *timestreamwrite.ListTablesInput, arg2 ...func(*timestreamwrite.Options)) (*timestreamwrite.ListTablesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &timestreamwrite.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTables")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTables", varargs...)
	ret0, _ := ret[0].(*timestreamwrite.ListTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTables indicates an expected call of ListTables.
func (mr *MockTimestreamwriteClientMockRecorder) ListTables(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTables", reflect.TypeOf((*MockTimestreamwriteClient)(nil).ListTables), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockTimestreamwriteClient) ListTagsForResource(arg0 context.Context, arg1 *timestreamwrite.ListTagsForResourceInput, arg2 ...func(*timestreamwrite.Options)) (*timestreamwrite.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &timestreamwrite.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTagsForResource")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*timestreamwrite.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockTimestreamwriteClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockTimestreamwriteClient)(nil).ListTagsForResource), varargs...)
}