// Code generated by MockGen. DO NOT EDIT.
// Source: /home/wil/dev/sro/gamesever-service/pkg/repository/map.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=/home/wil/dev/sro/gamesever-service/pkg/repository/map.go -destination=/home/wil/dev/sro/gamesever-service/pkg/mocks/map_mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	game "github.com/ShatteredRealms/gameserver-service/pkg/model/game"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockMapRepository is a mock of MapRepository interface.
type MockMapRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMapRepositoryMockRecorder
	isgomock struct{}
}

// MockMapRepositoryMockRecorder is the mock recorder for MockMapRepository.
type MockMapRepositoryMockRecorder struct {
	mock *MockMapRepository
}

// NewMockMapRepository creates a new mock instance.
func NewMockMapRepository(ctrl *gomock.Controller) *MockMapRepository {
	mock := &MockMapRepository{ctrl: ctrl}
	mock.recorder = &MockMapRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMapRepository) EXPECT() *MockMapRepositoryMockRecorder {
	return m.recorder
}

// CreateMap mocks base method.
func (m_2 *MockMapRepository) CreateMap(ctx context.Context, m *game.Map) (*game.Map, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "CreateMap", ctx, m)
	ret0, _ := ret[0].(*game.Map)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMap indicates an expected call of CreateMap.
func (mr *MockMapRepositoryMockRecorder) CreateMap(ctx, m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMap", reflect.TypeOf((*MockMapRepository)(nil).CreateMap), ctx, m)
}

// DeleteMap mocks base method.
func (m *MockMapRepository) DeleteMap(ctx context.Context, mapId *uuid.UUID) (*game.Map, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMap", ctx, mapId)
	ret0, _ := ret[0].(*game.Map)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMap indicates an expected call of DeleteMap.
func (mr *MockMapRepositoryMockRecorder) DeleteMap(ctx, mapId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMap", reflect.TypeOf((*MockMapRepository)(nil).DeleteMap), ctx, mapId)
}

// GetMapById mocks base method.
func (m *MockMapRepository) GetMapById(ctx context.Context, mapId *uuid.UUID) (*game.Map, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMapById", ctx, mapId)
	ret0, _ := ret[0].(*game.Map)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMapById indicates an expected call of GetMapById.
func (mr *MockMapRepositoryMockRecorder) GetMapById(ctx, mapId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMapById", reflect.TypeOf((*MockMapRepository)(nil).GetMapById), ctx, mapId)
}

// GetMaps mocks base method.
func (m *MockMapRepository) GetMaps(ctx context.Context) (*game.Maps, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaps", ctx)
	ret0, _ := ret[0].(*game.Maps)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaps indicates an expected call of GetMaps.
func (mr *MockMapRepositoryMockRecorder) GetMaps(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaps", reflect.TypeOf((*MockMapRepository)(nil).GetMaps), ctx)
}

// UpdateMap mocks base method.
func (m_2 *MockMapRepository) UpdateMap(ctx context.Context, m *game.Map) (*game.Map, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "UpdateMap", ctx, m)
	ret0, _ := ret[0].(*game.Map)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMap indicates an expected call of UpdateMap.
func (mr *MockMapRepositoryMockRecorder) UpdateMap(ctx, m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMap", reflect.TypeOf((*MockMapRepository)(nil).UpdateMap), ctx, m)
}
