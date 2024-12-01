package repository

import (
	"context"

	"github.com/ShatteredRealms/gameserver-service/pkg/model/gameserver"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type postgresConnectionRepository struct {
	gormdb *gorm.DB
}

func NewPostgresConnectionRepository(gormdb *gorm.DB) ConnectionRepository {
	gormdb.AutoMigrate(&gameserver.PendingConnection{})
	return &postgresConnectionRepository{gormdb: gormdb}
}

// CreatePendingConnection implements ConnectionRepository.
func (p *postgresConnectionRepository) CreatePendingConnection(
	ctx context.Context,
	character string,
	serverName string,
) (*gameserver.PendingConnection, error) {
	pc := &gameserver.PendingConnection{
		CharacterId: character,
		ServerName:  serverName,
	}
	return pc, p.db(ctx).Create(&pc).Error
}

// DeletePendingConnection implements ConnectionRepository.
func (p *postgresConnectionRepository) DeletePendingConnection(
	ctx context.Context,
	id *uuid.UUID,
) error {
	return p.db(ctx).Delete(&gameserver.PendingConnection{}, id).Error
}

// FindPendingConnection implements ConnectionRepository.
func (p *postgresConnectionRepository) FindPendingConnection(
	ctx context.Context,
	id *uuid.UUID,
) (pc *gameserver.PendingConnection, _ error) {
	return pc, p.db(ctx).First(&pc, id).Error
}

func (p *postgresConnectionRepository) db(ctx context.Context) *gorm.DB {
	return p.gormdb.WithContext(ctx)
}
