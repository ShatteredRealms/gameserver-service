package service

import (
	"context"

	"github.com/ShatteredRealms/dimension-service/pkg/model/game"
	"github.com/ShatteredRealms/dimension-service/pkg/repository"
)

type MapService interface {
	GetMaps(ctx context.Context) (*game.Maps, error)
	GetMapById(ctx context.Context, mapId string) (*game.Map, error)
	CreateMap(ctx context.Context, mapId string) (*game.Map, error)
	DeleteMap(ctx context.Context, mapId string) (*game.Map, error)
}

type mapService struct {
	repo repository.MapRepository
}

func NewMapService(repo repository.MapRepository) MapService {
	return &mapService{repo: repo}
}

// CreateMap implements MapService.
func (d *mapService) CreateMap(ctx context.Context, mapId string) (*game.Map, error) {
	return d.repo.CreateMap(ctx, mapId)
}

// DeleteMap implements MapService.
func (d *mapService) DeleteMap(ctx context.Context, mapId string) (*game.Map, error) {
	return d.repo.DeleteMap(ctx, mapId)
}

// GetMapById implements MapService.
func (d *mapService) GetMapById(ctx context.Context, mapId string) (*game.Map, error) {
	return d.repo.GetMapById(ctx, mapId)
}

// GetMaps implements MapService.
func (d *mapService) GetMaps(ctx context.Context) (*game.Maps, error) {
	return d.repo.GetMaps(ctx)
}
