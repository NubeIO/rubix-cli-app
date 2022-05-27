package dbhandler

import (
	dbase "gthub.com/NubeIO/rubix-cli-app/database"
)

var db *dbase.DB

type Handler struct {
	DB *dbase.DB
}

//Init give db access
func Init(h *Handler) {
	db = h.DB
}

func GetDB() *dbase.DB {
	return db
}
