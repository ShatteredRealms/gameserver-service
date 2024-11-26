package repository

import (
	"context"

	"github.com/ShatteredRealms/gameserver-service/pkg/model/game"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type postgresMapRepository struct {
	gormdb *gorm.DB
}

// CreateMap implements MapRepository.
func (p *postgresMapRepository) CreateMap(ctx context.Context, m *game.Map) (*game.Map, error) {
	if m.Id != nil {
		return nil, ErrDimensionIdProvided
	}

	err := p.db(ctx).Create(m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

// DeleteMap implements MapRepository.
func (p *postgresMapRepository) DeleteMap(ctx context.Context, mapId *uuid.UUID) (m *game.Map, _ error) {
	result := p.db(ctx).Clauses(clause.Returning{}).Delete(&m, "id = ?", mapId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}

	return m, nil
}

// GetMapById implements MapRepository.
func (p *postgresMapRepository) GetMapById(ctx context.Context, mapId *uuid.UUID) (m *game.Map, _ error) {
	result := p.db(ctx).Find(&m, "id = ?", mapId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}

	return m, nil
}

// GetMaps implements MapRepository.
func (p *postgresMapRepository) GetMaps(ctx context.Context) (m *game.Maps, _ error) {
	return m, p.db(ctx).Find(&m).Error
}

// UpdateMap implements MapRepository.
func (p *postgresMapRepository) UpdateMap(ctx context.Context, m *game.Map) (*game.Map, error) {
	return m, p.db(ctx).Save(m).Error
}

func NewPostgresMapRepository(db *gorm.DB) MapRepository {
	db.AutoMigrate(&game.Map{})
	return &postgresMapRepository{gormdb: db}
}

func (p *postgresMapRepository) db(ctx context.Context) *gorm.DB {
	return p.gormdb.WithContext(ctx)
}
