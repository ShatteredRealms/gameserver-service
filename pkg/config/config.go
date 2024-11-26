package config

import (
	"context"

	cconfig "github.com/ShatteredRealms/go-common-service/pkg/config"
	"github.com/sirupsen/logrus"
)

var (
	Version     = "v1.0.0"
	ServiceName = "GameServerService"
)

type DimensionConfig struct {
	cconfig.BaseConfig `yaml:",inline" dimensionstructure:",squash"`
	Postgres           cconfig.DBPoolConfig `yaml:"postgres"`
	Redis              cconfig.DBPoolConfig `yaml:"redis"`
	GameServerImage    string               `yaml:"gameServerImage"`
}

func NewDimensionConfig(ctx context.Context) (*DimensionConfig, error) {
	config := &DimensionConfig{
		BaseConfig: cconfig.BaseConfig{
			Server: cconfig.ServerAddress{
				Host: "localhost",
				Port: "8084",
			},
			Keycloak: cconfig.KeycloakConfig{
				BaseURL:      "http://localhost:8080",
				Realm:        "default",
				Id:           "ae593ef2-49d7-4ca1-8b8b-226f4e95b509",
				ClientId:     "sro-gameserver-service",
				ClientSecret: "**********",
			},
			Mode:                "local",
			LogLevel:            logrus.DebugLevel,
			OpenTelemtryAddress: "localhost:4317",
			Kafka: cconfig.ServerAddresses{
				{
					Host: "localhost",
					Port: "29092",
				},
			},
		},
		Postgres: cconfig.DBPoolConfig{
			Master: cconfig.DBConfig{
				ServerAddress: cconfig.ServerAddress{
					Host: "localhost",
					Port: "5432",
				},
				Name:     "gameserver_service",
				Username: "postgres",
				Password: "password",
			},
		},
		Redis: cconfig.DBPoolConfig{
			Master: cconfig.DBConfig{
				ServerAddress: cconfig.ServerAddress{
					Host: "localhost",
					Port: "7000",
				},
			},
		},
		GameServerImage: "sro-gameserver",
	}

	err := cconfig.BindConfigEnvs(ctx, "sro-gameserver-service", config)
	return config, err
}
