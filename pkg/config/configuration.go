package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path"

	"github.com/NubeIO/edge/pkg/logger"
)

var Config *Configuration
var rootCmd *cobra.Command

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
	Path     PathConfiguration
}

// Setup initialize configuration
func Setup(rootCmd_ *cobra.Command) error {
	rootCmd = rootCmd_
	var configuration *Configuration
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("Error reading config file, %s", err)
		fmt.Println(err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		logger.Errorf("Unable to decode into struct, %v", err)
		fmt.Println(err)
	}
	viper.SetDefault("database.driver", "sqlite")
	viper.SetDefault("database.name", "data.db")
	Config = configuration
	return nil
}

func (conf *Configuration) GetAbsDataDir() string {
	return path.Join(conf.getGlobalDir(), conf.getDataDir())
}

func (conf *Configuration) GetAbsConfigDir() string {
	return path.Join(conf.getGlobalDir(), conf.getConfigDir())
}

func (conf *Configuration) getGlobalDir() string {
	return rootCmd.PersistentFlags().Lookup("global-dir").Value.String()
}

func (conf *Configuration) getDataDir() string {
	return rootCmd.PersistentFlags().Lookup("data-dir").Value.String()
}

func (conf *Configuration) getConfigDir() string {
	return rootCmd.PersistentFlags().Lookup("config-dir").Value.String()
}
