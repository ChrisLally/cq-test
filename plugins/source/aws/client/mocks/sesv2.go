// Code generated by MockGen. DO NOT EDIT.
// Source: sesv2.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	sesv2 "github.com/aws/aws-sdk-go-v2/service/sesv2"
	gomock "github.com/golang/mock/gomock"
)

// MockSesv2Client is a mock of Sesv2Client interface.
type MockSesv2Client struct {
	ctrl     *gomock.Controller
	recorder *MockSesv2ClientMockRecorder
}

// MockSesv2ClientMockRecorder is the mock recorder for MockSesv2Client.
type MockSesv2ClientMockRecorder struct {
	mock *MockSesv2Client
}

// NewMockSesv2Client creates a new mock instance.
func NewMockSesv2Client(ctrl *gomock.Controller) *MockSesv2Client {
	mock := &MockSesv2Client{ctrl: ctrl}
	mock.recorder = &MockSesv2ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSesv2Client) EXPECT() *MockSesv2ClientMockRecorder {
	return m.recorder
}

// BatchGetMetricData mocks base method.
func (m *MockSesv2Client) BatchGetMetricData(arg0 context.Context, arg1 *sesv2.BatchGetMetricDataInput, arg2 ...func(*sesv2.Options)) (*sesv2.BatchGetMetricDataOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to BatchGetMetricData")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetMetricData", varargs...)
	ret0, _ := ret[0].(*sesv2.BatchGetMetricDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetMetricData indicates an expected call of BatchGetMetricData.
func (mr *MockSesv2ClientMockRecorder) BatchGetMetricData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetMetricData", reflect.TypeOf((*MockSesv2Client)(nil).BatchGetMetricData), varargs...)
}

// GetAccount mocks base method.
func (m *MockSesv2Client) GetAccount(arg0 context.Context, arg1 *sesv2.GetAccountInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetAccountOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetAccount")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccount", varargs...)
	ret0, _ := ret[0].(*sesv2.GetAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockSesv2ClientMockRecorder) GetAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockSesv2Client)(nil).GetAccount), varargs...)
}

