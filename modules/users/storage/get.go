package storage

import (
	"context"

	"github.com/user/Practice_api/common"
	"github.com/user/Practice_api/modules/users/models"

	"gorm.io/gorm"
)

func (s *SqlStore) FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*models.User, error) {
	db := s.db.Table(models.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user models.User

	if err := db.Where(condition).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrorDB(err)
	}
	return &user, nil

}
