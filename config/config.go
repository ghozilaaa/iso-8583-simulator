package config

import (
	"log"

	"github.com/spf13/viper"
)

const (
	configFileTypeYaml = "yaml"
)

type Config struct {
	App    App
	Server Server
}

type App struct {
	Name    string
	Version string
}

type Server struct {
	Host string
	Port string
}

// ParseConfigFile is used for parsing config file into struct
func ParseConfigFile(configName string) *Config {
	viperConfig := viper.New()
	viperConfig.SetConfigName(configName)
	viperConfig.SetConfigType(configFileTypeYaml)
	viperConfig.AddConfigPath(".")

	if err := viperConfig.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	config := new(Config)
	if err := viperConfig.Unmarshal(&config); err != nil {
		log.Fatalf("failed parsing config, %v", err)
	}

	return config
}
