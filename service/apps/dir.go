package apps

import (
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"os"
)

func (inst *EdgeApps) MakeTmpUploadDirHome() error {
	if err := emptyPath(inst.TmpUploadDirHome); err != nil {
		return err
	}
	path := fmt.Sprintf("%s/%s", userHomeDir(), inst.TmpUploadDirHome)
	return inst.App.MakeDirectoryIfNotExists(path, os.FileMode(inst.Perm))
}

// MakeTmpUploadDirHomeTmp make a tmp folder for uploading
func (inst *EdgeApps) MakeTmpUploadDirHomeTmp() (string, error) {
	if err := emptyPath(inst.TmpUploadDirHome); err != nil {
		return "", err
	}
	path := fmt.Sprintf("%s/%s/%s", userHomeDir(), inst.TmpUploadDirHome, uuid.ShortUUID("tmp"))
	return path, inst.App.MakeDirectoryIfNotExists(path, os.FileMode(inst.Perm))
}
