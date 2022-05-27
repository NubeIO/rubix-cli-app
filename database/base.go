package dbase

import (
	"errors"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func deleteResponse(query *gorm.DB) (*DeleteMessage, error) {
	msg := &DeleteMessage{
		Message: "delete failed",
	}
	if query.Error != nil {
		return msg, query.Error
	}
	r := query.RowsAffected
	if r == 0 {
		return msg, errors.New("not found")
	}
	msg.Message = "deleted ok"
	return msg, nil
}
