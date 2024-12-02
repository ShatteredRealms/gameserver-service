package service

import (
	"context"
	"fmt"
	"sync"

	"agones.dev/agones/pkg/client/clientset/versioned"
	"github.com/ShatteredRealms/gameserver-service/pkg/config"
	"github.com/ShatteredRealms/gameserver-service/pkg/model/game"
	"github.com/ShatteredRealms/go-common-service/pkg/log"
	"k8s.io/client-go/rest"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DimensionMap struct {
	Dimension *game.Dimension
	Map       *game.Map
	Created   bool
}

type GsmService struct {
	config           config.GameServerManagerConfig
	mapService       MapService
	dimensionService DimensionService

	DimensionMapsChanged chan DimensionMap

	agonesClient versioned.Interface

	mu         sync.Mutex
	ctx        context.Context
	cancelFunc context.CancelFunc
}

// Start implements GameServerManagerService.
func (g *GsmService) Start(ctx context.Context) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.ctx != nil {
		return
	}

	go func() {
		g.mu.Lock()
		g.ctx, g.cancelFunc = context.WithCancel(ctx)
		g.mu.Unlock()
		for {
			select {
			case data := <-g.DimensionMapsChanged:
				if data.Created {
					go func() {
						err := g.createGameServers(g.ctx, data.Dimension, data.Map)
						if err != nil {
							log.Logger.WithContext(ctx).Errorf("error creating dimension '%s' map '%s': %v", data.Dimension.Id, data.Map.Id, err)
						}
					}()
				} else {
					go func() {
						err := g.createGameServers(g.ctx, data.Dimension, data.Map)
						if err != nil {
							log.Logger.WithContext(ctx).Errorf("error deleting dimension '%s' map '%s': %v", data.Dimension.Id, data.Map.Id, err)
						}
					}()
				}
			case <-g.ctx.Done():
				log.Logger.WithContext(g.ctx).Info("game server manager service stopped")
				g.mu.Lock()
				g.ctx = nil
				g.cancelFunc = nil
				g.mu.Unlock()
				return
			}
		}
	}()
}

// Stop implements GameServerManagerService.
func (g *GsmService) Stop() {
	g.cancelFunc()
}

func (g *GsmService) createGameServers(ctx context.Context, dimension *game.Dimension, m *game.Map) error {
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

func (g *GsmService) deleteGameServers(ctx context.Context, dimension *game.Dimension, m *game.Map) error {
	err := g.agonesClient.AutoscalingV1().
		FleetAutoscalers(g.config.GameServerNamespace).
		Delete(ctx, g.config.GetFleetAutoscalerTemplate(dimension, m).GetName(), metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("delete fleet autoscaler: %w", err)
	}

	err = g.agonesClient.AgonesV1().
		Fleets(g.config.GameServerNamespace).
		Delete(ctx, g.config.GetFleetTemplate(dimension, m).GetName(), metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("delete fleet: %w", err)
	}

	return nil
}

func NewGameServerManagerService(
	gameServerConfig config.GameServerManagerConfig,
	mapService MapService,
	dimensionService DimensionService,
) (*GsmService, error) {
	g := &GsmService{
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

	g.DimensionMapsChanged = make(chan DimensionMap, 30)

	return g, nil
}
