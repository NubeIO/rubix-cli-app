package model

import "time"

type Network struct {
	UUID  string  `json:"uuid" gorm:"primary_key"`
	Name  string  `json:"name"  gorm:"type:varchar(255);unique;not null"`
	Hosts []*Host `json:"hosts" gorm:"constraint:OnDelete:CASCADE"`
}

type Host struct {
	UUID                 string    `json:"uuid" gorm:"primaryKey"  get:"true" delete:"true"`
	Name                 string    `json:"name"  gorm:"type:varchar(255);unique;not null" required:"true" get:"true" post:"true" patch:"true"`
	NetworkUUID          string    `json:"network_uuid,omitempty" gorm:"TYPE:varchar(255) REFERENCES networks;not null;default:null"`
	ProductType          string    //edge28, rubix-compute
	IP                   string    `json:"ip" required:"true" default:"192.168.15.10" get:"true" post:"true" patch:"true"`
	Port                 int       `json:"port" required:"true" default:"22" get:"true" post:"true" patch:"true"`
	HTTPS                *bool     `json:"https" get:"true" post:"true" patch:"true"`
	Username             string    `json:"username" required:"true" default:"admin" get:"true" post:"true" patch:"true"`
	Password             string    `json:"password" required:"true" get:"false" post:"true" patch:"true"`
	RubixPort            int       `json:"rubix_port" required:"true" default:"1660" get:"true" post:"true" patch:"true"`
	RubixUsername        string    `json:"rubix_username" required:"true" default:"admin" get:"true" post:"true" patch:"true"`
	RubixPassword        string    `json:"rubix_password" required:"true" post:"true" patch:"true"`
	BiosPort             int       `json:"bios_port" required:"true" default:"1660" get:"true" post:"true" patch:"true"`
	IsLocalhost          *bool     `json:"is_localhost" get:"true" post:"true" patch:"true"`
	PingEnable           *bool     `json:"ping_enable" get:"true" post:"true" patch:"false"`
	PingFrequency        int       `json:"ping_frequency" get:"true" post:"true" patch:"false"`
	IsOffline            *bool     `json:"is_offline" get:"true" post:"false" patch:"false"`
	OfflineCount         uint      `json:"offline_count" get:"true" post:"false" patch:"false"`
	RubixToken           string    `json:"-"`
	RubixTokenLastUpdate time.Time `json:"-"`
	BiosToken            string    `json:"-"`
}