// GetBlacklistReports mocks base method.
func (m *MockSesv2Client) GetBlacklistReports(arg0 context.Context, arg1 *sesv2.GetBlacklistReportsInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetBlacklistReportsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetBlacklistReports")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlacklistReports", varargs...)
	ret0, _ := ret[0].(*sesv2.GetBlacklistReportsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlacklistReports indicates an expected call of GetBlacklistReports.
func (mr *MockSesv2ClientMockRecorder) GetBlacklistReports(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlacklistReports", reflect.TypeOf((*MockSesv2Client)(nil).GetBlacklistReports), varargs...)
}

// GetConfigurationSet mocks base method.
func (m *MockSesv2Client) GetConfigurationSet(arg0 context.Context, arg1 *sesv2.GetConfigurationSetInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetConfigurationSetOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetConfigurationSet")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConfigurationSet", varargs...)
	ret0, _ := ret[0].(*sesv2.GetConfigurationSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigurationSet indicates an expected call of GetConfigurationSet.
func (mr *MockSesv2ClientMockRecorder) GetConfigurationSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigurationSet", reflect.TypeOf((*MockSesv2Client)(nil).GetConfigurationSet), varargs...)
}

// GetConfigurationSetEventDestinations mocks base method.
func (m *MockSesv2Client) GetConfigurationSetEventDestinations(arg0 context.Context, arg1 *sesv2.GetConfigurationSetEventDestinationsInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetConfigurationSetEventDestinationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetConfigurationSetEventDestinations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConfigurationSetEventDestinations", varargs...)
	ret0, _ := ret[0].(*sesv2.GetConfigurationSetEventDestinationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigurationSetEventDestinations indicates an expected call of GetConfigurationSetEventDestinations.
func (mr *MockSesv2ClientMockRecorder) GetConfigurationSetEventDestinations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigurationSetEventDestinations", reflect.TypeOf((*MockSesv2Client)(nil).GetConfigurationSetEventDestinations), varargs...)
}

// GetContact mocks base method.
func (m *MockSesv2Client) GetContact(arg0 context.Context, arg1 *sesv2.GetContactInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetContactOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetContact")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContact", varargs...)
	ret0, _ := ret[0].(*sesv2.GetContactOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContact indicates an expected call of GetContact.
func (mr *MockSesv2ClientMockRecorder) GetContact(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContact", reflect.TypeOf((*MockSesv2Client)(nil).GetContact), varargs...)
}

// GetContactList mocks base method.
func (m *MockSesv2Client) GetContactList(arg0 context.Context, arg1 *sesv2.GetContactListInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetContactListOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetContactList")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContactList", varargs...)
	ret0, _ := ret[0].(*sesv2.GetContactListOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContactList indicates an expected call of GetContactList.
func (mr *MockSesv2ClientMockRecorder) GetContactList(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContactList", reflect.TypeOf((*MockSesv2Client)(nil).GetContactList), varargs...)
}

// GetCustomVerificationEmailTemplate mocks base method.
func (m *MockSesv2Client) GetCustomVerificationEmailTemplate(arg0 context.Context, arg1 *sesv2.GetCustomVerificationEmailTemplateInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetCustomVerificationEmailTemplateOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetCustomVerificationEmailTemplate")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCustomVerificationEmailTemplate", varargs...)
	ret0, _ := ret[0].(*sesv2.GetCustomVerificationEmailTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomVerificationEmailTemplate indicates an expected call of GetCustomVerificationEmailTemplate.
func (mr *MockSesv2ClientMockRecorder) GetCustomVerificationEmailTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomVerificationEmailTemplate", reflect.TypeOf((*MockSesv2Client)(nil).GetCustomVerificationEmailTemplate), varargs...)
}

// GetDedicatedIp mocks base method.
func (m *MockSesv2Client) GetDedicatedIp(arg0 context.Context, arg1 *sesv2.GetDedicatedIpInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetDedicatedIpOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDedicatedIp")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDedicatedIp", varargs...)
	ret0, _ := ret[0].(*sesv2.GetDedicatedIpOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDedicatedIp indicates an expected call of GetDedicatedIp.
func (mr *MockSesv2ClientMockRecorder) GetDedicatedIp(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDedicatedIp", reflect.TypeOf((*MockSesv2Client)(nil).GetDedicatedIp), varargs...)
}

// GetDedicatedIpPool mocks base method.
func (m *MockSesv2Client) GetDedicatedIpPool(arg0 context.Context, arg1 *sesv2.GetDedicatedIpPoolInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetDedicatedIpPoolOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDedicatedIpPool")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDedicatedIpPool", varargs...)
	ret0, _ := ret[0].(*sesv2.GetDedicatedIpPoolOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDedicatedIpPool indicates an expected call of GetDedicatedIpPool.
func (mr *MockSesv2ClientMockRecorder) GetDedicatedIpPool(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDedicatedIpPool", reflect.TypeOf((*MockSesv2Client)(nil).GetDedicatedIpPool), varargs...)
}

// GetDedicatedIps mocks base method.
func (m *MockSesv2Client) GetDedicatedIps(arg0 context.Context, arg1 *sesv2.GetDedicatedIpsInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetDedicatedIpsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDedicatedIps")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDedicatedIps", varargs...)
	ret0, _ := ret[0].(*sesv2.GetDedicatedIpsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDedicatedIps indicates an expected call of GetDedicatedIps.
func (mr *MockSesv2ClientMockRecorder) GetDedicatedIps(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDedicatedIps", reflect.TypeOf((*MockSesv2Client)(nil).GetDedicatedIps), varargs...)
}

// GetDeliverabilityDashboardOptions mocks base method.
func (m *MockSesv2Client) GetDeliverabilityDashboardOptions(arg0 context.Context, arg1 *sesv2.GetDeliverabilityDashboardOptionsInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetDeliverabilityDashboardOptionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDeliverabilityDashboardOptions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDeliverabilityDashboardOptions", varargs...)
	ret0, _ := ret[0].(*sesv2.GetDeliverabilityDashboardOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeliverabilityDashboardOptions indicates an expected call of GetDeliverabilityDashboardOptions.
func (mr *MockSesv2ClientMockRecorder) GetDeliverabilityDashboardOptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeliverabilityDashboardOptions", reflect.TypeOf((*MockSesv2Client)(nil).GetDeliverabilityDashboardOptions), varargs...)
}

// GetDeliverabilityTestReport mocks base method.
func (m *MockSesv2Client) GetDeliverabilityTestReport(arg0 context.Context, arg1 *sesv2.GetDeliverabilityTestReportInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetDeliverabilityTestReportOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDeliverabilityTestReport")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDeliverabilityTestReport", varargs...)
	ret0, _ := ret[0].(*sesv2.GetDeliverabilityTestReportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeliverabilityTestReport indicates an expected call of GetDeliverabilityTestReport.
func (mr *MockSesv2ClientMockRecorder) GetDeliverabilityTestReport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeliverabilityTestReport", reflect.TypeOf((*MockSesv2Client)(nil).GetDeliverabilityTestReport), varargs...)
}

// GetDomainDeliverabilityCampaign mocks base method.
func (m *MockSesv2Client) GetDomainDeliverabilityCampaign(arg0 context.Context, arg1 *sesv2.GetDomainDeliverabilityCampaignInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetDomainDeliverabilityCampaignOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDomainDeliverabilityCampaign")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDomainDeliverabilityCampaign", varargs...)
	ret0, _ := ret[0].(*sesv2.GetDomainDeliverabilityCampaignOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomainDeliverabilityCampaign indicates an expected call of GetDomainDeliverabilityCampaign.
func (mr *MockSesv2ClientMockRecorder) GetDomainDeliverabilityCampaign(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainDeliverabilityCampaign", reflect.TypeOf((*MockSesv2Client)(nil).GetDomainDeliverabilityCampaign), varargs...)
}

// GetDomainStatisticsReport mocks base method.
func (m *MockSesv2Client) GetDomainStatisticsReport(arg0 context.Context, arg1 *sesv2.GetDomainStatisticsReportInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetDomainStatisticsReportOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDomainStatisticsReport")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDomainStatisticsReport", varargs...)
	ret0, _ := ret[0].(*sesv2.GetDomainStatisticsReportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomainStatisticsReport indicates an expected call of GetDomainStatisticsReport.
func (mr *MockSesv2ClientMockRecorder) GetDomainStatisticsReport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainStatisticsReport", reflect.TypeOf((*MockSesv2Client)(nil).GetDomainStatisticsReport), varargs...)
}

// GetEmailIdentity mocks base method.
func (m *MockSesv2Client) GetEmailIdentity(arg0 context.Context, arg1 *sesv2.GetEmailIdentityInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetEmailIdentityOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetEmailIdentity")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEmailIdentity", varargs...)
	ret0, _ := ret[0].(*sesv2.GetEmailIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmailIdentity indicates an expected call of GetEmailIdentity.
func (mr *MockSesv2ClientMockRecorder) GetEmailIdentity(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmailIdentity", reflect.TypeOf((*MockSesv2Client)(nil).GetEmailIdentity), varargs...)
}

// GetEmailIdentityPolicies mocks base method.
func (m *MockSesv2Client) GetEmailIdentityPolicies(arg0 context.Context, arg1 *sesv2.GetEmailIdentityPoliciesInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetEmailIdentityPoliciesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetEmailIdentityPolicies")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEmailIdentityPolicies", varargs...)
	ret0, _ := ret[0].(*sesv2.GetEmailIdentityPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmailIdentityPolicies indicates an expected call of GetEmailIdentityPolicies.
func (mr *MockSesv2ClientMockRecorder) GetEmailIdentityPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmailIdentityPolicies", reflect.TypeOf((*MockSesv2Client)(nil).GetEmailIdentityPolicies), varargs...)
}

// GetEmailTemplate mocks base method.
func (m *MockSesv2Client) GetEmailTemplate(arg0 context.Context, arg1 *sesv2.GetEmailTemplateInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetEmailTemplateOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetEmailTemplate")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEmailTemplate", varargs...)
	ret0, _ := ret[0].(*sesv2.GetEmailTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmailTemplate indicates an expected call of GetEmailTemplate.
func (mr *MockSesv2ClientMockRecorder) GetEmailTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmailTemplate", reflect.TypeOf((*MockSesv2Client)(nil).GetEmailTemplate), varargs...)
}

// GetImportJob mocks base method.
func (m *MockSesv2Client) GetImportJob(arg0 context.Context, arg1 *sesv2.GetImportJobInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetImportJobOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetImportJob")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetImportJob", varargs...)
	ret0, _ := ret[0].(*sesv2.GetImportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImportJob indicates an expected call of GetImportJob.
func (mr *MockSesv2ClientMockRecorder) GetImportJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImportJob", reflect.TypeOf((*MockSesv2Client)(nil).GetImportJob), varargs...)
}

// GetSuppressedDestination mocks base method.
func (m *MockSesv2Client) GetSuppressedDestination(arg0 context.Context, arg1 *sesv2.GetSuppressedDestinationInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetSuppressedDestinationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetSuppressedDestination")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSuppressedDestination", varargs...)
	ret0, _ := ret[0].(*sesv2.GetSuppressedDestinationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSuppressedDestination indicates an expected call of GetSuppressedDestination.
func (mr *MockSesv2ClientMockRecorder) GetSuppressedDestination(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSuppressedDestination", reflect.TypeOf((*MockSesv2Client)(nil).GetSuppressedDestination), varargs...)
}

// ListConfigurationSets mocks base method.
func (m *MockSesv2Client) ListConfigurationSets(arg0 context.Context, arg1 *sesv2.ListConfigurationSetsInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListConfigurationSetsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListConfigurationSets")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListConfigurationSets", varargs...)
	ret0, _ := ret[0].(*sesv2.ListConfigurationSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConfigurationSets indicates an expected call of ListConfigurationSets.
func (mr *MockSesv2ClientMockRecorder) ListConfigurationSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConfigurationSets", reflect.TypeOf((*MockSesv2Client)(nil).ListConfigurationSets), varargs...)
}

// ListContactLists mocks base method.
func (m *MockSesv2Client) ListContactLists(arg0 context.Context, arg1 *sesv2.ListContactListsInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListContactListsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListContactLists")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListContactLists", varargs...)
	ret0, _ := ret[0].(*sesv2.ListContactListsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContactLists indicates an expected call of ListContactLists.
func (mr *MockSesv2ClientMockRecorder) ListContactLists(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContactLists", reflect.TypeOf((*MockSesv2Client)(nil).ListContactLists), varargs...)
}

// ListContacts mocks base method.
func (m *MockSesv2Client) ListContacts(arg0 context.Context, arg1 *sesv2.ListContactsInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListContactsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListContacts")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListContacts", varargs...)
	ret0, _ := ret[0].(*sesv2.ListContactsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContacts indicates an expected call of ListContacts.
func (mr *MockSesv2ClientMockRecorder) ListContacts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContacts", reflect.TypeOf((*MockSesv2Client)(nil).ListContacts), varargs...)
}

// ListCustomVerificationEmailTemplates mocks base method.
func (m *MockSesv2Client) ListCustomVerificationEmailTemplates(arg0 context.Context, arg1 *sesv2.ListCustomVerificationEmailTemplatesInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListCustomVerificationEmailTemplatesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListCustomVerificationEmailTemplates")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCustomVerificationEmailTemplates", varargs...)
	ret0, _ := ret[0].(*sesv2.ListCustomVerificationEmailTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCustomVerificationEmailTemplates indicates an expected call of ListCustomVerificationEmailTemplates.
func (mr *MockSesv2ClientMockRecorder) ListCustomVerificationEmailTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCustomVerificationEmailTemplates", reflect.TypeOf((*MockSesv2Client)(nil).ListCustomVerificationEmailTemplates), varargs...)
}

// ListDedicatedIpPools mocks base method.
func (m *MockSesv2Client) ListDedicatedIpPools(arg0 context.Context, arg1 *sesv2.ListDedicatedIpPoolsInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListDedicatedIpPoolsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListDedicatedIpPools")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDedicatedIpPools", varargs...)
	ret0, _ := ret[0].(*sesv2.ListDedicatedIpPoolsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDedicatedIpPools indicates an expected call of ListDedicatedIpPools.
func (mr *MockSesv2ClientMockRecorder) ListDedicatedIpPools(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDedicatedIpPools", reflect.TypeOf((*MockSesv2Client)(nil).ListDedicatedIpPools), varargs...)
}

// ListDeliverabilityTestReports mocks base method.
func (m *MockSesv2Client) ListDeliverabilityTestReports(arg0 context.Context, arg1 *sesv2.ListDeliverabilityTestReportsInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListDeliverabilityTestReportsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListDeliverabilityTestReports")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDeliverabilityTestReports", varargs...)
	ret0, _ := ret[0].(*sesv2.ListDeliverabilityTestReportsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeliverabilityTestReports indicates an expected call of ListDeliverabilityTestReports.
func (mr *MockSesv2ClientMockRecorder) ListDeliverabilityTestReports(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeliverabilityTestReports", reflect.TypeOf((*MockSesv2Client)(nil).ListDeliverabilityTestReports), varargs...)
}

// ListDomainDeliverabilityCampaigns mocks base method.
func (m *MockSesv2Client) ListDomainDeliverabilityCampaigns(arg0 context.Context, arg1 *sesv2.ListDomainDeliverabilityCampaignsInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListDomainDeliverabilityCampaignsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListDomainDeliverabilityCampaigns")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDomainDeliverabilityCampaigns", varargs...)
	ret0, _ := ret[0].(*sesv2.ListDomainDeliverabilityCampaignsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDomainDeliverabilityCampaigns indicates an expected call of ListDomainDeliverabilityCampaigns.
func (mr *MockSesv2ClientMockRecorder) ListDomainDeliverabilityCampaigns(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDomainDeliverabilityCampaigns", reflect.TypeOf((*MockSesv2Client)(nil).ListDomainDeliverabilityCampaigns), varargs...)
}

// ListEmailIdentities mocks base method.
func (m *MockSesv2Client) ListEmailIdentities(arg0 context.Context, arg1 *sesv2.ListEmailIdentitiesInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListEmailIdentitiesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListEmailIdentities")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEmailIdentities", varargs...)
	ret0, _ := ret[0].(*sesv2.ListEmailIdentitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEmailIdentities indicates an expected call of ListEmailIdentities.
func (mr *MockSesv2ClientMockRecorder) ListEmailIdentities(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEmailIdentities", reflect.TypeOf((*MockSesv2Client)(nil).ListEmailIdentities), varargs...)
}

// ListEmailTemplates mocks base method.
func (m *MockSesv2Client) ListEmailTemplates(arg0 context.Context, arg1 *sesv2.ListEmailTemplatesInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListEmailTemplatesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListEmailTemplates")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEmailTemplates", varargs...)
	ret0, _ := ret[0].(*sesv2.ListEmailTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEmailTemplates indicates an expected call of ListEmailTemplates.
func (mr *MockSesv2ClientMockRecorder) ListEmailTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEmailTemplates", reflect.TypeOf((*MockSesv2Client)(nil).ListEmailTemplates), varargs...)
}

// ListImportJobs mocks base method.
func (m *MockSesv2Client) ListImportJobs(arg0 context.Context, arg1 *sesv2.ListImportJobsInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListImportJobsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListImportJobs")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListImportJobs", varargs...)
	ret0, _ := ret[0].(*sesv2.ListImportJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImportJobs indicates an expected call of ListImportJobs.
func (mr *MockSesv2ClientMockRecorder) ListImportJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImportJobs", reflect.TypeOf((*MockSesv2Client)(nil).ListImportJobs), varargs...)
}

// ListRecommendations mocks base method.
func (m *MockSesv2Client) ListRecommendations(arg0 context.Context, arg1 *sesv2.ListRecommendationsInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListRecommendationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListRecommendations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRecommendations", varargs...)
	ret0, _ := ret[0].(*sesv2.ListRecommendationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRecommendations indicates an expected call of ListRecommendations.
func (mr *MockSesv2ClientMockRecorder) ListRecommendations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRecommendations", reflect.TypeOf((*MockSesv2Client)(nil).ListRecommendations), varargs...)
}

// ListSuppressedDestinations mocks base method.
func (m *MockSesv2Client) ListSuppressedDestinations(arg0 context.Context, arg1 *sesv2.ListSuppressedDestinationsInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListSuppressedDestinationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListSuppressedDestinations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSuppressedDestinations", varargs...)
	ret0, _ := ret[0].(*sesv2.ListSuppressedDestinationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSuppressedDestinations indicates an expected call of ListSuppressedDestinations.
func (mr *MockSesv2ClientMockRecorder) ListSuppressedDestinations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSuppressedDestinations", reflect.TypeOf((*MockSesv2Client)(nil).ListSuppressedDestinations), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockSesv2Client) ListTagsForResource(arg0 context.Context, arg1 *sesv2.ListTagsForResourceInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &sesv2.Options{}
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
	ret0, _ := ret[0].(*sesv2.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockSesv2ClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockSesv2Client)(nil).ListTagsForResource), varargs...)
}
