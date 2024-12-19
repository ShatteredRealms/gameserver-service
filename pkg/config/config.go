package config

import (
	"context"
	"fmt"
	"strings"

	v1 "agones.dev/agones/pkg/apis/agones/v1"
	autoscalingv1 "agones.dev/agones/pkg/apis/autoscaling/v1"
	"github.com/ShatteredRealms/gameserver-service/pkg/model/game"
	"github.com/ShatteredRealms/go-common-service/pkg/config"
	cconfig "github.com/ShatteredRealms/go-common-service/pkg/config"
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	ManagedLabelKey   = "shatteredrealms.online/managed"
	ManagedLabelValue = "gameserver-service"

	MapLabel       = "shatteredrealms.online/map"
	DimensionLabel = "shatteredrealms.online/dimension"
)

var (
	Version     = "v1.0.0"
	ServiceName = "GameServerService"
)

type GameServerConfig struct {
	cconfig.BaseConfig `yaml:",inline" dimensionstructure:",squash"`
	Postgres           cconfig.DBConfig        `yaml:"postgres"`
	Redis              cconfig.DBPoolConfig    `yaml:"redis"`
	GsmConfig          GameServerManagerConfig `yaml:"gameServerManager"`
}

type GameServerManagerConfig struct {
	GameServerImage       string `yaml:"gameServerImage"`
	FleetPrefix           string `yaml:"fleetPrefix"`
	FleetAutoscalerPrefix string `yaml:"fleetAutoscalerPrefix"`
	GameServerNamespace   string `yaml:"gameServerNamespace"`
}

func NewGameServerConfig(ctx context.Context) (*GameServerConfig, error) {
	config := &GameServerConfig{
		BaseConfig: cconfig.BaseConfig{
			Server: cconfig.ServerAddress{
				Host: "localhost",
				Port: "8082",
			},
			Keycloak: cconfig.KeycloakConfig{
				BaseURL:      "http://localhost:8080",
				Realm:        "default",
				Id:           "ae593ef2-49d7-4ca1-8b8b-226f4e95b509",
				ClientId:     "sro-gameserver-service",
				ClientSecret: "**********",
			},
			Mode:                config.ModeLocal,
			LogLevel:            logrus.InfoLevel,
			OpenTelemtryAddress: "localhost:4317",
			Kafka: cconfig.ServerAddresses{
				{
					Host: "localhost",
					Port: "29092",
				},
			},
		},
		Postgres: cconfig.DBConfig{
			ServerAddress: cconfig.ServerAddress{
				Host: "localhost",
				Port: "5432",
			},
			Name:     "gameserver_service",
			Username: "postgres",
			Password: "password",
		},
		Redis: cconfig.DBPoolConfig{
			Master: cconfig.DBConfig{
				ServerAddress: cconfig.ServerAddress{
					Host: "localhost",
					Port: "7000",
				},
			},
		},
		GsmConfig: GameServerManagerConfig{
			GameServerImage:       "sro-gameserver",
			FleetPrefix:           "sro-f",
			FleetAutoscalerPrefix: "sro-fas",
		},
	}

	err := cconfig.BindConfigEnvs(ctx, "sro-gameserver-service", config)
	return config, err
}

func (c *GameServerManagerConfig) GetFleetTemplate(dimension *game.Dimension, m *game.Map) *v1.Fleet {
	labels := c.GetLabels(dimension.Id.String(), m.Id.String())
	return &v1.Fleet{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      c.getFleetName(dimension, m),
			Namespace: c.GameServerNamespace,
			Labels:    labels,
		},
		Spec: v1.FleetSpec{
			Replicas:   1,
			Scheduling: "",
			Strategy: appsv1.DeploymentStrategy{
				Type: "RollingUpdate",
				RollingUpdate: &appsv1.RollingUpdateDeployment{
					MaxUnavailable: &intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 25,
					},
					MaxSurge: &intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 25,
					},
				},
			},
			Template: v1.GameServerTemplateSpec{
				Spec: v1.GameServerSpec{
					Container: "",
					Ports: []v1.GameServerPort{
						{
							Name:          "default",
							PortPolicy:    "Dynamic",
							ContainerPort: 7777,
						},
					},
					Health: v1.Health{
						Disabled:            false,
						PeriodSeconds:       10,
						FailureThreshold:    3,
						InitialDelaySeconds: 300,
					},
					Template: corev1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Name:      c.getGameServerName(dimension, m),
							Namespace: c.GameServerNamespace,
							Labels:    labels,
						},
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name:  "gameserver",
									Image: c.GameServerImage,
									Args: []string{
										m.MapPath,
										"-log",
									},
									ImagePullPolicy: "Always",
								},
							},
							ImagePullSecrets: []corev1.LocalObjectReference{
								{
									Name: "regcred",
								},
							},
						},
					},
				},
			},
		},
	}
}

func (c *GameServerManagerConfig) GetFleetAutoscalerTemplate(dimension *game.Dimension, m *game.Map) *autoscalingv1.FleetAutoscaler {
	labels := c.GetLabels(dimension.Id.String(), m.Id.String())
	return &autoscalingv1.FleetAutoscaler{
		ObjectMeta: metav1.ObjectMeta{
			Name:      c.getFleetAutoscalerName(dimension, m),
			Namespace: c.GameServerNamespace,
			Labels:    labels,
		},
		Spec: autoscalingv1.FleetAutoscalerSpec{
			FleetName: c.getFleetName(dimension, m),
			Policy: autoscalingv1.FleetAutoscalerPolicy{
				Type: autoscalingv1.BufferPolicyType,
				Buffer: &autoscalingv1.BufferPolicy{
					MaxReplicas: 10,
					MinReplicas: 2,
					BufferSize: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 2,
					},
				},
			},
		},
	}
}

func (c *GameServerManagerConfig) GetLabels(dimensionId string, mapId string) map[string]string {
	labels := map[string]string{
		ManagedLabelKey: ManagedLabelValue,
	}
	if dimensionId != "" {
		labels[DimensionLabel] = dimensionId
	}
	if mapId != "" {
		labels[MapLabel] = mapId
	}
	return labels
}

func (c *GameServerManagerConfig) getGameServerName(dimension *game.Dimension, m *game.Map) string {
	return fmt.Sprintf("%s-%s",
		strings.ReplaceAll(strings.ToLower(dimension.Name), " ", "-"),
		strings.ReplaceAll(strings.ToLower(m.Name), " ", "-"),
	)
}

func (c *GameServerManagerConfig) getFleetName(dimension *game.Dimension, m *game.Map) string {
	return fmt.Sprintf("%s-%s",
		c.FleetPrefix,
		c.getGameServerName(dimension, m),
	)
}

func (c *GameServerManagerConfig) getFleetAutoscalerName(dimension *game.Dimension, m *game.Map) string {
	return fmt.Sprintf("%s-%s",
		c.FleetAutoscalerPrefix,
		c.getGameServerName(dimension, m),
	)
}
