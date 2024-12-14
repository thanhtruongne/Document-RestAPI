package storage

import "gorm.io/gorm"

type SqlStore struct {
	db *gorm.DB
}

func SqlInstance(db *gorm.DB) *SqlStore {
	return &SqlStore{db: db}
}
