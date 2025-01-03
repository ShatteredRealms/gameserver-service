// Code generated by MockGen. DO NOT EDIT.
// Source: /home/wil/sro/gameserver-service/pkg/service/map.go
//
// Generated by this command:
//
//	mockgen -source=/home/wil/sro/gameserver-service/pkg/service/map.go -destination=/home/wil/sro/gameserver-service/pkg/service/mocks/map.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	game "github.com/ShatteredRealms/gameserver-service/pkg/model/game"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockMapService is a mock of MapService interface.
type MockMapService struct {
	ctrl     *gomock.Controller
	recorder *MockMapServiceMockRecorder
	isgomock struct{}
}

// MockMapServiceMockRecorder is the mock recorder for MockMapService.
type MockMapServiceMockRecorder struct {
	mock *MockMapService
}

// NewMockMapService creates a new mock instance.
func NewMockMapService(ctrl *gomock.Controller) *MockMapService {
	mock := &MockMapService{ctrl: ctrl}
	mock.recorder = &MockMapServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMapService) EXPECT() *MockMapServiceMockRecorder {
	return m.recorder
}

// CreateMap mocks base method.
func (m *MockMapService) CreateMap(ctx context.Context, name, mapPath string) (*game.Map, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMap", ctx, name, mapPath)
	ret0, _ := ret[0].(*game.Map)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMap indicates an expected call of CreateMap.
func (mr *MockMapServiceMockRecorder) CreateMap(ctx, name, mapPath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMap", reflect.TypeOf((*MockMapService)(nil).CreateMap), ctx, name, mapPath)
}

// DeleteMap mocks base method.
func (m *MockMapService) DeleteMap(ctx context.Context, mapId *uuid.UUID) (*game.Map, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMap", ctx, mapId)
	ret0, _ := ret[0].(*game.Map)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMap indicates an expected call of DeleteMap.
func (mr *MockMapServiceMockRecorder) DeleteMap(ctx, mapId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMap", reflect.TypeOf((*MockMapService)(nil).DeleteMap), ctx, mapId)
}

// EditMap mocks base method.
func (m_2 *MockMapService) EditMap(ctx context.Context, m *game.Map) (*game.Map, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "EditMap", ctx, m)
	ret0, _ := ret[0].(*game.Map)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditMap indicates an expected call of EditMap.
func (mr *MockMapServiceMockRecorder) EditMap(ctx, m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditMap", reflect.TypeOf((*MockMapService)(nil).EditMap), ctx, m)
}

// GetDeletedMaps mocks base method.
func (m *MockMapService) GetDeletedMaps(ctx context.Context) (game.Maps, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeletedMaps", ctx)
	ret0, _ := ret[0].(game.Maps)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeletedMaps indicates an expected call of GetDeletedMaps.
func (mr *MockMapServiceMockRecorder) GetDeletedMaps(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeletedMaps", reflect.TypeOf((*MockMapService)(nil).GetDeletedMaps), ctx)
}

// GetMapById mocks base method.
func (m *MockMapService) GetMapById(ctx context.Context, mapId *uuid.UUID) (*game.Map, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMapById", ctx, mapId)
	ret0, _ := ret[0].(*game.Map)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMapById indicates an expected call of GetMapById.
func (mr *MockMapServiceMockRecorder) GetMapById(ctx, mapId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMapById", reflect.TypeOf((*MockMapService)(nil).GetMapById), ctx, mapId)
}

// GetMaps mocks base method.
func (m *MockMapService) GetMaps(ctx context.Context) (game.Maps, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaps", ctx)
	ret0, _ := ret[0].(game.Maps)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaps indicates an expected call of GetMaps.
func (mr *MockMapServiceMockRecorder) GetMaps(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaps", reflect.TypeOf((*MockMapService)(nil).GetMaps), ctx)
}
