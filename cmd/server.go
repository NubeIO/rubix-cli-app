package cmd

import (
	"fmt"
	"github.com/NubeIO/rubix-edge/pkg/config"
	"github.com/NubeIO/rubix-edge/pkg/logger"
	"github.com/NubeIO/rubix-edge/pkg/router"
	"github.com/spf13/cobra"
	"os"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "system admin for edge28",
	Long:  "pass in the host name and do operation like check arch type of the host",
	Run:   runServer,
}

func runServer(cmd *cobra.Command, args []string) {
	if err := config.Setup(RootCmd); err != nil {
		logger.Logger.Errorf("config.Setup() error: %s", err)
	}
	if err := os.MkdirAll(config.Config.GetAbsDataDir(), 0755); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(config.Config.GetAbsTempDir(), 0755); err != nil {
		panic(err)
	}
	logger.Init()
	logger.Logger.Infoln("starting edge...")

	r := router.Setup()

	host := "0.0.0.0"
	port := config.Config.GetPort()
	logger.Logger.Infof("server is starting at %s:%s", host, port)
	logger.Logger.Fatalf("%v", r.Run(fmt.Sprintf("%s:%s", host, port)))
}

func init() {
	RootCmd.AddCommand(serverCmd)
}
