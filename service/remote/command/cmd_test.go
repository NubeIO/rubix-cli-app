package command

import (
	"fmt"
	"strings"
	"testing"
)

func TestCMD(t *testing.T) {
	cmd := &Command{SetPath: "/home/aidan", Commands: []string{"ls"}}
	str := []string{"/home/pi", "ls"}
	fmt.Println(strings.Join(str[:], " "))
	out := cmd.RunCommand()
	fmt.Println(out.Out)

}
