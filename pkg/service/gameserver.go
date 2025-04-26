package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"

	v1 "agones.dev/agones/pkg/apis/agones/v1"
	aav1 "agones.dev/agones/pkg/apis/allocation/v1"
	"agones.dev/agones/pkg/client/clientset/versioned"
	"github.com/ShatteredRealms/gameserver-service/pkg/config"
	"github.com/ShatteredRealms/gameserver-service/pkg/model/game"
	"github.com/ShatteredRealms/go-common-service/pkg/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	ErrNoServersAvailable = errors.New("no servers available")
)

const (
	PLAYER_LIST_NAME     = "players"
	CHARACTERS_LIST_NAME = "characters"
)

type dimensionMap struct {
	Dimension *game.Dimension
	Map       *game.Map
	Created   bool
}

type GameServerManagerService interface {
	// FindAvailableGameServers will find the best game server(s) for the given dimension and map.
	FindAvailableGameServers(ctx context.Context, dimensionId, mapId string) (*aav1.GameServerAllocation, error)

	// DimensionMapChanged creates or deletes a game server fleets and autoscalers for the given dimension and game based on created.
	DimensionMapChanged(ctx context.Context, dimension *game.Dimension, m *game.Map, created bool)

	// SyncGameServers will create or delete game servers based on the maps in the dimension.
	// If the returned arrays is nil, then no maps were created or deleted and error will be returned.
	// Returns the maps that attempted to be created and deleted and any errors that occurred.
	SyncGameServers(ctx context.Context, dimension *game.Dimension) (mapsCreated []string, mapsDeleted []string, err error)

	// Start starts processing incoming dimension map changes on seperate threads.
	Start(ctx context.Context)

	// Stop stops processing incoming dimension map changes.
	Stop()

	AnyCharactersConneted(ctx context.Context, characterIds []string) (bool, error)
	AnyUsersConneted(ctx context.Context, userIds []string) (bool, error)

	CountGameServers(ctx context.Context) (int, error)
}

type gsmService struct {
	config *config.GameServerManagerConfig

	agones    versioned.Interface
	clientset *kubernetes.Clientset

	DimensionMapsChanged chan dimensionMap
	mu                   sync.Mutex
	ctx                  context.Context
	cancelFunc           context.CancelFunc
}

func NewGameServerManagerService(
	gameServerConfig *config.GameServerManagerConfig,
) (GameServerManagerService, error) {
	// If the kube config path is empty, then the in-cluster config will be used.
	path := os.ExpandEnv(gameServerConfig.KubeConfigPath)
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		return nil, fmt.Errorf("get kubernetes config: %w", err)
	}
	// Create the kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("create kubernetes clientset: %w", err)
	}

	g := &gsmService{
		config:               gameServerConfig,
		DimensionMapsChanged: make(chan dimensionMap, 30),
		clientset:            clientset,
	}

	g.agones, err = versioned.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("create agones client: %w", err)
	}

	return g, nil
}

func (g *gsmService) anyConnected(ctx context.Context, ids []string, key string) (bool, error) {

	selectors := make([]aav1.GameServerSelector, len(ids))
	for idx, id := range ids {
		selectors[idx] = aav1.GameServerSelector{
			Lists: map[string]aav1.ListSelector{
				key: {
					ContainsValue: id,
				},
			},
		}
	}

	servers, err := g.agones.AgonesV1().GameServers(g.config.GameServerNamespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return false, fmt.Errorf("list gameservers: %w", err)
	}

	mapIds := make(map[string]struct{})
	for _, id := range ids {
		mapIds[id] = struct{}{}
	}
	for _, server := range servers.Items {
		list, ok := server.Status.Lists[key]
		if !ok {
			return false, fmt.Errorf("gameserver %s does not have %s list", server.Name, key)
		}
		for _, player := range list.Values {
			if _, ok := mapIds[player]; ok {
				return true, nil
			}
		}
	}

	return false, nil
}

func (g *gsmService) AnyUsersConneted(ctx context.Context, characterIds []string) (bool, error) {
	return g.anyConnected(ctx, characterIds, PLAYER_LIST_NAME)
}

func (g *gsmService) AnyCharactersConneted(ctx context.Context, characterIds []string) (bool, error) {
	return g.anyConnected(ctx, characterIds, CHARACTERS_LIST_NAME)
}

