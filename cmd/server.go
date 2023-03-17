package cmd

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/rubix-edge/pkg/config"
	"github.com/NubeIO/rubix-edge/pkg/logger"
	"github.com/NubeIO/rubix-edge/pkg/router"
	"github.com/NubeIO/rubix-edge/service/system"
	"github.com/NubeIO/rubix-edge/utils"
	"github.com/go-co-op/gocron"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "system admin for edge28",
	Long:  "pass in the host name and do operation like check arch type of the host",
	Run:   runServer,
}

func setupCron() (*gocron.Scheduler, *systemctl.SystemCtl, *system.System) {
	scheduler := gocron.NewScheduler(time.Local)
	systemCtl := systemctl.New(false, 30)
	system_ := system.New(&system.System{})
	restartJobs := utils.GetRestartJobs()
	for _, restartJob := range restartJobs {
		_, err := cron.ParseStandard(restartJob.Expression)
		if err != nil {
			logger.Logger.Errorln(err)
		} else {
			_, err = scheduler.Cron(restartJob.Expression).Tag(restartJob.Unit).Do(func() {
				_ = systemCtl.Restart(restartJob.Unit)
			})
			if err != nil {
				logger.Logger.Errorln(err)
			}
		}
	}

	rebootJob := utils.GetRebootJob()
	if rebootJob != nil {
		_, err := cron.ParseStandard(rebootJob.Expression)
		if err != nil {
			logger.Logger.Errorln(err)
		}
		_, err = scheduler.Cron(rebootJob.Expression).Tag(rebootJob.Tag).Do(func() {
			_, _ = system_.RebootHost()
		})
	}
	scheduler.StartAsync()
	return scheduler, systemCtl, system_
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
	scheduler, systemCtl, system_ := setupCron()
	logger.Logger.Infoln("starting edge...")

	r := router.Setup(scheduler, systemCtl, system_)

	host := "0.0.0.0"
	port := config.Config.GetPort()
	logger.Logger.Infof("server is starting at %s:%s", host, port)
	logger.Logger.Fatalf("%v", r.Run(fmt.Sprintf("%s:%s", host, port)))
}

func init() {
	RootCmd.AddCommand(serverCmd)
}
