package model

//install rubix-wires
//install flow

type Apps struct {
	UUID        string `json:"uuid" gorm:"primaryKey"  get:"true" delete:"true"`
	AppName     string `json:"name"  gorm:"type:varchar(255);unique;not null"`
	ProductType string `json:"product_type"` // model.ProductType
	Port        int

	// git details
	Token string `json:"token"`
}
