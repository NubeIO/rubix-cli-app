package cmd

import (
	"github.com/spf13/cobra"
)

var (
//model.Host

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

var flgRoot struct {
	hostName      string
	hostIP        string
	hostPort      int
	hostUsername  string
	hostPassword  string
	rubixPort     int
	rubixUsername string
	rubixPassword string
	iface         string
}

func init() {

	RootCmd.PersistentFlags().StringVarP(&flgRoot.hostName, "host", "", "RC", "host name (default RC)")
	RootCmd.PersistentFlags().StringVarP(&flgRoot.hostIP, "ip", "", "192.168.15.10", "host ip (default 192.168.15.10)")
	RootCmd.PersistentFlags().IntVarP(&flgRoot.hostPort, "port", "", 22, "SSH Port")
	RootCmd.PersistentFlags().StringVarP(&flgRoot.iface, "iface", "", "", "pc or host network interface example: eth0")
	RootCmd.PersistentFlags().StringVarP(&flgRoot.hostUsername, "host-user", "", "pi", "host/linux username (default pi)")
	RootCmd.PersistentFlags().StringVarP(&flgRoot.hostPassword, "host-pass", "", "N00BRCRC", "host/linux password")
	RootCmd.PersistentFlags().IntVarP(&flgRoot.rubixPort, "rubix-port", "", 1616, "rubix port (default 1616)")
	RootCmd.PersistentFlags().StringVarP(&flgRoot.rubixUsername, "rubix-user", "", "admin", "rubix username (default admin)")
	RootCmd.PersistentFlags().StringVarP(&flgRoot.rubixPassword, "rubix-pass", "", "N00BWires", "rubix password")
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
