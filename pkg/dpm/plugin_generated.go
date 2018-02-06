// Code generated by MockGen. DO NOT EDIT.
// Source: plugin.go

// Package dpm is a generated GoMock package.
package dpm

import (
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	v1alpha "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1alpha"
	reflect "reflect"
)

// MockDevicePluginInterface is a mock of DevicePluginInterface interface
type MockDevicePluginInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDevicePluginInterfaceMockRecorder
}

// MockDevicePluginInterfaceMockRecorder is the mock recorder for MockDevicePluginInterface
type MockDevicePluginInterfaceMockRecorder struct {
	mock *MockDevicePluginInterface
}

// NewMockDevicePluginInterface creates a new mock instance
func NewMockDevicePluginInterface(ctrl *gomock.Controller) *MockDevicePluginInterface {
	mock := &MockDevicePluginInterface{ctrl: ctrl}
	mock.recorder = &MockDevicePluginInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDevicePluginInterface) EXPECT() *MockDevicePluginInterfaceMockRecorder {
	return m.recorder
}

// ListAndWatch mocks base method
func (m *MockDevicePluginInterface) ListAndWatch(arg0 *v1alpha.Empty, arg1 v1alpha.DevicePlugin_ListAndWatchServer) error {
	ret := m.ctrl.Call(m, "ListAndWatch", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListAndWatch indicates an expected call of ListAndWatch
func (mr *MockDevicePluginInterfaceMockRecorder) ListAndWatch(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAndWatch", reflect.TypeOf((*MockDevicePluginInterface)(nil).ListAndWatch), arg0, arg1)
}

// Allocate mocks base method
func (m *MockDevicePluginInterface) Allocate(arg0 context.Context, arg1 *v1alpha.AllocateRequest) (*v1alpha.AllocateResponse, error) {
	ret := m.ctrl.Call(m, "Allocate", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha.AllocateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Allocate indicates an expected call of Allocate
func (mr *MockDevicePluginInterfaceMockRecorder) Allocate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Allocate", reflect.TypeOf((*MockDevicePluginInterface)(nil).Allocate), arg0, arg1)
}

// Start mocks base method
func (m *MockDevicePluginInterface) Start() error {
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockDevicePluginInterfaceMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockDevicePluginInterface)(nil).Start))
}

// Stop mocks base method
func (m *MockDevicePluginInterface) Stop() error {
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockDevicePluginInterfaceMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockDevicePluginInterface)(nil).Stop))
}
