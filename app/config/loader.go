package config

import (
	"bytes"
	_ "embed"
	"strings"

	"github.com/spf13/viper"
)

type Database struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
}

type Config struct {
	DB *Database
}

//go:embed default-config.yml
var defaultConfiguration []byte

func Load() (*Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	if err := viper.ReadConfig(bytes.NewBuffer(defaultConfiguration)); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
