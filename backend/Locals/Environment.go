package Locals

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
)

type ConfigField struct {
	ThisDbHost      string `mapstructure:"THIS_DB_HOST"`
	ThisDbPass      string `mapstructure:"THIS_DB_PASS"`
	ThisDbUser      string `mapstructure:"THIS_DB_USER"`
	ThisRedisHost   string `mapstructure:"THIS_REDIS_HOST"`
	ThisRedisDbPass string `mapstructure:"THIS_REDIS_DB_PASS"`
	JwtKey          string `mapstructure:"THIS_JWT_KEY"`
	DbName          string `mapstructure:"THIS_DB_NAME"`
	ThisDbConfig    string `mapstructure:"THIS_DB_CONFIG"`
}

// Config Load data for all config
func Config() (config ConfigField, err error) {
	ConfigPath, err := filepath.Abs(".")
	viper.AddConfigPath(ConfigPath + "/backend")
	viper.AddConfigPath(ConfigPath)
	viper.AddConfigPath("$HOME/.ur/backend")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// External Default Config
	config.JwtKey = "testtoken"
	config.DbName = "urassets_db"

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	config.ThisDbConfig = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.ThisDbHost, config.ThisDbUser, config.ThisDbPass, config.DbName) // initiate db config

	return
}
