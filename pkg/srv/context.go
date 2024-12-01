package srv

import (
	"context"
	"fmt"

	"agones.dev/agones/pkg/client/clientset/versioned"
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
	"k8s.io/client-go/rest"
)

type GameServerContext struct {
	*commonsrv.Context

	DimensionBusWriter bus.MessageBusWriter[dimensionbus.Message]
	MapBusWriter       bus.MessageBusWriter[mapbus.Message]

	DimensionService  service.DimensionService
	MapService        service.MapService
	ConnectionService service.ConnectionService

	CharacterService characterbus.Service

	Agones *AgonesClient
}

func NewDimensionContext(ctx context.Context, cfg *config.DimensionConfig, serviceName string) (*GameServerContext, error) {
	dimensionCtx := &GameServerContext{
		Context:            commonsrv.NewContext(&cfg.BaseConfig, serviceName),
		DimensionBusWriter: bus.NewKafkaMessageBusWriter(cfg.Kafka, dimensionbus.Message{}),
		MapBusWriter:       bus.NewKafkaMessageBusWriter(cfg.Kafka, mapbus.Message{}),
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
	dimensionCtx.ConnectionService = service.NewConnectionService(
		repository.NewPostgresConnectionRepository(pg),
	)
	dimensionCtx.CharacterService = characterbus.NewService(
		characterbus.NewPostgresRepository(pg),
		bus.NewKafkaMessageBusReader(cfg.Kafka, serviceName, characterbus.Message{}),
	)

	if cfg.Mode != commoncfg.ModeLocal {
		conf, err := rest.InClusterConfig()
		if err != nil {
			return nil, fmt.Errorf("in cluster config: %w", err)
		}

		dimensionCtx.Agones, err := versioned.NewForConfig(conf)
		if err != nil {
			return nil, fmt.Errorf("agones config: %w", err)
		}
	}

	return dimensionCtx, nil
}
