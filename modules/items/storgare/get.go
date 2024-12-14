package storgare

import (
	"context"

	"github.com/user/Practice_api/common"
	"github.com/user/Practice_api/modules/items/models"
)

func (s *sqlStruct) GetItem(ctx context.Context, conds map[string]interface{}) (*models.Test, error) {
	var data models.Test
	if err := s.db.Table(models.Test{}.TableName()).Where(conds).First(&data).Error; err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrCanNotListEntity(common.EntityName, err)
		}

		return nil, err
	}
	return &data, nil

}
