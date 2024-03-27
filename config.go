package main

import (
	"github.com/spf13/viper"
)

type AssetDashConfig struct {
	Token string `mapstructure:"token"`
}

func setupConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
}

func NewAssetDashConfig() (*AssetDashConfig, error) {
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := AssetDashConfig{}
	if err := viper.UnmarshalExact(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
