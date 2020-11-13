// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/octant/internal/octant (interfaces: PodMetricsCRUD)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// MockPodMetricsCRUD is a mock of PodMetricsCRUD interface
type MockPodMetricsCRUD struct {
	ctrl     *gomock.Controller
	recorder *MockPodMetricsCRUDMockRecorder
}

// MockPodMetricsCRUDMockRecorder is the mock recorder for MockPodMetricsCRUD
type MockPodMetricsCRUDMockRecorder struct {
	mock *MockPodMetricsCRUD
}

// NewMockPodMetricsCRUD creates a new mock instance
func NewMockPodMetricsCRUD(ctrl *gomock.Controller) *MockPodMetricsCRUD {
	mock := &MockPodMetricsCRUD{ctrl: ctrl}
	mock.recorder = &MockPodMetricsCRUDMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPodMetricsCRUD) EXPECT() *MockPodMetricsCRUDMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockPodMetricsCRUD) Get(arg0 context.Context, arg1, arg2 string) (*unstructured.Unstructured, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get
func (mr *MockPodMetricsCRUDMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPodMetricsCRUD)(nil).Get), arg0, arg1, arg2)
}
