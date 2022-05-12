package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

var ip string
var modbusIp string
var modbusPort int

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "nube-cli",
	Short: "description",
	Long:  `description`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.baccli.yaml)")
	RootCmd.PersistentFlags().StringVarP(&modbusIp, "modbus-ip", "", "192.168.15.93", "host ip")
	RootCmd.PersistentFlags().IntVarP(&modbusPort, "modbus-port", "", 502, "Port")
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// We want to allow this to be accessed
	viper.BindPFlag("interface", RootCmd.PersistentFlags().Lookup("interface"))
	viper.BindPFlag("port", RootCmd.PersistentFlags().Lookup("port"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".baccli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".baccli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
