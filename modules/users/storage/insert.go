package storage

import (
	"context"

	"github.com/user/Practice_api/common"
	"github.com/user/Practice_api/modules/users/models"
)

func (s *SqlStore) CreateUser(ctx context.Context, data *models.UserCreate) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(&data).Error; err != nil {
		db.Rollback()
		return common.ErrorDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrorDB(err)
	}

	return nil

}
