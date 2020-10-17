package config

import (
	"context"
	"github.com/heirko/go-contrib/logrusHelper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitLogging(ctx context.Context) error {
	logConfig := viper.New()
	logConfig.SetConfigName("logrus")
	logConfig.SetConfigType("yaml")
	logConfig.AddConfigPath(".")
	err := logConfig.ReadInConfig()
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("could not load logging configuration")
		return err
	}

	var c = logrusHelper.UnmarshalConfiguration(logConfig)
	err = logrusHelper.SetConfig(log.StandardLogger(), c)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("setting up logger configuration failed")
		return err
	}
	return nil
}
