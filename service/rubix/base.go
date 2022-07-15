package rubix

import "fmt"

const nonRoot = 0700
const root = 0777

var FilePerm = nonRoot
var DataDir = "/home/aidan/rubix-edge-testing/data"
var TmpDir = fmt.Sprintf("%s/tmp", DataDir)
var AppsInstallDir = fmt.Sprintf("%s/rubix-service/apps/install", DataDir)
