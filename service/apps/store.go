package apps

import (
	"gorm.io/datatypes"
)

type InstalledApp struct {
	UUID string `json:"uuid" gorm:"primaryKey"`
	//Name         string `json:"name"  gorm:"type:varchar(255);unique;not null"`
	AppStoreName string `json:"app_store_name"`
	AppStoreUUID string `json:"app_store_uuid"`
}

type Store struct {
	UUID                    string         `json:"uuid" gorm:"primaryKey"`
	Name                    string         `json:"name"  gorm:"type:varchar(255);unique;not null"`
	AllowableProducts       datatypes.JSON `json:"allowable_products"` // All RubixCompute RubixIO
	Port                    int            `json:"port"`
	AppTypeName             string         `json:"app_type_name"`                                          //go, node
	AppType                 Type           `json:"-"`                                                      //go, node
	Repo                    string         `json:"repo"  gorm:"type:varchar(255);unique;not null"`         // wires-build
	ServiceName             string         `json:"service_name"  gorm:"type:varchar(255);unique;not null"` // nubeio-rubix-wires
	RubixRootPath           string         `json:"rubix_root_path"`                                        // /data
	AppsPath                string         `json:"apps_path"`                                              // /data/rubix-apps/install/flow-framework
	AppPath                 string         `json:"app_path"`                                               // /data/flow-framework
	DownloadPath            string         `json:"download_path"`                                          // home/user/downloads
	AssetZipName            string         `json:"-"`                                                      // (auto added)
	Owner                   string         `json:"owner"`                                                  // NubeIO
	RunAsUser               string         `json:"run_as_user"`                                            // root
	ServiceDescription      string         `json:"description"`                                            // nube-io app
	ServiceWorkingDirectory string         `json:"service_working_directory"`                              // MainDir/apps/install/
	ServiceExecStart        string         `json:"service_exec_start"`                                     // npm run prod:start --prod -- --datadir /data/rubix-wires/data --envFile /data/rubix-wires/config/.env
	ProductType             string         `json:"product_type"`                                           // RubixCompute (auto added)
	Arch                    string         `json:"arch"`                                                   // amd64 (auto added)

}

type InstallOptions struct {
	UpgradeToLatest *bool  `json:"upgrade_to_latest"`
	InstallVersion  string `json:"install_version"`
	CleanInstall    *bool  `json:"clean_install"`
}

type UnInstallOptions struct {
	UnInstall      *bool  `json:"upgrade_to_latest"`
	CleanUnInstall *bool  `json:"clean_un_install"`
	Downgrade      string `json:"downgrade"`
}

type ServiceActionOptions struct {
	TimeOut int `json:"time_out"`
}

type AppService struct {
	AppStoreUUID     string            `json:"app_store_uuid"`
	InstallOptions   *InstallOptions   `json:"install_options"`
	UnInstallOptions *UnInstallOptions `json:"un_install_options"`
	ServiceAction    string            `json:"service_action"` // start, stop
}
