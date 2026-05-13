package config

import (
	"strings"

	"github.com/spf13/viper"
)

func Load() (*Config, error) {

	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	// Allow env var overrides
	//
	// Example:
	// SERVER_PORT=9090
	//
	viper.SetEnvKeyReplacer(
		strings.NewReplacer(".", "_"),
	)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
