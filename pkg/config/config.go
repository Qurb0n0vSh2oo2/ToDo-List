package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type DbConfig struct {
	Databse  string `json:"database"`
	Username string `json:"user"`
	Password string `json:"password"`
}

type Config struct {
	Db   DbConfig `json:"db"`
	Port string   `json:"port"`
	Host string   `json:"host"`
}

var Module = fx.Provide(LoadConfig)

func LoadConfig() (config *Config) {
	viper.SetConfigFile("pkg/config/config.json")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return
}

