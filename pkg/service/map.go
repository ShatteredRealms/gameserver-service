package service

import (
	"context"

	"github.com/ShatteredRealms/gameserver-service/pkg/model/game"
	"github.com/ShatteredRealms/gameserver-service/pkg/repository"
	"github.com/google/uuid"
)

type MapService interface {
	GetMaps(ctx context.Context) (*game.Maps, error)
	GetMapById(ctx context.Context, mapId *uuid.UUID) (*game.Map, error)
	CreateMap(ctx context.Context, name, mapPath string) (*game.Map, error)
	DeleteMap(ctx context.Context, mapId *uuid.UUID) (*game.Map, error)
	EditMap(ctx context.Context, m *game.Map) (*game.Map, error)
}

type mapService struct {
	repo repository.MapRepository
}

// EditMap implements MapService.
func (s *mapService) EditMap(ctx context.Context, m *game.Map) (*game.Map, error) {
	return s.repo.UpdateMap(ctx, m)
}

// CreateMap implements MapService.
func (s *mapService) CreateMap(ctx context.Context, name string, mapPath string) (*game.Map, error) {
	m := &game.Map{
		Name:    name,
		MapPath: mapPath,
	}

	return s.repo.CreateMap(ctx, m)
}

// DeleteMap implements MapService.
func (s *mapService) DeleteMap(ctx context.Context, mapId *uuid.UUID) (*game.Map, error) {
	return s.repo.DeleteMap(ctx, mapId)
}

// GetMapById implements MapService.
func (s *mapService) GetMapById(ctx context.Context, mapId *uuid.UUID) (*game.Map, error) {
	return s.repo.GetMapById(ctx, mapId)
}

// GetMaps implements MapService.
func (s *mapService) GetMaps(ctx context.Context) (*game.Maps, error) {
	return s.repo.GetMaps(ctx)
}

func NewMapService(repo repository.MapRepository) MapService {
	return &mapService{repo: repo}
}
