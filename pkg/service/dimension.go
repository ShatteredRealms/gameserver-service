package service

import (
	"context"

	"github.com/ShatteredRealms/gameserver-service/pkg/model/game"
	"github.com/ShatteredRealms/gameserver-service/pkg/repository"
	"github.com/google/uuid"
)

type DimensionService interface {
	GetDimensions(ctx context.Context) (game.Dimensions, error)
	GetDeletedDimensions(ctx context.Context) (game.Dimensions, error)
	GetDimensionById(ctx context.Context, dimensionId *uuid.UUID) (*game.Dimension, error)
	CreateDimension(ctx context.Context, name, version, location string, mapIds []uuid.UUID) (*game.Dimension, error)
	DeleteDimension(ctx context.Context, dimensionId *uuid.UUID) (*game.Dimension, error)
	EditDimension(ctx context.Context, dimension *game.Dimension) (*game.Dimension, error)
}

type dimensionService struct {
	repo repository.DimensionRepository
}

func NewDimensionService(repo repository.DimensionRepository) DimensionService {
	return &dimensionService{repo: repo}
}

// CreateDimension implements DimesionService.
func (d *dimensionService) CreateDimension(ctx context.Context, name, version, location string, mapIds []uuid.UUID) (*game.Dimension, error) {
	maps := make([]*game.Map, len(mapIds))
	for idx, mapId := range mapIds {
		maps[idx] = &game.Map{
			Id: mapId,
		}
	}

	dimension := &game.Dimension{
		Name:     name,
		Version:  version,
		Maps:     maps,
		Location: location,
	}

	err := dimension.Validate()
	if err != nil {
		return nil, err
	}

	return d.repo.CreateDimension(ctx, dimension)
}

// DeleteDimension implements DimesionService.
func (d *dimensionService) DeleteDimension(ctx context.Context, dimensionId *uuid.UUID) (*game.Dimension, error) {
	return d.repo.DeleteDimension(ctx, dimensionId)
}

// EditDimension implements DimesionService.
func (d *dimensionService) EditDimension(ctx context.Context, dimension *game.Dimension) (*game.Dimension, error) {
	err := dimension.Validate()
	if err != nil {
		return nil, err
	}

	return d.repo.UpdateDimension(ctx, dimension)
}

// GetDimensionById implements DimesionService.
func (d *dimensionService) GetDimensionById(ctx context.Context, dimensionId *uuid.UUID) (*game.Dimension, error) {
	return d.repo.GetDimensionById(ctx, dimensionId)
}

// GetDimensions implements DimesionService.
func (d *dimensionService) GetDimensions(ctx context.Context) (game.Dimensions, error) {
	return d.repo.GetDimensions(ctx)
}

// GetDimensions implements DimesionService.
func (d *dimensionService) GetDeletedDimensions(ctx context.Context) (game.Dimensions, error) {
	return d.repo.GetDeletedDimensions(ctx)
}
