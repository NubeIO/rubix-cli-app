package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path"
)

var Config *Configuration
var rootCmd *cobra.Command

type Configuration struct {
	Server   ServerConfiguration
	Gin      GinConfiguration
	Database DatabaseConfiguration
	Path     PathConfiguration
}

func Setup(rootCmd_ *cobra.Command) error {
	rootCmd = rootCmd_
	configuration := &Configuration{}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configuration.GetAbsConfigDir())

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Println(err)
	}
	viper.SetDefault("database.driver", "sqlite")
	viper.SetDefault("database.name", "data.db")
	Config = configuration
	return nil
}
func (conf *Configuration) Prod() bool {
	return rootCmd.PersistentFlags().Lookup("prod").Value.String() == "true"
}

func (conf *Configuration) GetPort() string {
	return rootCmd.PersistentFlags().Lookup("port").Value.String()
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
