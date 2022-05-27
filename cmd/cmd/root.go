package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	dbase "gthub.com/NubeIO/rubix-cli-app/database"
	"gthub.com/NubeIO/rubix-cli-app/pkg/config"
	"gthub.com/NubeIO/rubix-cli-app/pkg/database"
)

var (
	//model.Host
	hostName     string
	hostIP       string
	hostPort     int
	hostUsername string
	hostPassword string

	rubixPort     int
	rubixUsername string
	rubixPassword string

	iface string
)

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
	}
}

func initDB() *dbase.DB {
	if err := config.Setup(); err != nil {
		log.Errorln("config.Setup() error: %s", err)
	}
	if err := database.Setup(); err != nil {
		log.Errorln("database.Setup() error: %s", err)
	}
	db := database.GetDB()
	appDB := &dbase.DB{
		DB: db,
	}

	return appDB
}

func init() {

	RootCmd.PersistentFlags().StringVarP(&hostName, "host", "", "RC", "host name (default RC)")
	RootCmd.PersistentFlags().StringVarP(&hostIP, "ip", "", "192.168.15.10", "host ip (default 192.168.15.10)")
	RootCmd.PersistentFlags().IntVarP(&hostPort, "port", "", 22, "SSH Port")
	RootCmd.PersistentFlags().StringVarP(&iface, "iface", "", "", "pc or host network interface example: eth0")
	RootCmd.PersistentFlags().StringVarP(&hostUsername, "host-user", "", "pi", "host/linux username (default pi)")
	RootCmd.PersistentFlags().StringVarP(&hostPassword, "host-pass", "", "N00BRCRC", "host/linux password")
	RootCmd.PersistentFlags().IntVarP(&rubixPort, "rubix-port", "", 1616, "rubix port (default 1616)")
	RootCmd.PersistentFlags().StringVarP(&rubixUsername, "rubix-user", "", "admin", "rubix username (default admin)")
	RootCmd.PersistentFlags().StringVarP(&rubixPassword, "rubix-pass", "", "N00BWires", "rubix password")
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
