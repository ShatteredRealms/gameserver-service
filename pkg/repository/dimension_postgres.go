package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/ShatteredRealms/dimension-service/pkg/model/game"
	"github.com/ShatteredRealms/go-common-service/pkg/srospan"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrDimension           = errors.New("dimension repository")
	ErrDimensionIdProvided = fmt.Errorf("%w: id provided", ErrDimension)
)

type postgresDimensionRepository struct {
	gormdb *gorm.DB
}

func NewPostgresDimensionRepository(gormdb *gorm.DB) DimensionRepository {
	gormdb.AutoMigrate(&game.Dimension{})
	return &postgresDimensionRepository{gormdb: gormdb}
}

// CreateDimension implements DimensionRepository.
func (p *postgresDimensionRepository) CreateDimension(ctx context.Context, dimension *game.Dimension) (*game.Dimension, error) {
	if dimension.Id != nil {
		return nil, ErrDimensionIdProvided
	}

	err := p.db(ctx).Create(dimension).Error
	if err != nil {
		return nil, err
	}

	updateSpanWithDimension(ctx, dimension.Id.String())
	return dimension, nil
}

// DeleteDimension implements DimensionRepository.
func (p *postgresDimensionRepository) DeleteDimension(ctx context.Context, dimensionId *uuid.UUID) (dimension *game.Dimension, err error) {
	result := p.db(ctx).Clauses(clause.Returning{}).Delete(&dimension, "id = ?", dimensionId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}

	updateSpanWithDimension(ctx, dimension.Id.String())
	return dimension, nil
}

// GetDimensionById implements DimensionRepository.
func (p *postgresDimensionRepository) GetDimensionById(ctx context.Context, dimensionId *uuid.UUID) (dimension *game.Dimension, _ error) {
	result := p.db(ctx).Find(&dimension, "id = ?", dimensionId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}

	updateSpanWithDimension(ctx, dimension.Id.String())
	return dimension, nil
}

// GetDimensions implements DimensionRepository.
func (p *postgresDimensionRepository) GetDimensions(ctx context.Context) (dimensions *game.Dimensions, _ error) {
	return dimensions, p.db(ctx).Find(&dimensions).Error
}

// UpdateDimension implements DimensionRepository.
func (p *postgresDimensionRepository) UpdateDimension(ctx context.Context, dimension *game.Dimension) (*game.Dimension, error) {
	return dimension, p.db(ctx).Save(dimension).Error
}

func (p *postgresDimensionRepository) db(ctx context.Context) *gorm.DB {
	return p.gormdb.WithContext(ctx)
}

func updateSpanWithDimension(ctx context.Context, dimensionId string) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		srospan.DimensionId(dimensionId),
	)
}

