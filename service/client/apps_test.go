package client

import (
	"fmt"
	dbase "github.com/NubeIO/rubix-cli-app/database"
	pprint "github.com/NubeIO/rubix-cli-app/pkg/helpers/print"
	"testing"
)

func TestHost(*testing.T) {

	client := New("0.0.0.0", 8090)
	data, res := client.GetApps()
	for i, datum := range data {
		fmt.Println(i, datum)
	}
	install, res := client.InstallApp(&dbase.App{AppName: "flow-framework", Version: "latest", Token: ""})
	fmt.Println(res.GetStatus())
	fmt.Println(res.AsString())
	pprint.PrintJOSN(install)

}
