package apps

import (
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestInstall(t *testing.T) {
	inst := &Apps{
		AppName: "rubixWires",
		//AppName:            "flowFramework",
		Token:              "",
		Version:            "latest",
		DownloadPath:       "/home/aidan/apps-test",
		DownloadPathSubDir: "wires",
	}

	apps, err := New(inst)
	fmt.Println(err)
	fmt.Println(apps)

	if err != nil {
		log.Errorln("failed", err)
	}

	download, err := apps.GitDownload(inst.DownloadPath)
	if err != nil {
		//return
	}
	fmt.Println(download)
	asset := fileutils.New(&fileutils.Dirs{Path: inst.DownloadPath})
	fmt.Println(download.AssetName)
	zip := fmt.Sprintf("%s/%s", asset.GetPath(), download.AssetName)

	//unzip the asset
	unZip, err := asset.UnZip(zip, "bin", 0777)
	if err != nil {
		//return
	}
	fmt.Println(unZip)
	fmt.Println(err)

}
