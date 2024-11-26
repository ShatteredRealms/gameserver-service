package repository

import (
	"context"

	"github.com/ShatteredRealms/dimension-service/pkg/model/game"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type postgresMapRepository struct {
	gormdb *gorm.DB
}

func NewPostgresMapRepository(db *gorm.DB) MapRepository {
	db.AutoMigrate(&game.Map{})
	return &postgresMapRepository{gormdb: db}
}

// CreateMap implements MapRepository.
func (p *postgresMapRepository) CreateMap(ctx context.Context, mapId string) (m *game.Map, _ error) {
	m.Id = mapId
	return m, p.db(ctx).Create(m).Error
}

// DeleteMap implements MapRepository.
func (p *postgresMapRepository) DeleteMap(ctx context.Context, mapId string) (m *game.Map, _ error) {
	return m, p.db(ctx).Clauses(clause.Returning{}).Delete(m, "id = ?", mapId).Error
}

// GetMapById implements MapRepository.
func (p *postgresMapRepository) GetMapById(ctx context.Context, mapId string) (m *game.Map, _ error) {
	result := p.db(ctx).Where("id = ?", mapId).Find(&m)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return m, nil
}

// GetMaps implements MapRepository.
func (p *postgresMapRepository) GetMaps(ctx context.Context) (maps *game.Maps, _ error) {
	return maps, p.db(ctx).Find(maps).Error
}

func (p *postgresMapRepository) db(ctx context.Context) *gorm.DB {
	return p.gormdb.WithContext(ctx)
}
