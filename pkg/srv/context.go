package srv

import (
	"context"
	"fmt"

	"github.com/ShatteredRealms/gameserver-service/pkg/config"
	"github.com/ShatteredRealms/gameserver-service/pkg/repository"
	"github.com/ShatteredRealms/gameserver-service/pkg/service"
	"github.com/ShatteredRealms/go-common-service/pkg/bus"
	"github.com/ShatteredRealms/go-common-service/pkg/bus/character/characterbus"
	"github.com/ShatteredRealms/go-common-service/pkg/bus/gameserver/dimensionbus"
	"github.com/ShatteredRealms/go-common-service/pkg/bus/gameserver/mapbus"
	commoncfg "github.com/ShatteredRealms/go-common-service/pkg/config"
	commonrepo "github.com/ShatteredRealms/go-common-service/pkg/repository"
	commonsrv "github.com/ShatteredRealms/go-common-service/pkg/srv"
)

type GameServerContext struct {
	*commonsrv.Context

	DimensionBusWriter bus.MessageBusWriter[dimensionbus.Message]
	MapBusWriter       bus.MessageBusWriter[mapbus.Message]

	DimensionService  service.DimensionService
	MapService        service.MapService
	ConnectionService service.ConnectionService
	GsmService        service.GameServerManagerService

	CharacterService characterbus.Service
}

func NewGameServerContext(ctx context.Context, cfg *config.GameServerConfig, serviceName string) (*GameServerContext, error) {
	gsCtx := &GameServerContext{
		Context:            commonsrv.NewContext(&cfg.BaseConfig, serviceName),
		DimensionBusWriter: bus.NewKafkaMessageBusWriter(cfg.Kafka, dimensionbus.Message{}),
		MapBusWriter:       bus.NewKafkaMessageBusWriter(cfg.Kafka, mapbus.Message{}),
	}
	ctx, span := gsCtx.Tracer.Start(ctx, "context.dimension.new")
	defer span.End()

	pg, err := commonrepo.ConnectDB(ctx, cfg.Postgres, cfg.Redis)
	if err != nil {
		return nil, fmt.Errorf("connect db: %w", err)
	}

	gsCtx.DimensionService = service.NewDimensionService(
		repository.NewPostgresDimensionRepository(pg),
	)
	gsCtx.MapService = service.NewMapService(
		repository.NewPostgresMapRepository(pg),
	)
	gsCtx.ConnectionService = service.NewConnectionService(
		repository.NewPostgresConnectionRepository(pg),
	)
	gsCtx.CharacterService = characterbus.NewService(
		characterbus.NewPostgresRepository(pg),
		bus.NewKafkaMessageBusReader(cfg.Kafka, serviceName, characterbus.Message{}),
	)
	gsCtx.CharacterService.StartProcessing(ctx)

	if gsCtx.UsingAgones() {
		gsCtx.GsmService, err = service.NewGameServerManagerService(
			cfg.GsmConfig,
		)
		if err != nil {
			return nil, fmt.Errorf("new gsm service: %w", err)
		}
		gsCtx.GsmService.Start(ctx)
	}

	return gsCtx, nil
}

func (gsCtx *GameServerContext) UsingAgones() bool {
	return gsCtx.Config.Mode != commoncfg.ModeLocal
}
