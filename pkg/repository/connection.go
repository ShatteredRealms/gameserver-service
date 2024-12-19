package repository

import (
	"context"

	"github.com/ShatteredRealms/gameserver-service/pkg/model/gameserver"
	"github.com/google/uuid"
)

type ConnectionRepository interface {
	CreatePendingConnection(
		ctx context.Context,
		characterId *uuid.UUID,
		serverName string,
	) (*gameserver.PendingConnection, error)
	DeletePendingConnection(ctx context.Context, id *uuid.UUID) error
	FindPendingConnection(
		ctx context.Context,
		id *uuid.UUID,
	) (*gameserver.PendingConnection, error)
}
