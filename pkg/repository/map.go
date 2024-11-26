package repository

import (
	"context"

	"github.com/ShatteredRealms/dimension-service/pkg/model/game"
)

type MapRepository interface {
	GetMapById(ctx context.Context, mapId string) (*game.Map, error)

	GetMaps(ctx context.Context) (*game.Maps, error)

	CreateMap(ctx context.Context, mapId string) (*game.Map, error)

	DeleteMap(ctx context.Context, mapId string) (*game.Map, error)
}
