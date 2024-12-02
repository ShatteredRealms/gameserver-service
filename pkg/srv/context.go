package srv

import (
	"context"
	"fmt"

	"github.com/ShatteredRealms/gameserver-service/pkg/config"
	"github.com/ShatteredRealms/gameserver-service/pkg/repository"
	"github.com/ShatteredRealms/gameserver-service/pkg/service"
	"github.com/ShatteredRealms/go-common-service/pkg/bus"
	commonrepo "github.com/ShatteredRealms/go-common-service/pkg/repository"
	commonsrv "github.com/ShatteredRealms/go-common-service/pkg/srv"
)

type GameServerContext struct {
	*commonsrv.Context

	DimensionBusWriter bus.MessageBusWriter[bus.DimensionMessage]
	MapBusWriter       bus.MessageBusWriter[bus.MapMessage]

	DimensionService service.DimensionService
	MapService       service.MapService
}

func NewDimensionContext(ctx context.Context, cfg *config.GameServerConfig, serviceName string) (*GameServerContext, error) {
	dimensionCtx := &GameServerContext{
		Context:            commonsrv.NewContext(&cfg.BaseConfig, serviceName),
		DimensionBusWriter: bus.NewKafkaMessageBusWriter(cfg.Kafka, bus.DimensionMessage{}),
		MapBusWriter:       bus.NewKafkaMessageBusWriter(cfg.Kafka, bus.MapMessage{}),
	}
	ctx, span := dimensionCtx.Tracer.Start(ctx, "context.dimension.new")
	defer span.End()

	pg, err := commonrepo.ConnectDB(ctx, cfg.Postgres, cfg.Redis)
	if err != nil {
		return nil, fmt.Errorf("connect db: %w", err)
	}

	dimensionCtx.DimensionService = service.NewDimensionService(
		repository.NewPostgresDimensionRepository(pg),
	)
	dimensionCtx.MapService = service.NewMapService(
		repository.NewPostgresMapRepository(pg),
	)

	return dimensionCtx, nil
}
