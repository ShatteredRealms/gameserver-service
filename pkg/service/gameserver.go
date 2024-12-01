package service

import (
	"context"
	"errors"
	"fmt"

	"agones.dev/agones/pkg/client/clientset/versioned"
	"github.com/ShatteredRealms/gameserver-service/pkg/config"
	"github.com/ShatteredRealms/gameserver-service/pkg/model/game"
	"k8s.io/client-go/rest"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GameServerManagerService interface {
}

type gsmService struct {
	config           config.GameServerManagerConfig
	mapService       MapService
	dimensionService DimensionService

	agonesClient versioned.Interface
}

func (g *gsmService) setupNewDimension(ctx context.Context, dimension *game.Dimension) error {
	var errs error
	for _, m := range dimension.Maps {
		err := g.createGameServers(ctx, dimension, m)
		if err != nil {
			errs = errors.Join(errs, err)
		}
	}

	return errs
}

func (g *gsmService) createGameServers(ctx context.Context, dimension *game.Dimension, m *game.Map) error {
	fleet, err := g.agonesClient.AgonesV1().
		Fleets(g.config.GameServerNamespace).
		Create(ctx, g.config.GetFleetTemplate(dimension, m), metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("create fleet: %w", err)
	}

	_, err = g.agonesClient.AutoscalingV1().
		FleetAutoscalers(fleet.Namespace).
		Create(ctx, g.config.GetFleetAutoscalerTemplate(dimension, m), metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("create fleet autoscaler: %w", err)
	}

	return nil
}

func NewGameServerManagerService(
	gameServerConfig config.GameServerManagerConfig,
	mapService MapService,
	dimensionService DimensionService,
) (GameServerManagerService, error) {
	g := &gsmService{
		mapService:       mapService,
		dimensionService: dimensionService,
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("get kubernetes config: %w", err)
	}

	g.agonesClient, err = versioned.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("create agones client: %w", err)
	}

	return g, nil
}
