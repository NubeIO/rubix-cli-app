package model

type Apps struct {
	UUID        string `json:"uuid" gorm:"primaryKey"  get:"true" delete:"true"`
	Name        string `json:"name"  gorm:"type:varchar(255);unique;not null"`
	ProductType string `json:"product_type"` // model.ProductType

	// git details
	Token              string `json:"token"`
	Owner              string `json:"owner"`
	Repo               string `json:"repo"`
	Arch               string `json:"arch"`
	Tag                string `json:"tag"`
	DownloadPath       string `json:"download_path"`         //home/user
	DownloadPathSubDir string `json:"download_path_sub_dir"` //home/user /bios

	// Service file
	ServiceName string `json:"service_name"` // nubeio-rubix-bios
	//[Unit]
	Description string `json:"description"`
	After       string `json:"after"` //network.target
	//[Service]
	Type             string `json:"type"` //simple
	User             string `json:"user"`
	WorkingDirectory string `json:"working_directory"`
	ExecStart        string `json:"exec_start"`
	Restart          string `json:"restart"`
	RestartSec       int    `json:"restart_sec"`
	StandardOutput   string `json:"standard_output"`
	StandardError    string `json:"standard_error"`
	SyslogIdentifier string `json:"syslog_identifier"`
	//[Install]
	WantedBy string `json:"wanted_by"` //multi-user.target
}
