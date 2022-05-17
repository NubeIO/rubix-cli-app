package bios

import (
	"testing"
)

/*
to run
export GITHUB_TOKEN=YOUR-token
(cd pkg/github && go test -run TestInfo)
*/

func TestDownload(t *testing.T) {
	n := New(&Bios{})
	n.Download()

}
