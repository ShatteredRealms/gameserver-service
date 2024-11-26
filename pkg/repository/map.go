package repository

import (
	"context"

	"github.com/ShatteredRealms/gameserver-service/pkg/model/game"
	"github.com/google/uuid"
)

type MapRepository interface {
	GetMaps(ctx context.Context) (*game.Maps, error)

	GetMapById(ctx context.Context, mapId *uuid.UUID) (*game.Map, error)

	CreateMap(ctx context.Context, m *game.Map) (*game.Map, error)

	UpdateMap(ctx context.Context, m *game.Map) (*game.Map, error)

	DeleteMap(ctx context.Context, mapId *uuid.UUID) (*game.Map, error)
}
