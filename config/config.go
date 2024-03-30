package config

import (
	"github.com/leenzstra/assetdash-go/client"
	"github.com/spf13/viper"
)

type Config struct {
	Token     string            `mapstructure:"token"`
	Endpoints *client.Endpoints `mapstructure:"endpoints"`
}

func SetupConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
}

func New() (*Config, error) {
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := viper.UnmarshalExact(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
