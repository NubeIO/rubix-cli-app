package rubix

import (
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	pprint "github.com/NubeIO/rubix-edge/pkg/helpers/print"
	"testing"
)

func Test_checkVersion(t *testing.T) {

	homeDir, _ := fileutils.Dir()
	fmt.Println(homeDir)
	app := New(&Rubix{DataDir: "/data", Perm: nonRoot})
	err := app.MakeAllDirs()
	fmt.Println(err)
	if err != nil {
		return
	}

	apps, err := app.InstalledApps()
	if err != nil {
		return
	}
	pprint.PrintJOSN(apps)
	apps, err = app.ConfirmInstalledApps([]string{"nubeio-flow-framework", "non-exists"})
	if err != nil {
		return
	}
	pprint.PrintJOSN(apps)
}
