package rubix

import (
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	"testing"
)

func Test_checkVersion(t *testing.T) {

	homeDir, _ := fileutils.Dir()
	fmt.Println(homeDir)
	app := New(&App{DataDir: "/data", Perm: nonRoot, HostDownloadPath: fmt.Sprintf("%s/Downloads", homeDir)})

	version := app.GetAppVersion("wires-builds")

	fmt.Println(version)

	err := app.InstallApp("rubix-wires", "wires-builds", "v2.7.2", nil, "wires-builds-2.7.2.zip")
	fmt.Println(err)
	if err != nil {
		return
	}
	version = app.GetAppVersion("wires-builds")

	fmt.Println(version)

	err = app.InstallApp("rubix-wires", "wires-builds-2.7.3.zip", "v2.7.3", nil, "wires-builds-2.7.3.zip")
	fmt.Println(err)
	if err != nil {
		return
	}

	files, err := app.listFiles(fmt.Sprintf("%s/%s", AppsInstallDir, "wires-builds"))
	fmt.Println(err)
	if err != nil {
		return
	}

	fmt.Println(files)

	version = app.GetAppVersion("wires-builds")

	fmt.Println(version)

}
