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

	pg, err := commonrepo.ConnectDB(ctx, commoncfg.DBPoolConfig{Master: cfg.Postgres}, cfg.Redis)
	if err != nil {
		return nil, fmt.Errorf("connect db: %w", err)
	}

	migrater, err := commonrepo.NewPgxMigrater(ctx, cfg.Postgres.PostgresDSN(), "migrations")
	if err != nil {
		return nil, fmt.Errorf("new migrater: %w", err)
	}

	gsCtx.DimensionService = service.NewDimensionService(
		repository.NewPgxDimensionRepository(migrater),
	)
	gsCtx.MapService = service.NewMapService(
		repository.NewPgxMapRepository(migrater),
	)
	gsCtx.ConnectionService = service.NewConnectionService(
		repository.NewPgxConnectionRepository(migrater),
	)
	gsCtx.CharacterService = characterbus.NewService(
		characterbus.NewPostgresRepository(pg),
		bus.NewKafkaMessageBusReader(cfg.Kafka, serviceName, characterbus.Message{}),
	)
	gsCtx.CharacterService.StartProcessing(ctx)

	if gsCtx.UsingAgones() {
		gsCtx.GsmService, err = service.NewGameServerManagerService(
			&cfg.Gsm,
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

func (c *GameServerContext) ResetDimensionBus() commonsrv.WriterResetCallback {
	return func(ctx context.Context) error {
		ctx, span := c.Context.Tracer.Start(ctx, "dimension.reset_dimension_bus")
		defer span.End()
		dimensions, err := c.DimensionService.GetDimensions(ctx)
		if err != nil {
			return fmt.Errorf("get dimensions: %w", err)
		}
		deletedDimensions, err := c.DimensionService.GetDeletedDimensions(ctx)
		if err != nil {
			return fmt.Errorf("get deleted dimensions: %w", err)
		}

		msgs := make([]dimensionbus.Message, len(dimensions)+len(deletedDimensions))
		for idx, char := range dimensions {
			msgs[idx] = dimensionbus.Message{
				Id:      char.Id,
				Deleted: false,
			}
		}
		for idx, char := range deletedDimensions {
			msgs[idx+len(dimensions)] = dimensionbus.Message{
				Id:      char.Id,
				Deleted: true,
			}
		}

		return c.DimensionBusWriter.PublishMany(ctx, msgs)
	}
}

func (c *GameServerContext) ResetMapBus() commonsrv.WriterResetCallback {
	return func(ctx context.Context) error {
		ctx, span := c.Context.Tracer.Start(ctx, "map.reset_map_bus")
		defer span.End()
		maps, err := c.MapService.GetMaps(ctx)
		if err != nil {
			return fmt.Errorf("get maps: %w", err)
		}
		deletedMaps, err := c.MapService.GetDeletedMaps(ctx)
		if err != nil {
			return fmt.Errorf("get deleted maps: %w", err)
		}

		msgs := make([]mapbus.Message, len(maps)+len(deletedMaps))
		for idx, char := range maps {
			msgs[idx] = mapbus.Message{
				Id:      char.Id,
				Deleted: false,
			}
		}
		for idx, char := range deletedMaps {
			msgs[idx+len(maps)] = mapbus.Message{
				Id:      char.Id,
				Deleted: true,
			}
		}

		return c.MapBusWriter.PublishMany(ctx, msgs)
	}
}
