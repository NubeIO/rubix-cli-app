package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	pprint "gthub.com/NubeIO/rubix-cli-app/pkg/helpers/print"
	"gthub.com/NubeIO/rubix-cli-app/service/apps"
	"gthub.com/NubeIO/rubix-cli-app/service/apps/app"
	"os"
)

var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "manage rubix service apps",
	Long:  `do things like install an app, the device must have internet access to download the apps`,
	Run:   runApps,
}

type InstallResp struct {
	RespBuilder *apps.RespBuilder `json:"response_builder"`
}

func runApps(cmd *cobra.Command, args []string) {
	var err error

	_, appName, err := app.CheckAppName(flgApp.appName)
	log.Infof("try and install app:%s \n", appName)
	var perm os.FileMode = 0777
	inst := &apps.Apps{
		AppName:       appName,
		Token:         flgApp.token,
		Version:       flgApp.version,
		DownloadPath:  "/home/aidan/apps-test/new",
		RubixRootPath: "/data",
		Perm:          perm,
		ServiceName:   "nubeio-flow-framework",
	}
	newApp, err := apps.New(inst, appName)
	if err != nil {
		log.Errorln("new app: failed to init a new app", err)
		return
	}

	err = inst.MakeDownloadDir()

	_, err = newApp.GitDownload(inst.DownloadPath)
	if err != nil {
		log.Errorf("git: download error %s \n", err.Error())
		return
	}

	err = inst.MakeInstallDir()
	if err != nil {
		return
	}

	err = inst.UnpackBuild()
	if err != nil {
		return
	}

	tmpFileDir := "/tmp"
	_, err = newApp.GenerateServiceFile(newApp.GeneratedApp, tmpFileDir)
	if err != nil {
		log.Errorf("make service file build: failed error:%s \n", err.Error())
		return
	}

	tmpServiceFile := fmt.Sprintf("%s/%s.service", tmpFileDir, newApp.GeneratedApp.ServiceName)
	installService, err := newApp.InstallService(newApp.GeneratedApp.ServiceName, tmpServiceFile)
	if err != nil {
		return
	}

	pprint.PrintJOSN(installService)
	err = inst.CleanUp()
	if err != nil {
		return
	}

	if false {
		// -------------clean up all made dirs-------------
		// -------------remove unzip dir-------------
		//dirs := fileutils.New(&fileutils.Dirs{})
		//err := dirs.RmRF(inst.DownloadPath)
		//if err != nil {
		//	log.Errorf("deleted dir:%s \n", err.Error())
		//	return
		//} else {
		//	log.Infof("deleted dir:%s \n", inst.DownloadPath)
		//}
		//// -------------remove install dir-------------
		//err = dirs.RmRF(installPath)
		//if err != nil {
		//	log.Errorf("deleted dir:%s \n", err.Error())
		//	return
		//} else {
		//	log.Infof("deleted dir:%s \n", installPath)
		//}
	}

}

var flgApp struct {
	token    string
	owner    string
	appName  string
	arch     string
	version  string
	destPath string
	target   string
}

func init() {
	RootCmd.AddCommand(appsCmd)
	flagSet := appsCmd.Flags()
	flagSet.StringVar(&flgApp.token, "token", "", "github oauth2 token value (optional)")
	flagSet.StringVarP(&flgApp.appName, "app", "", "", "rubix-wires, wires or RubixWires")
	flagSet.StringVar(&flgApp.version, "version", "latest", "version of build")
	flagSet.StringVar(&flgApp.destPath, "dest", "/data", "destination path")

}
