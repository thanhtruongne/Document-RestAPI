package storgare

import "gorm.io/gorm"

type sqlStruct struct {
	db *gorm.DB
}

func SqlStore(db *gorm.DB) *sqlStruct {
	return &sqlStruct{db: db}
}
