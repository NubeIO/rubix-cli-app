package config

import (
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	"github.com/spf13/viper"

	"github.com/NubeIO/edge/pkg/logger"
)

var Config *Configuration

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
	Path     PathConfiguration
}

// Setup initialize configuration
func Setup() error {

	var configuration *Configuration
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	home, err := fileutils.Dir()
	if err != nil {
		fmt.Println(err)
	}
	viper.AddConfigPath(home + "/rubix-assist-cli")

	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("Error reading config file, %s", err)
		fmt.Println(err)
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		logger.Errorf("Unable to decode into struct, %v", err)
		fmt.Println(err)
	}
	viper.SetDefault("server.port", "1661")
	viper.SetDefault("database.driver", "sqlite")
	viper.SetDefault("database.name", "updater.db")

	Config = configuration
	return nil
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
