package config

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Feeds []feed
}

type feed struct {
	Address string
}

func InitApplicationConfiguration(ctx context.Context) (*Config, error) {
	appConfig := viper.New()
	appConfig.SetConfigName("config")
	appConfig.SetConfigType("yaml")
	appConfig.AddConfigPath(".")
	err := appConfig.ReadInConfig()
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("could not load application configuration")
		return nil, err
	}
	var config Config
	err = appConfig.Unmarshal(&config)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("invalid config format")
		return nil, err
	}

	log.Tracef("config: %+v", config)

	return &config, nil
}
