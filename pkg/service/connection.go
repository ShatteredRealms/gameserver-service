package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ShatteredRealms/gameserver-service/pkg/model/gameserver"
	"github.com/ShatteredRealms/gameserver-service/pkg/repository"
	"github.com/ShatteredRealms/go-common-service/pkg/common"
	"github.com/ShatteredRealms/go-common-service/pkg/log"
	"github.com/google/uuid"
)

var (
	ErrInvalidServerName = fmt.Errorf("invalid server name")
	ErrExpiredConnection = errors.New("connection request expired")
)

type ConnectionService interface {
	CreatePendingConnection(
		ctx context.Context,
		characterId string,
		serverName string,
	) (*gameserver.PendingConnection, error)
	CheckPlayerConnection(
		ctx context.Context,
		id *uuid.UUID,
		serverName string,
	) (*gameserver.PendingConnection, error)
}

func NewConnectionService(repo repository.ConnectionRepository) ConnectionService {
	return &connectionService{repo: repo}
}

type connectionService struct {
	repo repository.ConnectionRepository
}

// CheckPlayerConnection implements ConnectionService.
func (c *connectionService) CheckPlayerConnection(
	ctx context.Context,
	id *uuid.UUID,
	serverName string,
) (*gameserver.PendingConnection, error) {
	pc, err := c.repo.FindPendingConnection(ctx, id)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := c.repo.DeletePendingConnection(ctx, id)
		if err != nil {
			log.Logger.WithContext(ctx).Errorf("failed to delete connection request: %s", err)
		}
	}()

	if pc.ServerName != serverName {
		log.Logger.WithContext(ctx).Errorf("connection request server name mismatch: %s != %s, genereated by character %s", pc.ServerName, serverName, pc.CharacterId)
		return nil, ErrInvalidServerName
	}

	if pc.CreatedAt.Add(time.Minute).Before(time.Now()) {
		return nil, ErrExpiredConnection
	}

	return pc, nil
}

// CreatePendingConnection implements ConnectionService.
func (c *connectionService) CreatePendingConnection(
	ctx context.Context,
	characterId string,
	serverName string,
) (*gameserver.PendingConnection, error) {
	id, err := uuid.Parse(characterId)
	if err != nil {
		return nil, common.ErrInvalidId
	}
	return c.repo.CreatePendingConnection(ctx, &id, serverName)
}