// DeleteExtra implements GameServerManagerService.
func (g *gsmService) SyncGameServers(
	ctx context.Context,
	dimension *game.Dimension,
) (mapsCreated []string, mapsDeleted []string, err error) {
	pendingCreation := make([]*game.Map, 0)
	mapsCreated = make([]string, 0)
	mapsDeleted = make([]string, 0)

	// Find all the fleets that are associated with the dimension
	labels, err := metav1.LabelSelectorAsSelector(
		&metav1.LabelSelector{
			MatchLabels: g.config.GetLabels(dimension.Id.String(), ""),
		},
	)
	if err != nil {
		return nil, nil, fmt.Errorf("create label selector: %w", err)
	}

	fleetList, err := g.agones.AgonesV1().Fleets(g.config.GameServerNamespace).List(ctx, metav1.ListOptions{
		LabelSelector: labels.String(),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("list fleets: %w", err)
	}

	// Find maps that are required, but not created and add them to a pending creation list.
	// Also, remove items from the fleetList that already exist so the remaining items can be deleted.
	for _, m := range dimension.Maps {
		found := false
		for idx, fleet := range fleetList.Items {
			if fleet.Labels[config.MapLabel] == m.Id.String() {
				found = true
				fleetList.Items[idx] = fleetList.Items[len(fleetList.Items)-1]
				fleetList.Items = fleetList.Items[:len(fleetList.Items)-1]
				break
			}
		}

		if !found {
			pendingCreation = append(pendingCreation, m)
		}
	}

	var errs error

	// Delete any remaining fleets and their associated autoscalers because they are not needed
	for _, fleet := range fleetList.Items {
		mapId := fleet.Labels[config.MapLabel]
		mapsDeleted = append(mapsDeleted, mapId)
		mapLabels, err := metav1.LabelSelectorAsSelector(
			&metav1.LabelSelector{
				MatchLabels: g.config.GetLabels("", mapId),
			},
		)
		if err != nil {
			errs = errors.Join(errs, fmt.Errorf("create label selector for map '%s': %w", mapId, err))
		} else {
			err = g.agones.AutoscalingV1().
				FleetAutoscalers(g.config.GameServerNamespace).
				DeleteCollection(
					ctx,
					metav1.DeleteOptions{},
					metav1.ListOptions{
						LabelSelector: mapLabels.String(),
					},
				)
			if err != nil {
				errs = errors.Join(errs, fmt.Errorf("delete fleet autoscaler for map '%s': %w", mapId, err))
			}

			err = g.agones.AgonesV1().Fleets(g.config.GameServerNamespace).Delete(ctx, fleet.Name, metav1.DeleteOptions{})
			if err != nil {
				errs = errors.Join(errs, fmt.Errorf("delete fleet '%s': %w", fleet.Name, err))
			}
		}
	}

	// Create any missing gameserver fleets and autoscalers
	for _, m := range pendingCreation {
		mapsCreated = append(mapsCreated, m.Id.String())
		err := g.createGameServers(ctx, dimension, m)
		if err != nil {
			errs = errors.Join(errs, fmt.Errorf("create gameserver for map '%s': %w", m.Id.String(), err))
		}
	}

	return mapsCreated, mapsDeleted, errs
}

// DimensionMapChanged implements GameServerManagerService.
func (g *gsmService) DimensionMapChanged(ctx context.Context, dimension *game.Dimension, m *game.Map, created bool) {
	if g.ctx != nil {
		g.DimensionMapsChanged <- dimensionMap{
			Dimension: dimension,
			Map:       m,
			Created:   created,
		}
	} else {
		log.Logger.WithContext(ctx).Warnf("game server manager service not started")
	}
}

// FindAvailableGameServers implements GameServerManagerService.
func (g *gsmService) FindAvailableGameServers(ctx context.Context, dimensionId, mapId string) (*aav1.GameServerAllocation, error) {
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
						err := g.deleteGameServers(g.ctx, data.Dimension, data.Map)
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

// CountGameServers implements GameServerManagerService.
func (g *gsmService) CountGameServers(ctx context.Context) (int, error) {
	fleetList, err := g.agones.AgonesV1().Fleets(g.config.GameServerNamespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("list fleets: %v", err)
		return 0, err
	}

	return len(fleetList.Items), nil
}
