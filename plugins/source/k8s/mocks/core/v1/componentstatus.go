// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/core/v1 (interfaces: ComponentStatusesGetter,ComponentStatusInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/core/v1"
	v12 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// MockComponentStatusesGetter is a mock of ComponentStatusesGetter interface.
type MockComponentStatusesGetter struct {
	ctrl     *gomock.Controller
	recorder *MockComponentStatusesGetterMockRecorder
}

// MockComponentStatusesGetterMockRecorder is the mock recorder for MockComponentStatusesGetter.
type MockComponentStatusesGetterMockRecorder struct {
	mock *MockComponentStatusesGetter
}

// NewMockComponentStatusesGetter creates a new mock instance.
func NewMockComponentStatusesGetter(ctrl *gomock.Controller) *MockComponentStatusesGetter {
	mock := &MockComponentStatusesGetter{ctrl: ctrl}
	mock.recorder = &MockComponentStatusesGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComponentStatusesGetter) EXPECT() *MockComponentStatusesGetterMockRecorder {
	return m.recorder
}

// ComponentStatuses mocks base method.
func (m *MockComponentStatusesGetter) ComponentStatuses() v12.ComponentStatusInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComponentStatuses")
	ret0, _ := ret[0].(v12.ComponentStatusInterface)
	return ret0
}

// ComponentStatuses indicates an expected call of ComponentStatuses.
func (mr *MockComponentStatusesGetterMockRecorder) ComponentStatuses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComponentStatuses", reflect.TypeOf((*MockComponentStatusesGetter)(nil).ComponentStatuses))
}

// MockComponentStatusInterface is a mock of ComponentStatusInterface interface.
type MockComponentStatusInterface struct {
	ctrl     *gomock.Controller
	recorder *MockComponentStatusInterfaceMockRecorder
}

// MockComponentStatusInterfaceMockRecorder is the mock recorder for MockComponentStatusInterface.
type MockComponentStatusInterfaceMockRecorder struct {
	mock *MockComponentStatusInterface
}

// NewMockComponentStatusInterface creates a new mock instance.
func NewMockComponentStatusInterface(ctrl *gomock.Controller) *MockComponentStatusInterface {
	mock := &MockComponentStatusInterface{ctrl: ctrl}
	mock.recorder = &MockComponentStatusInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComponentStatusInterface) EXPECT() *MockComponentStatusInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockComponentStatusInterface) Apply(arg0 context.Context, arg1 *v11.ComponentStatusApplyConfiguration, arg2 v10.ApplyOptions) (*v1.ComponentStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ComponentStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockComponentStatusInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockComponentStatusInterface)(nil).Apply), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockComponentStatusInterface) Create(arg0 context.Context, arg1 *v1.ComponentStatus, arg2 v10.CreateOptions) (*v1.ComponentStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ComponentStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockComponentStatusInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockComponentStatusInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockComponentStatusInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockComponentStatusInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockComponentStatusInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockComponentStatusInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockComponentStatusInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockComponentStatusInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockComponentStatusInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.ComponentStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ComponentStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockComponentStatusInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockComponentStatusInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockComponentStatusInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.ComponentStatusList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.ComponentStatusList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockComponentStatusInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockComponentStatusInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockComponentStatusInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.ComponentStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.ComponentStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockComponentStatusInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockComponentStatusInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockComponentStatusInterface) Update(arg0 context.Context, arg1 *v1.ComponentStatus, arg2 v10.UpdateOptions) (*v1.ComponentStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ComponentStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockComponentStatusInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockComponentStatusInterface)(nil).Update), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockComponentStatusInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockComponentStatusInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockComponentStatusInterface)(nil).Watch), arg0, arg1)
}
