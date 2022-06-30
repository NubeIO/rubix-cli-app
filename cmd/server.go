package cmd

import (
	"fmt"
	"github.com/NubeIO/edge/pkg/config"
	"github.com/NubeIO/edge/pkg/database"
	"github.com/NubeIO/edge/pkg/logger"
	"github.com/NubeIO/edge/pkg/router"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "system admin for edge28",
	Long:  "pass in the host name and do operation like check arch type of the host",
	Run:   runServer,
}

func runServer(cmd *cobra.Command, args []string) {
	logger.Init()
	logger.SetLogLevel(logrus.InfoLevel)
	logger.InfoLn("starting edge...")

	if err := config.Setup(RootCmd); err != nil {
		logger.Errorf("config.Setup() error: %s", err)
	}

	if err := os.MkdirAll(config.Config.GetAbsDataDir(), 0755); err != nil {
		panic(err)
	}

	if err := database.Setup(); err != nil {
		logger.Fatalf("database.Setup() error: %s", err)
	}

	r := router.Setup(database.DB)

	host := "0.0.0.0"
	port := config.Config.GetPort()
	logger.Infof("server is starting at %s:%s", host, port)
	log.Fatalf("%v", r.Run(fmt.Sprintf("%s:%s", host, port)))
}

func init() {
	RootCmd.AddCommand(serverCmd)
}
