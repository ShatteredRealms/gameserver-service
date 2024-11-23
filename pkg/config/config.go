package config

import (
	"context"

	cconfig "github.com/ShatteredRealms/go-common-service/pkg/config"
)

var (
	Version = "v1.0.0"
)

type DimensionConfig struct {
	cconfig.BaseConfig `yaml:",inline" dimensionstructure:",squash"`
	Postgres           cconfig.DBPoolConfig `yaml:"postgres"`
}

func NewDimensionConfig(ctx context.Context) (*DimensionConfig, error) {
	config := &DimensionConfig{
		BaseConfig: cconfig.BaseConfig{
			Server: cconfig.ServerAddress{
				Host: "localhost",
				Port: "8085",
			},
			Keycloak: cconfig.KeycloakConfig{
				BaseURL:      "localhost:8080",
				Realm:        "default",
				Id:           "7b575e9b-c687-4cdc-b210-67c59b5f380f",
				ClientId:     "sro-dimension-service",
				ClientSecret: "**********",
			},
			Mode:                "local",
			LogLevel:            0,
			OpenTelemtryAddress: "localhost:4317",
		},
		Postgres: cconfig.DBPoolConfig{
			Master: cconfig.DBConfig{
				ServerAddress: cconfig.ServerAddress{},
				Name:          "dimension-service",
				Username:      "postgres",
				Password:      "password",
			},
		},
	}

	err := cconfig.BindConfigEnvs(ctx, "sro-dimension", config)
	return config, err
}
