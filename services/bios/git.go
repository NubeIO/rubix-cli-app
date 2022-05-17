package bios

import (
	"context"
	"fmt"
	pprint "github.com/NubeIO/git/pkg/helpers/print"

	"github.com/NubeIO/git/pkg/git"
	"os"
)

type Bios struct {
}

func New(bios *Bios) *Bios {
	return bios
}

func (inst *Bios) githubToken() string {
	return os.Getenv("GITHUB_TOKEN")
}

var (
	page    = 1
	perPage = 2
)

func (inst *Bios) Download() {

	opt := &git.AssetOptions{
		Owner:         git.NubeIO,
		Repo:          git.NubeRubixService,
		Tag:           git.TagLatest,
		Arch:          git.ArchAmd64,
		DestPath:      "../../bin",
		ManualInstall: git.ManualInstall{},
	}
	ctx := context.Background()
	client := git.NewClient(inst.githubToken(), opt, ctx)

	resp, err := client.DownloadReleaseAsset()
	if err != nil {
		fmt.Println(err)
	}

	pprint.PrintJOSN(resp)

}

func (inst *Bios) List() {

	opt := &git.AssetOptions{
		Owner:         git.NubeIO,
		Repo:          git.NubeRubixService,
		Tag:           git.TagLatest,
		Arch:          git.ArchAmd64,
		DestPath:      "bin",
		ManualInstall: git.ManualInstall{},
	}

	ctx := context.Background()
	client := git.NewClient(inst.githubToken(), opt, ctx)

	gitOpts := &git.ListOptions{
		Page:    page,
		PerPage: perPage,
	}
	resp, err := client.ListReleases(gitOpts)
	fmt.Println(err)
	if err != nil {

	}

	pprint.PrintJOSN(resp)

}
