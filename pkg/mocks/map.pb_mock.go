// Code generated by MockGen. DO NOT EDIT.
// Source: /home/wil/sro/gameserver-service/pkg/pb/map.pb.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=/home/wil/sro/gameserver-service/pkg/pb/map.pb.go -destination=/home/wil/sro/gameserver-service/pkg/mocks/map.pb_mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockisEditMapRequest_OptionalName is a mock of isEditMapRequest_OptionalName interface.
type MockisEditMapRequest_OptionalName struct {
	ctrl     *gomock.Controller
	recorder *MockisEditMapRequest_OptionalNameMockRecorder
	isgomock struct{}
}

// MockisEditMapRequest_OptionalNameMockRecorder is the mock recorder for MockisEditMapRequest_OptionalName.
type MockisEditMapRequest_OptionalNameMockRecorder struct {
	mock *MockisEditMapRequest_OptionalName
}

// NewMockisEditMapRequest_OptionalName creates a new mock instance.
func NewMockisEditMapRequest_OptionalName(ctrl *gomock.Controller) *MockisEditMapRequest_OptionalName {
	mock := &MockisEditMapRequest_OptionalName{ctrl: ctrl}
	mock.recorder = &MockisEditMapRequest_OptionalNameMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisEditMapRequest_OptionalName) EXPECT() *MockisEditMapRequest_OptionalNameMockRecorder {
	return m.recorder
}

// isEditMapRequest_OptionalName mocks base method.
func (m *MockisEditMapRequest_OptionalName) isEditMapRequest_OptionalName() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isEditMapRequest_OptionalName")
}

// isEditMapRequest_OptionalName indicates an expected call of isEditMapRequest_OptionalName.
func (mr *MockisEditMapRequest_OptionalNameMockRecorder) isEditMapRequest_OptionalName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isEditMapRequest_OptionalName", reflect.TypeOf((*MockisEditMapRequest_OptionalName)(nil).isEditMapRequest_OptionalName))
}

// MockisEditMapRequest_OptionalMapPath is a mock of isEditMapRequest_OptionalMapPath interface.
type MockisEditMapRequest_OptionalMapPath struct {
	ctrl     *gomock.Controller
	recorder *MockisEditMapRequest_OptionalMapPathMockRecorder
	isgomock struct{}
}

// MockisEditMapRequest_OptionalMapPathMockRecorder is the mock recorder for MockisEditMapRequest_OptionalMapPath.
type MockisEditMapRequest_OptionalMapPathMockRecorder struct {
	mock *MockisEditMapRequest_OptionalMapPath
}

// NewMockisEditMapRequest_OptionalMapPath creates a new mock instance.
func NewMockisEditMapRequest_OptionalMapPath(ctrl *gomock.Controller) *MockisEditMapRequest_OptionalMapPath {
	mock := &MockisEditMapRequest_OptionalMapPath{ctrl: ctrl}
	mock.recorder = &MockisEditMapRequest_OptionalMapPathMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisEditMapRequest_OptionalMapPath) EXPECT() *MockisEditMapRequest_OptionalMapPathMockRecorder {
	return m.recorder
}

// isEditMapRequest_OptionalMapPath mocks base method.
func (m *MockisEditMapRequest_OptionalMapPath) isEditMapRequest_OptionalMapPath() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isEditMapRequest_OptionalMapPath")
}

// isEditMapRequest_OptionalMapPath indicates an expected call of isEditMapRequest_OptionalMapPath.
func (mr *MockisEditMapRequest_OptionalMapPathMockRecorder) isEditMapRequest_OptionalMapPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isEditMapRequest_OptionalMapPath", reflect.TypeOf((*MockisEditMapRequest_OptionalMapPath)(nil).isEditMapRequest_OptionalMapPath))
}
