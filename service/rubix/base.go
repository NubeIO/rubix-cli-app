package rubix

import "fmt"

const nonRoot = 0700
const root = 0777

var FilePerm = root
var DataDir = "/data"
var TmpDir = ""
var AppsInstallDir = ""

type App struct {
	Name        string `json:"app"`          // rubix-wires
	Version     string `json:"version"`      // v1.1.1
	DirName     string `json:"dir_name"`     // wires-builds
	ServiceName string `json:"service_name"` // nubeio-rubix-wires
}

type Rubix struct {
	DataDir string `json:"data_dir"`
	Perm    int
}

func New(r *Rubix) *Rubix {
	if r == nil {
		r = &Rubix{}
	}
	if r.DataDir == "" {
		r.DataDir = DataDir
	}
	if r.Perm == 0 {
		r.Perm = FilePerm
	}
	DataDir = r.DataDir
	TmpDir = fmt.Sprintf("%s/tmp", DataDir)
	AppsInstallDir = fmt.Sprintf("%s/rubix-service/apps/install", DataDir)
	return r
}
