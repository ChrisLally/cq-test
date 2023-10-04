// Code generated by MockGen. DO NOT EDIT.
// Source: resiliencehub.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	resiliencehub "github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	gomock "github.com/golang/mock/gomock"
)

// MockResiliencehubClient is a mock of ResiliencehubClient interface.
type MockResiliencehubClient struct {
	ctrl     *gomock.Controller
	recorder *MockResiliencehubClientMockRecorder
}

// MockResiliencehubClientMockRecorder is the mock recorder for MockResiliencehubClient.
type MockResiliencehubClientMockRecorder struct {
	mock *MockResiliencehubClient
}

// NewMockResiliencehubClient creates a new mock instance.
func NewMockResiliencehubClient(ctrl *gomock.Controller) *MockResiliencehubClient {
	mock := &MockResiliencehubClient{ctrl: ctrl}
	mock.recorder = &MockResiliencehubClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResiliencehubClient) EXPECT() *MockResiliencehubClientMockRecorder {
	return m.recorder
}

// DescribeApp mocks base method.
func (m *MockResiliencehubClient) DescribeApp(arg0 context.Context, arg1 *resiliencehub.DescribeAppInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.DescribeAppOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeApp")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeApp", varargs...)
	ret0, _ := ret[0].(*resiliencehub.DescribeAppOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeApp indicates an expected call of DescribeApp.
func (mr *MockResiliencehubClientMockRecorder) DescribeApp(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeApp", reflect.TypeOf((*MockResiliencehubClient)(nil).DescribeApp), varargs...)
}

// DescribeAppAssessment mocks base method.
func (m *MockResiliencehubClient) DescribeAppAssessment(arg0 context.Context, arg1 *resiliencehub.DescribeAppAssessmentInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.DescribeAppAssessmentOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAppAssessment")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAppAssessment", varargs...)
	ret0, _ := ret[0].(*resiliencehub.DescribeAppAssessmentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAppAssessment indicates an expected call of DescribeAppAssessment.
func (mr *MockResiliencehubClientMockRecorder) DescribeAppAssessment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAppAssessment", reflect.TypeOf((*MockResiliencehubClient)(nil).DescribeAppAssessment), varargs...)
}

// DescribeAppVersion mocks base method.
func (m *MockResiliencehubClient) DescribeAppVersion(arg0 context.Context, arg1 *resiliencehub.DescribeAppVersionInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.DescribeAppVersionOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAppVersion")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAppVersion", varargs...)
	ret0, _ := ret[0].(*resiliencehub.DescribeAppVersionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAppVersion indicates an expected call of DescribeAppVersion.
func (mr *MockResiliencehubClientMockRecorder) DescribeAppVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAppVersion", reflect.TypeOf((*MockResiliencehubClient)(nil).DescribeAppVersion), varargs...)
}

// DescribeAppVersionAppComponent mocks base method.
func (m *MockResiliencehubClient) DescribeAppVersionAppComponent(arg0 context.Context, arg1 *resiliencehub.DescribeAppVersionAppComponentInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.DescribeAppVersionAppComponentOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAppVersionAppComponent")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAppVersionAppComponent", varargs...)
	ret0, _ := ret[0].(*resiliencehub.DescribeAppVersionAppComponentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAppVersionAppComponent indicates an expected call of DescribeAppVersionAppComponent.
func (mr *MockResiliencehubClientMockRecorder) DescribeAppVersionAppComponent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAppVersionAppComponent", reflect.TypeOf((*MockResiliencehubClient)(nil).DescribeAppVersionAppComponent), varargs...)
}

// DescribeAppVersionResource mocks base method.
func (m *MockResiliencehubClient) DescribeAppVersionResource(arg0 context.Context, arg1 *resiliencehub.DescribeAppVersionResourceInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.DescribeAppVersionResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAppVersionResource")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAppVersionResource", varargs...)
	ret0, _ := ret[0].(*resiliencehub.DescribeAppVersionResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAppVersionResource indicates an expected call of DescribeAppVersionResource.
func (mr *MockResiliencehubClientMockRecorder) DescribeAppVersionResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAppVersionResource", reflect.TypeOf((*MockResiliencehubClient)(nil).DescribeAppVersionResource), varargs...)
}

// DescribeAppVersionResourcesResolutionStatus mocks base method.
func (m *MockResiliencehubClient) DescribeAppVersionResourcesResolutionStatus(arg0 context.Context, arg1 *resiliencehub.DescribeAppVersionResourcesResolutionStatusInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.DescribeAppVersionResourcesResolutionStatusOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAppVersionResourcesResolutionStatus")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAppVersionResourcesResolutionStatus", varargs...)
	ret0, _ := ret[0].(*resiliencehub.DescribeAppVersionResourcesResolutionStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAppVersionResourcesResolutionStatus indicates an expected call of DescribeAppVersionResourcesResolutionStatus.
func (mr *MockResiliencehubClientMockRecorder) DescribeAppVersionResourcesResolutionStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAppVersionResourcesResolutionStatus", reflect.TypeOf((*MockResiliencehubClient)(nil).DescribeAppVersionResourcesResolutionStatus), varargs...)
}

// DescribeAppVersionTemplate mocks base method.
func (m *MockResiliencehubClient) DescribeAppVersionTemplate(arg0 context.Context, arg1 *resiliencehub.DescribeAppVersionTemplateInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.DescribeAppVersionTemplateOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAppVersionTemplate")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAppVersionTemplate", varargs...)
	ret0, _ := ret[0].(*resiliencehub.DescribeAppVersionTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAppVersionTemplate indicates an expected call of DescribeAppVersionTemplate.
func (mr *MockResiliencehubClientMockRecorder) DescribeAppVersionTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAppVersionTemplate", reflect.TypeOf((*MockResiliencehubClient)(nil).DescribeAppVersionTemplate), varargs...)
}

// DescribeDraftAppVersionResourcesImportStatus mocks base method.
func (m *MockResiliencehubClient) DescribeDraftAppVersionResourcesImportStatus(arg0 context.Context, arg1 *resiliencehub.DescribeDraftAppVersionResourcesImportStatusInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.DescribeDraftAppVersionResourcesImportStatusOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeDraftAppVersionResourcesImportStatus")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDraftAppVersionResourcesImportStatus", varargs...)
	ret0, _ := ret[0].(*resiliencehub.DescribeDraftAppVersionResourcesImportStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDraftAppVersionResourcesImportStatus indicates an expected call of DescribeDraftAppVersionResourcesImportStatus.
func (mr *MockResiliencehubClientMockRecorder) DescribeDraftAppVersionResourcesImportStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDraftAppVersionResourcesImportStatus", reflect.TypeOf((*MockResiliencehubClient)(nil).DescribeDraftAppVersionResourcesImportStatus), varargs...)
}

// DescribeResiliencyPolicy mocks base method.
func (m *MockResiliencehubClient) DescribeResiliencyPolicy(arg0 context.Context, arg1 *resiliencehub.DescribeResiliencyPolicyInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.DescribeResiliencyPolicyOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeResiliencyPolicy")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeResiliencyPolicy", varargs...)
	ret0, _ := ret[0].(*resiliencehub.DescribeResiliencyPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeResiliencyPolicy indicates an expected call of DescribeResiliencyPolicy.
func (mr *MockResiliencehubClientMockRecorder) DescribeResiliencyPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeResiliencyPolicy", reflect.TypeOf((*MockResiliencehubClient)(nil).DescribeResiliencyPolicy), varargs...)
}

// ListAlarmRecommendations mocks base method.
func (m *MockResiliencehubClient) ListAlarmRecommendations(arg0 context.Context, arg1 *resiliencehub.ListAlarmRecommendationsInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.ListAlarmRecommendationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAlarmRecommendations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAlarmRecommendations", varargs...)
	ret0, _ := ret[0].(*resiliencehub.ListAlarmRecommendationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAlarmRecommendations indicates an expected call of ListAlarmRecommendations.
func (mr *MockResiliencehubClientMockRecorder) ListAlarmRecommendations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAlarmRecommendations", reflect.TypeOf((*MockResiliencehubClient)(nil).ListAlarmRecommendations), varargs...)
}

// ListAppAssessmentComplianceDrifts mocks base method.
func (m *MockResiliencehubClient) ListAppAssessmentComplianceDrifts(arg0 context.Context, arg1 *resiliencehub.ListAppAssessmentComplianceDriftsInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.ListAppAssessmentComplianceDriftsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAppAssessmentComplianceDrifts")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAppAssessmentComplianceDrifts", varargs...)
	ret0, _ := ret[0].(*resiliencehub.ListAppAssessmentComplianceDriftsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAppAssessmentComplianceDrifts indicates an expected call of ListAppAssessmentComplianceDrifts.
func (mr *MockResiliencehubClientMockRecorder) ListAppAssessmentComplianceDrifts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAppAssessmentComplianceDrifts", reflect.TypeOf((*MockResiliencehubClient)(nil).ListAppAssessmentComplianceDrifts), varargs...)
}

// ListAppAssessments mocks base method.
func (m *MockResiliencehubClient) ListAppAssessments(arg0 context.Context, arg1 *resiliencehub.ListAppAssessmentsInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.ListAppAssessmentsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAppAssessments")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAppAssessments", varargs...)
	ret0, _ := ret[0].(*resiliencehub.ListAppAssessmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAppAssessments indicates an expected call of ListAppAssessments.
func (mr *MockResiliencehubClientMockRecorder) ListAppAssessments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAppAssessments", reflect.TypeOf((*MockResiliencehubClient)(nil).ListAppAssessments), varargs...)
}

// ListAppComponentCompliances mocks base method.
func (m *MockResiliencehubClient) ListAppComponentCompliances(arg0 context.Context, arg1 *resiliencehub.ListAppComponentCompliancesInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.ListAppComponentCompliancesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAppComponentCompliances")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAppComponentCompliances", varargs...)
	ret0, _ := ret[0].(*resiliencehub.ListAppComponentCompliancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAppComponentCompliances indicates an expected call of ListAppComponentCompliances.
func (mr *MockResiliencehubClientMockRecorder) ListAppComponentCompliances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAppComponentCompliances", reflect.TypeOf((*MockResiliencehubClient)(nil).ListAppComponentCompliances), varargs...)
}

// ListAppComponentRecommendations mocks base method.
func (m *MockResiliencehubClient) ListAppComponentRecommendations(arg0 context.Context, arg1 *resiliencehub.ListAppComponentRecommendationsInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.ListAppComponentRecommendationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAppComponentRecommendations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAppComponentRecommendations", varargs...)
	ret0, _ := ret[0].(*resiliencehub.ListAppComponentRecommendationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAppComponentRecommendations indicates an expected call of ListAppComponentRecommendations.
func (mr *MockResiliencehubClientMockRecorder) ListAppComponentRecommendations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAppComponentRecommendations", reflect.TypeOf((*MockResiliencehubClient)(nil).ListAppComponentRecommendations), varargs...)
}

// ListAppInputSources mocks base method.
func (m *MockResiliencehubClient) ListAppInputSources(arg0 context.Context, arg1 *resiliencehub.ListAppInputSourcesInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.ListAppInputSourcesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAppInputSources")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAppInputSources", varargs...)
	ret0, _ := ret[0].(*resiliencehub.ListAppInputSourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAppInputSources indicates an expected call of ListAppInputSources.
func (mr *MockResiliencehubClientMockRecorder) ListAppInputSources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAppInputSources", reflect.TypeOf((*MockResiliencehubClient)(nil).ListAppInputSources), varargs...)
}

// ListAppVersionAppComponents mocks base method.
func (m *MockResiliencehubClient) ListAppVersionAppComponents(arg0 context.Context, arg1 *resiliencehub.ListAppVersionAppComponentsInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.ListAppVersionAppComponentsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAppVersionAppComponents")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAppVersionAppComponents", varargs...)
	ret0, _ := ret[0].(*resiliencehub.ListAppVersionAppComponentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAppVersionAppComponents indicates an expected call of ListAppVersionAppComponents.
func (mr *MockResiliencehubClientMockRecorder) ListAppVersionAppComponents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAppVersionAppComponents", reflect.TypeOf((*MockResiliencehubClient)(nil).ListAppVersionAppComponents), varargs...)
}

// ListAppVersionResourceMappings mocks base method.
func (m *MockResiliencehubClient) ListAppVersionResourceMappings(arg0 context.Context, arg1 *resiliencehub.ListAppVersionResourceMappingsInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.ListAppVersionResourceMappingsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAppVersionResourceMappings")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAppVersionResourceMappings", varargs...)
	ret0, _ := ret[0].(*resiliencehub.ListAppVersionResourceMappingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAppVersionResourceMappings indicates an expected call of ListAppVersionResourceMappings.
func (mr *MockResiliencehubClientMockRecorder) ListAppVersionResourceMappings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAppVersionResourceMappings", reflect.TypeOf((*MockResiliencehubClient)(nil).ListAppVersionResourceMappings), varargs...)
}

// ListAppVersionResources mocks base method.
func (m *MockResiliencehubClient) ListAppVersionResources(arg0 context.Context, arg1 *resiliencehub.ListAppVersionResourcesInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.ListAppVersionResourcesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAppVersionResources")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAppVersionResources", varargs...)
	ret0, _ := ret[0].(*resiliencehub.ListAppVersionResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAppVersionResources indicates an expected call of ListAppVersionResources.
func (mr *MockResiliencehubClientMockRecorder) ListAppVersionResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAppVersionResources", reflect.TypeOf((*MockResiliencehubClient)(nil).ListAppVersionResources), varargs...)
}

// ListAppVersions mocks base method.
func (m *MockResiliencehubClient) ListAppVersions(arg0 context.Context, arg1 *resiliencehub.ListAppVersionsInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.ListAppVersionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAppVersions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAppVersions", varargs...)
	ret0, _ := ret[0].(*resiliencehub.ListAppVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAppVersions indicates an expected call of ListAppVersions.
func (mr *MockResiliencehubClientMockRecorder) ListAppVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAppVersions", reflect.TypeOf((*MockResiliencehubClient)(nil).ListAppVersions), varargs...)
}

// ListApps mocks base method.
func (m *MockResiliencehubClient) ListApps(arg0 context.Context, arg1 *resiliencehub.ListAppsInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.ListAppsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListApps")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListApps", varargs...)
	ret0, _ := ret[0].(*resiliencehub.ListAppsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApps indicates an expected call of ListApps.
func (mr *MockResiliencehubClientMockRecorder) ListApps(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApps", reflect.TypeOf((*MockResiliencehubClient)(nil).ListApps), varargs...)
}

// ListRecommendationTemplates mocks base method.
func (m *MockResiliencehubClient) ListRecommendationTemplates(arg0 context.Context, arg1 *resiliencehub.ListRecommendationTemplatesInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.ListRecommendationTemplatesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListRecommendationTemplates")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRecommendationTemplates", varargs...)
	ret0, _ := ret[0].(*resiliencehub.ListRecommendationTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRecommendationTemplates indicates an expected call of ListRecommendationTemplates.
func (mr *MockResiliencehubClientMockRecorder) ListRecommendationTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRecommendationTemplates", reflect.TypeOf((*MockResiliencehubClient)(nil).ListRecommendationTemplates), varargs...)
}

// ListResiliencyPolicies mocks base method.
func (m *MockResiliencehubClient) ListResiliencyPolicies(arg0 context.Context, arg1 *resiliencehub.ListResiliencyPoliciesInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.ListResiliencyPoliciesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListResiliencyPolicies")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResiliencyPolicies", varargs...)
	ret0, _ := ret[0].(*resiliencehub.ListResiliencyPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResiliencyPolicies indicates an expected call of ListResiliencyPolicies.
func (mr *MockResiliencehubClientMockRecorder) ListResiliencyPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResiliencyPolicies", reflect.TypeOf((*MockResiliencehubClient)(nil).ListResiliencyPolicies), varargs...)
}

// ListSopRecommendations mocks base method.
func (m *MockResiliencehubClient) ListSopRecommendations(arg0 context.Context, arg1 *resiliencehub.ListSopRecommendationsInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.ListSopRecommendationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListSopRecommendations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSopRecommendations", varargs...)
	ret0, _ := ret[0].(*resiliencehub.ListSopRecommendationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSopRecommendations indicates an expected call of ListSopRecommendations.
func (mr *MockResiliencehubClientMockRecorder) ListSopRecommendations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSopRecommendations", reflect.TypeOf((*MockResiliencehubClient)(nil).ListSopRecommendations), varargs...)
}

// ListSuggestedResiliencyPolicies mocks base method.
func (m *MockResiliencehubClient) ListSuggestedResiliencyPolicies(arg0 context.Context, arg1 *resiliencehub.ListSuggestedResiliencyPoliciesInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.ListSuggestedResiliencyPoliciesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListSuggestedResiliencyPolicies")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSuggestedResiliencyPolicies", varargs...)
	ret0, _ := ret[0].(*resiliencehub.ListSuggestedResiliencyPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSuggestedResiliencyPolicies indicates an expected call of ListSuggestedResiliencyPolicies.
func (mr *MockResiliencehubClientMockRecorder) ListSuggestedResiliencyPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSuggestedResiliencyPolicies", reflect.TypeOf((*MockResiliencehubClient)(nil).ListSuggestedResiliencyPolicies), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockResiliencehubClient) ListTagsForResource(arg0 context.Context, arg1 *resiliencehub.ListTagsForResourceInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
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
	ret0, _ := ret[0].(*resiliencehub.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockResiliencehubClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockResiliencehubClient)(nil).ListTagsForResource), varargs...)
}

// ListTestRecommendations mocks base method.
func (m *MockResiliencehubClient) ListTestRecommendations(arg0 context.Context, arg1 *resiliencehub.ListTestRecommendationsInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.ListTestRecommendationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTestRecommendations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTestRecommendations", varargs...)
	ret0, _ := ret[0].(*resiliencehub.ListTestRecommendationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTestRecommendations indicates an expected call of ListTestRecommendations.
func (mr *MockResiliencehubClientMockRecorder) ListTestRecommendations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTestRecommendations", reflect.TypeOf((*MockResiliencehubClient)(nil).ListTestRecommendations), varargs...)
}

// ListUnsupportedAppVersionResources mocks base method.
func (m *MockResiliencehubClient) ListUnsupportedAppVersionResources(arg0 context.Context, arg1 *resiliencehub.ListUnsupportedAppVersionResourcesInput, arg2 ...func(*resiliencehub.Options)) (*resiliencehub.ListUnsupportedAppVersionResourcesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &resiliencehub.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListUnsupportedAppVersionResources")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUnsupportedAppVersionResources", varargs...)
	ret0, _ := ret[0].(*resiliencehub.ListUnsupportedAppVersionResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUnsupportedAppVersionResources indicates an expected call of ListUnsupportedAppVersionResources.
func (mr *MockResiliencehubClientMockRecorder) ListUnsupportedAppVersionResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUnsupportedAppVersionResources", reflect.TypeOf((*MockResiliencehubClient)(nil).ListUnsupportedAppVersionResources), varargs...)
}