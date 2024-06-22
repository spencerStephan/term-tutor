package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database DatabasePaths
}

type DatabasePaths struct {
	Windows string `mapstructure:"windows"`
	Mac     string `mapstructure:"mac"`
	Linux   string `mapstructure:"linux"`
}

var Cfg Config

func InitConfig(configPath string) {
	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading in config file, %s", err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("Error unmarshalling config file, %s", err)
	}

}
