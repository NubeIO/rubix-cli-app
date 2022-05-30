package dbase

import (
	"encoding/json"
	"errors"
	"fmt"
	"gthub.com/NubeIO/rubix-cli-app/pkg/helpers/uuid"
	"gthub.com/NubeIO/rubix-cli-app/pkg/logger"
	"gthub.com/NubeIO/rubix-cli-app/service/apps"
	"gthub.com/NubeIO/rubix-cli-app/service/product"
)

func (d *DB) GetApps() ([]*apps.Store, error) {
	var m []*apps.Store
	if err := d.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) GetApp(uuid string) (*apps.Store, error) {
	var m *apps.Store
	if err := d.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		logger.Errorf("GetApp error: %v", err)
		return nil, err
	}
	return m, nil
}

func (d *DB) GetAppByName(name string) (*apps.Store, error) {
	var m *apps.Store
	if err := d.DB.Where("name = ? ", name).First(&m).Error; err != nil {
		logger.Errorf("GetApp error: %v", err)
		return nil, err
	}
	return m, nil
}

func checkProduct(products []string) error {
	if len(products) == 0 {
		return errors.New("product list can not be empty try, RubixCompute, Edge28")
	}
	for _, pro := range products {
		_, err := product.CheckProduct(pro)
		if err != nil {
			return err
		}
	}
	return nil
}

const (
	Owner   = "NubeIO"
	User    = "root"
	TempDir = "/tmp"
)

func validateAllowableProducts(store *apps.Store) ([]byte, []string, error) {
	if store.AllowableProducts == nil {
		return nil, nil, nil
	}
	var data []string
	err := json.Unmarshal(store.AllowableProducts, &data)
	if err != nil {
		return nil, nil, err
	}
	valid, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}
	return valid, data, nil
}

func (d *DB) CreateApp(body *apps.Store) (*apps.Store, error) {
	body.UUID = fmt.Sprintf("app_%s", uuid.SmallUUID())
	pro, err := product.Get()
	appType, appTypeName, err := apps.CheckType(body.AppTypeName)
	if err != nil {
		return nil, err
	}
	products, productsList, err := validateAllowableProducts(body)
	if err != nil {
		return nil, err
	}
	err = checkProduct(productsList)
	if err != nil {
	}

	if body.RubixRootPath == "" {
		body.RubixRootPath = "/data"
	}
	if body.AppsPath == "" {
		body.AppsPath = "/data/rubix-apps/installed"
	}

	if body.Owner == "" {
		body.Owner = Owner
	}
	if body.RunAsUser == "" {
		body.RunAsUser = User
	}
	if body.DownloadPath == "" {
		body.DownloadPath = TempDir
	}

	body.AllowableProducts = products
	body.AppTypeName = appTypeName
	body.AppType = appType
	get, err := product.Get()
	if err != nil {
		return nil, err
	}
	body.Arch = get.Arch
	if err != nil {
		return nil, err
	}
	body.ProductType = pro.Type
	if err != nil {
		return nil, err
	}
	fmt.Println(body)
	if err := d.DB.Create(&body).Error; err != nil {
		return nil, err
	} else {
		return body, nil
	}
}

func (d *DB) UpdateApp(uuid string, app *apps.Store) (*apps.Store, error) {
	var m *apps.Store
	query := d.DB.Where("uuid = ?", uuid).Find(&m).Updates(app)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return app, query.Error
	}
}

func (d *DB) DeleteApp(uuid string) (*DeleteMessage, error) {
	var m *apps.Store
	query := d.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

type DeleteMessage struct {
	Message string `json:"message"`
}

// DropApps delete all.
func (d *DB) DropApps() (*DeleteMessage, error) {
	var m *apps.Store
	query := d.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
