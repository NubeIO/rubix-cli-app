package cmd

import (
	"fmt"
	"github.com/NubeIO/edge/pkg/config"
	"github.com/NubeIO/edge/pkg/database"
	"github.com/NubeIO/edge/pkg/logger"
	"github.com/NubeIO/edge/pkg/router"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "system admin for edge28",
	Long:  `pass in the host name and do operation like check arch type of the host`,
	Run:   runServer,
}

func setup() {
	logger.Init()
	logger.SetLogLevel(logrus.InfoLevel)
	logger.InfoLn("try and start rubix-updater")
	if err := config.Setup(); err != nil {
		logger.Errorf("config.Setup() error: %s", err)
	}
	if err := database.Setup(); err != nil {
		logger.Fatalf("database.Setup() error: %s", err)
	}
}

func runServer(cmd *cobra.Command, args []string) {
	setup()
	db := database.GetDB()
	r := router.Setup(db)

	host := "0.0.0.0"
	if h := viper.GetString("server.host"); h != "" {
		host = h
	}
	port := RootCmd.PersistentFlags().Lookup("port").Value.String()
	logger.Infof("Server is starting at %s:%s", host, port)
	fmt.Printf("server is running at %s:%s Check logs for details\n", host, port)
	log.Fatalf("%v", r.Run(fmt.Sprintf("%s:%s", host, port)))
}

func init() {
	RootCmd.AddCommand(serverCmd)
}
