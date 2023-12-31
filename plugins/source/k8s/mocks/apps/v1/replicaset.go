// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/apps/v1 (interfaces: ReplicaSetsGetter,ReplicaSetInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/api/autoscaling/v1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v12 "k8s.io/client-go/applyconfigurations/apps/v1"
	v13 "k8s.io/client-go/applyconfigurations/autoscaling/v1"
	v14 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

// MockReplicaSetsGetter is a mock of ReplicaSetsGetter interface.
type MockReplicaSetsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockReplicaSetsGetterMockRecorder
}

// MockReplicaSetsGetterMockRecorder is the mock recorder for MockReplicaSetsGetter.
type MockReplicaSetsGetterMockRecorder struct {
	mock *MockReplicaSetsGetter
}

// NewMockReplicaSetsGetter creates a new mock instance.
func NewMockReplicaSetsGetter(ctrl *gomock.Controller) *MockReplicaSetsGetter {
	mock := &MockReplicaSetsGetter{ctrl: ctrl}
	mock.recorder = &MockReplicaSetsGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReplicaSetsGetter) EXPECT() *MockReplicaSetsGetterMockRecorder {
	return m.recorder
}

// ReplicaSets mocks base method.
func (m *MockReplicaSetsGetter) ReplicaSets(arg0 string) v14.ReplicaSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplicaSets", arg0)
	ret0, _ := ret[0].(v14.ReplicaSetInterface)
	return ret0
}

// ReplicaSets indicates an expected call of ReplicaSets.
func (mr *MockReplicaSetsGetterMockRecorder) ReplicaSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicaSets", reflect.TypeOf((*MockReplicaSetsGetter)(nil).ReplicaSets), arg0)
}

// MockReplicaSetInterface is a mock of ReplicaSetInterface interface.
type MockReplicaSetInterface struct {
	ctrl     *gomock.Controller
	recorder *MockReplicaSetInterfaceMockRecorder
}

// MockReplicaSetInterfaceMockRecorder is the mock recorder for MockReplicaSetInterface.
type MockReplicaSetInterfaceMockRecorder struct {
	mock *MockReplicaSetInterface
}

// NewMockReplicaSetInterface creates a new mock instance.
func NewMockReplicaSetInterface(ctrl *gomock.Controller) *MockReplicaSetInterface {
	mock := &MockReplicaSetInterface{ctrl: ctrl}
	mock.recorder = &MockReplicaSetInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReplicaSetInterface) EXPECT() *MockReplicaSetInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockReplicaSetInterface) Apply(arg0 context.Context, arg1 *v12.ReplicaSetApplyConfiguration, arg2 v11.ApplyOptions) (*v1.ReplicaSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockReplicaSetInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockReplicaSetInterface)(nil).Apply), arg0, arg1, arg2)
}

// ApplyScale mocks base method.
func (m *MockReplicaSetInterface) ApplyScale(arg0 context.Context, arg1 string, arg2 *v13.ScaleApplyConfiguration, arg3 v11.ApplyOptions) (*v10.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyScale", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v10.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyScale indicates an expected call of ApplyScale.
func (mr *MockReplicaSetInterfaceMockRecorder) ApplyScale(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyScale", reflect.TypeOf((*MockReplicaSetInterface)(nil).ApplyScale), arg0, arg1, arg2, arg3)
}

// ApplyStatus mocks base method.
func (m *MockReplicaSetInterface) ApplyStatus(arg0 context.Context, arg1 *v12.ReplicaSetApplyConfiguration, arg2 v11.ApplyOptions) (*v1.ReplicaSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyStatus indicates an expected call of ApplyStatus.
func (mr *MockReplicaSetInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyStatus", reflect.TypeOf((*MockReplicaSetInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockReplicaSetInterface) Create(arg0 context.Context, arg1 *v1.ReplicaSet, arg2 v11.CreateOptions) (*v1.ReplicaSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockReplicaSetInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockReplicaSetInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockReplicaSetInterface) Delete(arg0 context.Context, arg1 string, arg2 v11.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockReplicaSetInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockReplicaSetInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockReplicaSetInterface) DeleteCollection(arg0 context.Context, arg1 v11.DeleteOptions, arg2 v11.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockReplicaSetInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockReplicaSetInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockReplicaSetInterface) Get(arg0 context.Context, arg1 string, arg2 v11.GetOptions) (*v1.ReplicaSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockReplicaSetInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockReplicaSetInterface)(nil).Get), arg0, arg1, arg2)
}

// GetScale mocks base method.
func (m *MockReplicaSetInterface) GetScale(arg0 context.Context, arg1 string, arg2 v11.GetOptions) (*v10.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScale", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScale indicates an expected call of GetScale.
func (mr *MockReplicaSetInterfaceMockRecorder) GetScale(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScale", reflect.TypeOf((*MockReplicaSetInterface)(nil).GetScale), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockReplicaSetInterface) List(arg0 context.Context, arg1 v11.ListOptions) (*v1.ReplicaSetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.ReplicaSetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockReplicaSetInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockReplicaSetInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockReplicaSetInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v11.PatchOptions, arg5 ...string) (*v1.ReplicaSet, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.ReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockReplicaSetInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockReplicaSetInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockReplicaSetInterface) Update(arg0 context.Context, arg1 *v1.ReplicaSet, arg2 v11.UpdateOptions) (*v1.ReplicaSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockReplicaSetInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockReplicaSetInterface)(nil).Update), arg0, arg1, arg2)
}

// UpdateScale mocks base method.
func (m *MockReplicaSetInterface) UpdateScale(arg0 context.Context, arg1 string, arg2 *v10.Scale, arg3 v11.UpdateOptions) (*v10.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScale", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v10.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateScale indicates an expected call of UpdateScale.
func (mr *MockReplicaSetInterfaceMockRecorder) UpdateScale(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScale", reflect.TypeOf((*MockReplicaSetInterface)(nil).UpdateScale), arg0, arg1, arg2, arg3)
}

// UpdateStatus mocks base method.
func (m *MockReplicaSetInterface) UpdateStatus(arg0 context.Context, arg1 *v1.ReplicaSet, arg2 v11.UpdateOptions) (*v1.ReplicaSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockReplicaSetInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockReplicaSetInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockReplicaSetInterface) Watch(arg0 context.Context, arg1 v11.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockReplicaSetInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockReplicaSetInterface)(nil).Watch), arg0, arg1)
}
