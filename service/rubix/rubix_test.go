package rubix

import (
	"fmt"
	pprint "github.com/NubeIO/rubix-edge/pkg/helpers/print"
	"testing"
)

func Test_checkAppsService(t *testing.T) {

	apps, err := checkAppsDir("")

	fmt.Println(err)

	pprint.PrintJOSN(apps)

	service, err := checkAppsService(apps)

	fmt.Println(err)
	pprint.PrintJOSN(service)

}
