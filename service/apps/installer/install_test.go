package installer

import (
	"fmt"
	"testing"
)

func Test_extractRepoName(t *testing.T) {

	//slices.Contains(things, "foo") // true
	match, count, version, archMatch, arch := matchRepoName("rubix-plat-build-2.7.0.zip", "rubix-plat-build")
	fmt.Println(match, count, version, archMatch, arch)
	match, count, version, archMatch, arch = matchRepoName("flow-framework-0.5.5-1575cf89.amd64.zip", "rubix-plat-build")
	fmt.Println(match, count, version, archMatch, arch)
	match, count, version, archMatch, arch = matchRepoName("flow-framework-0.5.5-1575cf89.amd32.zip", "flow-framework")
	fmt.Println(match, count, version, archMatch, arch)

}
