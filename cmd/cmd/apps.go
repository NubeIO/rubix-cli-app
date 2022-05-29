package cmd

import (
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gthub.com/NubeIO/rubix-cli-app/service/apps"
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
	RespInstall *apps.RespInstall `json:"response_install"`
}

func runApps(cmd *cobra.Command, args []string) {

	name := flgApp.appName
	log.Infof("try and install app:%s \n", name)
	var perm os.FileMode = 0777
	inst := &apps.Apps{
		AppName:             name,
		Token:               flgApp.token,
		Version:             flgApp.version,
		DownloadPath:        "/home/aidan/apps-test/new",
		OverrideInstallPath: "/home/aidan/apps-test/new",
	}

	dirs := fileutils.New(&fileutils.Dirs{})
	// -------------check download dir-------------
	if !dirs.DirExists(inst.DownloadPath) {
		log.Errorf("no dir exists %s \n", inst.DownloadPath)
		err := dirs.MkdirAll(inst.DownloadPath, perm)
		if err != nil {
			log.Errorf("unzip build: failed to make new dir %s \n", inst.DownloadPath)
			return
		}
		log.Infof("unzip build: made new dir:%s \n", inst.DownloadPath)
	} else {
		log.Infof("unzip build: existing dir to download zip:%s \n", inst.DownloadPath)
	}

	newApp, err := apps.New(inst)
	if err != nil {
		log.Errorln("new app: failed to init a new app", err)
		return
	}

	// -------------download build-------------

	download, err := newApp.GitDownload(inst.DownloadPath)
	if err != nil {
		log.Errorf("git: download error %s \n", err.Error())
		return
	}
	log.Infof("downloaded app name:%s  asset name:%s \n", name, download.AssetName)

	// -------------check build unzip path-------------

	zip := fmt.Sprintf("%s/%s", inst.DownloadPath, download.AssetName)
	installPath := "/home/aidan/flow-framework"
	if !dirs.DirExists(installPath) {
		log.Errorf("no dir exists %s \n", installPath)
		err := dirs.MkdirAll(installPath, perm)
		if err != nil {
			log.Errorf("install dir: failed to make new dir %s \n", installPath)
			return
		}
	} else {
		log.Infof("install dir: existing install dir existed:%s \n", installPath)
	}

	// -------------unzip build-------------
	_, err = dirs.UnZip(zip, installPath, perm)
	if err != nil {
		log.Errorf("unzip build: failed to make new dir %s \n", installPath)
		return
	} else {
		log.Infof("unzip build: existing install dir existed:%s \n", installPath)
	}

	if false {
		// -------------clean up all made dirs-------------
		// -------------remove unzip dir-------------
		err := dirs.RmRF(inst.DownloadPath)
		if err != nil {
			log.Errorf("deleted dir:%s \n", err.Error())
			return
		} else {
			log.Infof("deleted dir:%s \n", inst.DownloadPath)
		}
		// -------------remove install dir-------------
		err = dirs.RmRF(installPath)
		if err != nil {
			log.Errorf("deleted dir:%s \n", err.Error())
			return
		} else {
			log.Infof("deleted dir:%s \n", installPath)
		}
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
