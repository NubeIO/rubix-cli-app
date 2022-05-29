package apps

import (
	"gorm.io/datatypes"
	"gthub.com/NubeIO/rubix-cli-app/service/apps/app"
)

type Store struct {
	UUID              string         `json:"uuid" gorm:"primaryKey"`
	AppName           string         `json:"name"  gorm:"type:varchar(255);unique;not null"`
	Product           string         `json:"product_type"`       // model.ProductType
	AllowableProducts datatypes.JSON `json:"allowable_products"` // All RubixCompute RubixIO
	Port              int
	AppTypeName       string   `json:"app_type_name"` //go, node
	AppType           app.Type `json:"-"`             //go, node

	Repo          string `json:"repo"`         // wires-build
	ServiceName   string `json:"service_name"` // nubeio-rubix-wires
	RubixRootPath string // /data
	AppsPath      string `json:"apps_path"` // /data/rubix-apps/install/flow-framework
	AppPath       string `json:"app_path"`  // /data/flow-framework

	// git details
	Token string `json:"token"`
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
