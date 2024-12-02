package service

import (
	"context"
	"errors"
	"fmt"
	"sync"

	v1 "agones.dev/agones/pkg/apis/agones/v1"
	aav1 "agones.dev/agones/pkg/apis/allocation/v1"
	"agones.dev/agones/pkg/client/clientset/versioned"
	"github.com/ShatteredRealms/gameserver-service/pkg/config"
	"github.com/ShatteredRealms/gameserver-service/pkg/model/game"
	"github.com/ShatteredRealms/go-common-service/pkg/log"
	"k8s.io/client-go/rest"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	ErrNoServersAvailable = errors.New("no servers available")
)

type dimensionMap struct {
	Dimension *game.Dimension
	Map       *game.Map
	Created   bool
}

type GameServerManagerService interface {
	RequestConnection(ctx context.Context, characterId, dimensionId, mapId string) (*aav1.GameServerAllocation, error)
	DimensionMapChanged(dimension *game.Dimension, m *game.Map, created bool)
	Start(ctx context.Context)
	Stop()
}

type gsmService struct {
	config *config.GameServerManagerConfig

	agones versioned.Interface

	DimensionMapsChanged chan dimensionMap
	mu                   sync.Mutex
	ctx                  context.Context
	cancelFunc           context.CancelFunc
}

// DimensionMapChanged implements GameServerManagerService.
func (g *gsmService) DimensionMapChanged(dimension *game.Dimension, m *game.Map, created bool) {
	g.DimensionMapsChanged <- dimensionMap{
		Dimension: dimension,
		Map:       m,
		Created:   created,
	}
}

// RequestConnection implements GameServerManagerService.
func (g *gsmService) RequestConnection(ctx context.Context, characterId, dimensionId, mapId string) (*aav1.GameServerAllocation, error) {
	log.Logger.WithContext(ctx).Debugf(
		"character '%s' in dimension '%s' requesting connection to gameserver with map '%s'",
		characterId, dimensionId, mapId,
	)

	allocatedState := v1.GameServerStateAllocated
	readyState := v1.GameServerStateReady
	gsAlloc, err := g.agones.AllocationV1().GameServerAllocations(g.config.GameServerNamespace).Create(
		ctx,
		&aav1.GameServerAllocation{
			TypeMeta:   metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{},
			Spec: aav1.GameServerAllocationSpec{
				Selectors: []aav1.GameServerSelector{
					{
						GameServerState: &allocatedState,
						Players: &aav1.PlayerSelector{
							MinAvailable: 1,
							MaxAvailable: 1000,
						},
						LabelSelector: metav1.LabelSelector{
							MatchLabels: map[string]string{
								config.MapLabel:       mapId,
								config.DimensionLabel: dimensionId,
							},
						},
					},
					{
						GameServerState: &readyState,
						Players: &aav1.PlayerSelector{
							MinAvailable: 1,
							MaxAvailable: 1000,
						},
						LabelSelector: metav1.LabelSelector{
							MatchLabels: map[string]string{
								config.MapLabel:       mapId,
								config.DimensionLabel: dimensionId,
							},
						},
					},
				},
			},
		},
		metav1.CreateOptions{},
	)
	if err != nil {
		return nil, fmt.Errorf("create gameserver allocation: %w", err)
	}

	if len(gsAlloc.Status.Ports) <= 0 {
		log.Logger.WithContext(ctx).Errorf("no servers available for gameserver allocation for dimension '%s' map '%s'", dimensionId, mapId)
		return nil, ErrNoServersAvailable
	}

	return gsAlloc, nil
}

// Start implements GameServerManagerService.
func (g *gsmService) Start(ctx context.Context) {
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
func (g *gsmService) Stop() {
	g.cancelFunc()
}

func (g *gsmService) createGameServers(ctx context.Context, dimension *game.Dimension, m *game.Map) error {
	fleet, err := g.agones.AgonesV1().
		Fleets(g.config.GameServerNamespace).
		Create(ctx, g.config.GetFleetTemplate(dimension, m), metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("create fleet: %w", err)
	}

	_, err = g.agones.AutoscalingV1().
		FleetAutoscalers(fleet.Namespace).
		Create(ctx, g.config.GetFleetAutoscalerTemplate(dimension, m), metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("create fleet autoscaler: %w", err)
	}

	return nil
}

func (g *gsmService) deleteGameServers(ctx context.Context, dimension *game.Dimension, m *game.Map) error {
	err := g.agones.AutoscalingV1().
		FleetAutoscalers(g.config.GameServerNamespace).
		Delete(ctx, g.config.GetFleetAutoscalerTemplate(dimension, m).GetName(), metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("delete fleet autoscaler: %w", err)
	}

	err = g.agones.AgonesV1().
		Fleets(g.config.GameServerNamespace).
		Delete(ctx, g.config.GetFleetTemplate(dimension, m).GetName(), metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("delete fleet: %w", err)
	}

	return nil
}

func NewGameServerManagerService(
	gameServerConfig config.GameServerManagerConfig,
) (GameServerManagerService, error) {
	g := &gsmService{}

	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("get kubernetes config: %w", err)
	}

	g.agones, err = versioned.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("create agones client: %w", err)
	}

	g.DimensionMapsChanged = make(chan dimensionMap, 30)

	return g, nil
}
