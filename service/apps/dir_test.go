package apps

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"testing"
)

func TestEdgeApps_MakeTmpUploadDirHome(t *testing.T) {

	apps, err := New(&EdgeApps{App: &installer.App{}})
	fmt.Println(err)
	err = apps.MakeTmpUploadDirHome()
	fmt.Println(err)
	if err != nil {
		return
	}
}
