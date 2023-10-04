// Code generated by MockGen. DO NOT EDIT.
// Source: cognitoidentity.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	cognitoidentity "github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	gomock "github.com/golang/mock/gomock"
)

// MockCognitoidentityClient is a mock of CognitoidentityClient interface.
type MockCognitoidentityClient struct {
	ctrl     *gomock.Controller
	recorder *MockCognitoidentityClientMockRecorder
}

// MockCognitoidentityClientMockRecorder is the mock recorder for MockCognitoidentityClient.
type MockCognitoidentityClientMockRecorder struct {
	mock *MockCognitoidentityClient
}

// NewMockCognitoidentityClient creates a new mock instance.
func NewMockCognitoidentityClient(ctrl *gomock.Controller) *MockCognitoidentityClient {
	mock := &MockCognitoidentityClient{ctrl: ctrl}
	mock.recorder = &MockCognitoidentityClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCognitoidentityClient) EXPECT() *MockCognitoidentityClientMockRecorder {
	return m.recorder
}

// DescribeIdentity mocks base method.
func (m *MockCognitoidentityClient) DescribeIdentity(arg0 context.Context, arg1 *cognitoidentity.DescribeIdentityInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.DescribeIdentityOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentity.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeIdentity")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeIdentity", varargs...)
	ret0, _ := ret[0].(*cognitoidentity.DescribeIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeIdentity indicates an expected call of DescribeIdentity.
func (mr *MockCognitoidentityClientMockRecorder) DescribeIdentity(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeIdentity", reflect.TypeOf((*MockCognitoidentityClient)(nil).DescribeIdentity), varargs...)
}

// DescribeIdentityPool mocks base method.
func (m *MockCognitoidentityClient) DescribeIdentityPool(arg0 context.Context, arg1 *cognitoidentity.DescribeIdentityPoolInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.DescribeIdentityPoolOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentity.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeIdentityPool")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeIdentityPool", varargs...)
	ret0, _ := ret[0].(*cognitoidentity.DescribeIdentityPoolOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeIdentityPool indicates an expected call of DescribeIdentityPool.
func (mr *MockCognitoidentityClientMockRecorder) DescribeIdentityPool(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeIdentityPool", reflect.TypeOf((*MockCognitoidentityClient)(nil).DescribeIdentityPool), varargs...)
}

// GetCredentialsForIdentity mocks base method.
func (m *MockCognitoidentityClient) GetCredentialsForIdentity(arg0 context.Context, arg1 *cognitoidentity.GetCredentialsForIdentityInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.GetCredentialsForIdentityOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentity.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetCredentialsForIdentity")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCredentialsForIdentity", varargs...)
	ret0, _ := ret[0].(*cognitoidentity.GetCredentialsForIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentialsForIdentity indicates an expected call of GetCredentialsForIdentity.
func (mr *MockCognitoidentityClientMockRecorder) GetCredentialsForIdentity(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentialsForIdentity", reflect.TypeOf((*MockCognitoidentityClient)(nil).GetCredentialsForIdentity), varargs...)
}

// GetId mocks base method.
func (m *MockCognitoidentityClient) GetId(arg0 context.Context, arg1 *cognitoidentity.GetIdInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.GetIdOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentity.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetId")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetId", varargs...)
	ret0, _ := ret[0].(*cognitoidentity.GetIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetId indicates an expected call of GetId.
func (mr *MockCognitoidentityClientMockRecorder) GetId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetId", reflect.TypeOf((*MockCognitoidentityClient)(nil).GetId), varargs...)
}

// GetIdentityPoolRoles mocks base method.
func (m *MockCognitoidentityClient) GetIdentityPoolRoles(arg0 context.Context, arg1 *cognitoidentity.GetIdentityPoolRolesInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.GetIdentityPoolRolesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentity.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetIdentityPoolRoles")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIdentityPoolRoles", varargs...)
	ret0, _ := ret[0].(*cognitoidentity.GetIdentityPoolRolesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIdentityPoolRoles indicates an expected call of GetIdentityPoolRoles.
func (mr *MockCognitoidentityClientMockRecorder) GetIdentityPoolRoles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdentityPoolRoles", reflect.TypeOf((*MockCognitoidentityClient)(nil).GetIdentityPoolRoles), varargs...)
}

// GetOpenIdToken mocks base method.
func (m *MockCognitoidentityClient) GetOpenIdToken(arg0 context.Context, arg1 *cognitoidentity.GetOpenIdTokenInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.GetOpenIdTokenOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentity.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetOpenIdToken")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOpenIdToken", varargs...)
	ret0, _ := ret[0].(*cognitoidentity.GetOpenIdTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOpenIdToken indicates an expected call of GetOpenIdToken.
func (mr *MockCognitoidentityClientMockRecorder) GetOpenIdToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenIdToken", reflect.TypeOf((*MockCognitoidentityClient)(nil).GetOpenIdToken), varargs...)
}

// GetOpenIdTokenForDeveloperIdentity mocks base method.
func (m *MockCognitoidentityClient) GetOpenIdTokenForDeveloperIdentity(arg0 context.Context, arg1 *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentity.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetOpenIdTokenForDeveloperIdentity")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOpenIdTokenForDeveloperIdentity", varargs...)
	ret0, _ := ret[0].(*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOpenIdTokenForDeveloperIdentity indicates an expected call of GetOpenIdTokenForDeveloperIdentity.
func (mr *MockCognitoidentityClientMockRecorder) GetOpenIdTokenForDeveloperIdentity(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenIdTokenForDeveloperIdentity", reflect.TypeOf((*MockCognitoidentityClient)(nil).GetOpenIdTokenForDeveloperIdentity), varargs...)
}

// GetPrincipalTagAttributeMap mocks base method.
func (m *MockCognitoidentityClient) GetPrincipalTagAttributeMap(arg0 context.Context, arg1 *cognitoidentity.GetPrincipalTagAttributeMapInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.GetPrincipalTagAttributeMapOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentity.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetPrincipalTagAttributeMap")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPrincipalTagAttributeMap", varargs...)
	ret0, _ := ret[0].(*cognitoidentity.GetPrincipalTagAttributeMapOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrincipalTagAttributeMap indicates an expected call of GetPrincipalTagAttributeMap.
func (mr *MockCognitoidentityClientMockRecorder) GetPrincipalTagAttributeMap(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrincipalTagAttributeMap", reflect.TypeOf((*MockCognitoidentityClient)(nil).GetPrincipalTagAttributeMap), varargs...)
}

// ListIdentities mocks base method.
func (m *MockCognitoidentityClient) ListIdentities(arg0 context.Context, arg1 *cognitoidentity.ListIdentitiesInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.ListIdentitiesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentity.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListIdentities")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIdentities", varargs...)
	ret0, _ := ret[0].(*cognitoidentity.ListIdentitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIdentities indicates an expected call of ListIdentities.
func (mr *MockCognitoidentityClientMockRecorder) ListIdentities(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIdentities", reflect.TypeOf((*MockCognitoidentityClient)(nil).ListIdentities), varargs...)
}

// ListIdentityPools mocks base method.
func (m *MockCognitoidentityClient) ListIdentityPools(arg0 context.Context, arg1 *cognitoidentity.ListIdentityPoolsInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.ListIdentityPoolsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentity.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListIdentityPools")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIdentityPools", varargs...)
	ret0, _ := ret[0].(*cognitoidentity.ListIdentityPoolsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIdentityPools indicates an expected call of ListIdentityPools.
func (mr *MockCognitoidentityClientMockRecorder) ListIdentityPools(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIdentityPools", reflect.TypeOf((*MockCognitoidentityClient)(nil).ListIdentityPools), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockCognitoidentityClient) ListTagsForResource(arg0 context.Context, arg1 *cognitoidentity.ListTagsForResourceInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cognitoidentity.Options{}
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
	ret0, _ := ret[0].(*cognitoidentity.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockCognitoidentityClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockCognitoidentityClient)(nil).ListTagsForResource), varargs...)
}