package client

import (
	"fmt"
	"testing"
)

func TestHost(*testing.T) {

	client := New("0.0.0.0", 8080)

	apps, _ := client.GetApps()
	fmt.Println(222, apps)
	uuid := ""
	fmt.Println(apps)
	for _, app := range apps {
		uuid = app.UUID
	}
	if uuid == "" {
		return
	}

}
