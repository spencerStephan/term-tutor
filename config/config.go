package config

import (
	"log"

	"github.com/spf13/viper"
)

type RootConfig struct {
	Database DatabasePaths
}

type DatabasePaths struct {
	Windows string `mapstructure:"windows"`
	Mac     string `mapstructure:"mac"`
	Linux   string `mapstructure:"linux"`
}

var RootCfg RootConfig

func InitRootConfig(configPath string) {
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

	if err := viper.Unmarshal(&RootCfg); err != nil {
		log.Fatalf("Error unmarshalling config file, %s", err)
	}
}

func OverrideDatabasePaths(newPath string) {
	if newPath != "" {
		RootCfg.Database.Windows = newPath
		RootCfg.Database.Mac = newPath
		RootCfg.Database.Linux = newPath
	}
}
