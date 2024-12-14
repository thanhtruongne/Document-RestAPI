package storgare

import (
	"context"

	"github.com/user/Practice_api/modules/items/models"
)

func (s *sqlStruct) UpdateItem(ctx context.Context, conditions map[string]interface{}, model *models.ItemUpdate) error {
	if err := s.db.Table(models.Test{}.TableName()).Where(conditions).Updates(&model).Error; err != nil {
		return err
	}
	return nil
}
