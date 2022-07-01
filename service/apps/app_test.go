package apps

import (
	"fmt"
	"testing"
)

func TestInstall(t *testing.T) {
	inst := &Apps{

		App: &Store{
			ServiceName: "nubeio-flow-framework",
		},
	}

	apps, err := New(inst)

	fmt.Println(apps, err)

	action, err := apps.SystemCtlAction("stop", 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(action)

	// inst := &Apps{
	//	AppName: app.FlowFramework,
	//	//AppName:            "flowFramework", //rubixWires
	//	Token:              "",
	//	Version:            "latest",
	//	DownloadPath:       "/home/aidan/apps-test",
	//	//DownloadPathSubDir: "wires",
	// }
	//
	// apps, err := New(inst)
	// fmt.Println(err)
	// fmt.Println(apps)
	//
	// if err != nil {
	//	fmt.Println("failed", err)
	// }
	//
	// download, err := apps.GitDownload(inst.DownloadPath)
	// if err != nil {
	//	//return
	// }
	// fmt.Println(download)
	// asset := fileutils.New(&fileutils.Dirs{Path: inst.DownloadPath})
	// fmt.Println(download.AssetName)
	// zip := fmt.Sprintf("%s/%s", asset.GetPath(), download.AssetName)
	//
	// //unzip the asset
	// unZip, err := asset.UnZip(zip, "bin", 0777)
	// if err != nil {
	//	//return
	// }
	// fmt.Println(unZip)
	// fmt.Println(err)
}
