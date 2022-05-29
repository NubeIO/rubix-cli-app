package cmd

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	pprint "gthub.com/NubeIO/rubix-cli-app/pkg/helpers/print"
	"gthub.com/NubeIO/rubix-cli-app/service/apps"
)

var storesCmd = &cobra.Command{
	Use:   "store",
	Short: "manage rubix app store",
	Long:  `do things like install an app, the device must have internet access to download the apps`,
	Run:   runStore,
}

func runStore(cmd *cobra.Command, args []string) {
	var err error

	db := initDB()
	products, _ := json.Marshal([]string{"RubixCompute"})

	store := &apps.Store{
		AppName:           "ff",
		AppTypeName:       "Go",
		AllowableProducts: products,
	}

	app, err := db.CreateApp(store)
	if err != nil {
		log.Errorln(err)
		return
	}

	if true {
		db.DeleteApp(app.UUID)
	}

	pprint.PrintJOSN(app)

}

var flgStore struct {
	token         string
	owner         string
	appName       string
	arch          string
	version       string
	downloadPath  string
	rubixRootPath string
}

func init() {
	RootCmd.AddCommand(storesCmd)
	flagSet := storesCmd.Flags()
	flagSet.StringVar(&flgStore.token, "token", "", "github oauth2 token value (optional)")
	flagSet.StringVarP(&flgStore.appName, "app", "", "", "RubixWires")
	flagSet.StringVar(&flgStore.version, "version", "latest", "version of build")
	flagSet.StringVar(&flgStore.downloadPath, "download", "/tmp", "download path")
	flagSet.StringVar(&flgStore.rubixRootPath, "rubix-path", "/data", "rubix main path")

}
