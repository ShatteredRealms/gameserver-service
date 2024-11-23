package config

import "github.com/ShatteredRealms/go-common-service/pkg/config"

var (
	Version = "v1.0.0"
)

type DimensionConfig struct {
	config.BaseConfig `yaml:",inline" mapstructure:",squash"`
	Postgres          config.DBPoolConfig `yaml:"postgres"`
}

func NewDimensionConfig() *DimensionConfig {
	return &DimensionConfig{
		BaseConfig: config.BaseConfig{
			Server: config.ServerAddress{
				Host: "localhost",
				Port: "8083",
			},
			Keycloak: config.KeycloakConfig{
				BaseURL:      "localhost:8080",
				Realm:        "default",
				Id:           "ae593ef2-49d7-4ca1-8b8b-226f4e95b509",
				ClientId:     "sro-dimension-service",
				ClientSecret: "**********",
			},
			Mode:                "local",
			LogLevel:            0,
			OpenTelemtryAddress: "localhost:4317",
		},
		Postgres: config.DBPoolConfig{
			Master: config.DBConfig{
				ServerAddress: config.ServerAddress{},
				Name:          "dimension-service",
				Username:      "postgres",
				Password:      "password",
			},
		},
	}
}
