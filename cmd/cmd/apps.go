package cmd

import (
	"github.com/spf13/cobra"
	"gthub.com/NubeIO/rubix-cli-app/service/apps"
)

var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "manage rubix service apps",
	Long:  `do things like install an app, the device must have internet access to download the apps`,
	Run:   runApps,
}

type InstallResp struct {
	RespDownload *apps.RespDownload `json:"response_download"`
	RespBuilder  *apps.RespBuilder  `json:"response_builder"`
	RespInstall  *apps.RespInstall  `json:"response_install"`
}

func runApps(cmd *cobra.Command, args []string) {

	//serviceFile := &builder.SystemDBuilder{}
	//
	//inst := &apps.Installer{
	//	Token:              flgApp.token,
	//	Owner:              flgApp.owner,
	//	Repo:               flgApp.repo,
	//	Arch:               flgApp.arch,
	//	Tag:                flgApp.tag,
	//	DownloadPath:       flgApp.destPath,
	//	DownloadPathSubDir: flgApp.target,
	//	ServiceFile:        serviceFile,
	//}
	//newInstall := apps.New(inst)
	//resp := &InstallResp{}
	//
	////DOWNLOAD
	//download, err := newInstall.GitDownload()
	//resp.RespDownload = download
	//if err != nil {
	//	return
	//}
	////Build service file
	//build, err := newInstall.GenerateServiceFile()
	//resp.RespBuilder = build
	//if err != nil {
	//
	//	return
	//}
	////Install
	//install, err := newInstall.InstallService("nubeio-rubix-bios", "/home/aidan/bios-test/nubeio-rubix-bios.service")
	//resp.RespInstall = install

	//inst := &apps.Installer{
	//	Token:    flgApp.token,
	//	Owner:    flgApp.owner,
	//	Repo:     flgApp.app,
	//	Arch:     flgApp.arch,
	//	Tag:      flgApp.tag,
	//	DestPath: flgApp.destPath,
	//	Target:   flgApp.target,
	//}
	//
	//install := apps.New(inst)
	//
	//downloadInstall, _ := install.Download()
	//pprint.PrintJOSN(downloadInstall)
	//pprint.Print(downloadInstall)

}

var flgApp struct {
	token    string
	owner    string
	repo     string
	arch     string
	tag      string
	destPath string
	target   string
}

func init() {
	RootCmd.AddCommand(appsCmd)
	flagSet := appsCmd.Flags()
	flagSet.StringVar(&flgApp.token, "token", "", "github oauth2 token value (optional)")
	flagSet.StringVarP(&flgApp.owner, "owner", "", "NubeIO", "github repository (OWNER/name)")
	flagSet.StringVarP(&flgApp.repo, "app", "", "rubix-bios", "github repository (owner/NAME)")
	flagSet.StringVar(&flgApp.tag, "tag", "latest", "version of build")
	flagSet.StringVar(&flgApp.destPath, "dest", "/data", "destination path")
	flagSet.StringVar(&flgApp.target, "target", "", "rename destination file (optional)")
	flagSet.StringVar(&flgApp.arch, "arch", "amd64", "arch keyword")

}
