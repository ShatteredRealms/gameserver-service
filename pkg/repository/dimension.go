package repository

import (
	"context"

	"github.com/ShatteredRealms/dimension-service/pkg/model/game"
	"github.com/google/uuid"
)

type DimensionRepository interface {
	GetDimensions(ctx context.Context) (*game.Dimensions, error)

	GetDimensionById(ctx context.Context, dimensionId *uuid.UUID) (*game.Dimension, error)

	CreateDimension(ctx context.Context, dimension *game.Dimension) (*game.Dimension, error)

	UpdateDimension(ctx context.Context, dimension *game.Dimension) (*game.Dimension, error)

	DeleteDimension(ctx context.Context, dimensionId *uuid.UUID) (*game.Dimension, error)
}
