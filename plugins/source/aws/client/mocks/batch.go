// Code generated by MockGen. DO NOT EDIT.
// Source: batch.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	batch "github.com/aws/aws-sdk-go-v2/service/batch"
	gomock "github.com/golang/mock/gomock"
)

// MockBatchClient is a mock of BatchClient interface.
type MockBatchClient struct {
	ctrl     *gomock.Controller
	recorder *MockBatchClientMockRecorder
}

// MockBatchClientMockRecorder is the mock recorder for MockBatchClient.
type MockBatchClientMockRecorder struct {
	mock *MockBatchClient
}

// NewMockBatchClient creates a new mock instance.
func NewMockBatchClient(ctrl *gomock.Controller) *MockBatchClient {
	mock := &MockBatchClient{ctrl: ctrl}
	mock.recorder = &MockBatchClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBatchClient) EXPECT() *MockBatchClientMockRecorder {
	return m.recorder
}

// DescribeComputeEnvironments mocks base method.
func (m *MockBatchClient) DescribeComputeEnvironments(arg0 context.Context, arg1 *batch.DescribeComputeEnvironmentsInput, arg2 ...func(*batch.Options)) (*batch.DescribeComputeEnvironmentsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &batch.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeComputeEnvironments")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeComputeEnvironments", varargs...)
	ret0, _ := ret[0].(*batch.DescribeComputeEnvironmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeComputeEnvironments indicates an expected call of DescribeComputeEnvironments.
func (mr *MockBatchClientMockRecorder) DescribeComputeEnvironments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeComputeEnvironments", reflect.TypeOf((*MockBatchClient)(nil).DescribeComputeEnvironments), varargs...)
}

// DescribeJobDefinitions mocks base method.
func (m *MockBatchClient) DescribeJobDefinitions(arg0 context.Context, arg1 *batch.DescribeJobDefinitionsInput, arg2 ...func(*batch.Options)) (*batch.DescribeJobDefinitionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &batch.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeJobDefinitions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeJobDefinitions", varargs...)
	ret0, _ := ret[0].(*batch.DescribeJobDefinitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeJobDefinitions indicates an expected call of DescribeJobDefinitions.
func (mr *MockBatchClientMockRecorder) DescribeJobDefinitions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeJobDefinitions", reflect.TypeOf((*MockBatchClient)(nil).DescribeJobDefinitions), varargs...)
}

// DescribeJobQueues mocks base method.
func (m *MockBatchClient) DescribeJobQueues(arg0 context.Context, arg1 *batch.DescribeJobQueuesInput, arg2 ...func(*batch.Options)) (*batch.DescribeJobQueuesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &batch.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeJobQueues")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeJobQueues", varargs...)
	ret0, _ := ret[0].(*batch.DescribeJobQueuesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeJobQueues indicates an expected call of DescribeJobQueues.
func (mr *MockBatchClientMockRecorder) DescribeJobQueues(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeJobQueues", reflect.TypeOf((*MockBatchClient)(nil).DescribeJobQueues), varargs...)
}

// DescribeJobs mocks base method.
func (m *MockBatchClient) DescribeJobs(arg0 context.Context, arg1 *batch.DescribeJobsInput, arg2 ...func(*batch.Options)) (*batch.DescribeJobsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &batch.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeJobs")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeJobs", varargs...)
	ret0, _ := ret[0].(*batch.DescribeJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeJobs indicates an expected call of DescribeJobs.
func (mr *MockBatchClientMockRecorder) DescribeJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeJobs", reflect.TypeOf((*MockBatchClient)(nil).DescribeJobs), varargs...)
}

// DescribeSchedulingPolicies mocks base method.
func (m *MockBatchClient) DescribeSchedulingPolicies(arg0 context.Context, arg1 *batch.DescribeSchedulingPoliciesInput, arg2 ...func(*batch.Options)) (*batch.DescribeSchedulingPoliciesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &batch.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeSchedulingPolicies")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSchedulingPolicies", varargs...)
	ret0, _ := ret[0].(*batch.DescribeSchedulingPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSchedulingPolicies indicates an expected call of DescribeSchedulingPolicies.
func (mr *MockBatchClientMockRecorder) DescribeSchedulingPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSchedulingPolicies", reflect.TypeOf((*MockBatchClient)(nil).DescribeSchedulingPolicies), varargs...)
}

// ListJobs mocks base method.
func (m *MockBatchClient) ListJobs(arg0 context.Context, arg1 *batch.ListJobsInput, arg2 ...func(*batch.Options)) (*batch.ListJobsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &batch.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListJobs")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListJobs", varargs...)
	ret0, _ := ret[0].(*batch.ListJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListJobs indicates an expected call of ListJobs.
func (mr *MockBatchClientMockRecorder) ListJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListJobs", reflect.TypeOf((*MockBatchClient)(nil).ListJobs), varargs...)
}

// ListSchedulingPolicies mocks base method.
func (m *MockBatchClient) ListSchedulingPolicies(arg0 context.Context, arg1 *batch.ListSchedulingPoliciesInput, arg2 ...func(*batch.Options)) (*batch.ListSchedulingPoliciesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &batch.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListSchedulingPolicies")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSchedulingPolicies", varargs...)
	ret0, _ := ret[0].(*batch.ListSchedulingPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchedulingPolicies indicates an expected call of ListSchedulingPolicies.
func (mr *MockBatchClientMockRecorder) ListSchedulingPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedulingPolicies", reflect.TypeOf((*MockBatchClient)(nil).ListSchedulingPolicies), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockBatchClient) ListTagsForResource(arg0 context.Context, arg1 *batch.ListTagsForResourceInput, arg2 ...func(*batch.Options)) (*batch.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &batch.Options{}
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
	ret0, _ := ret[0].(*batch.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockBatchClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockBatchClient)(nil).ListTagsForResource), varargs...)
}
