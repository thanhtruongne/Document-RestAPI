package storgare

import (
	"context"

	"github.com/user/Practice_api/common"
	"github.com/user/Practice_api/modules/items/models"
)

func (s *sqlStruct) CreateItem(ctx context.Context, data *models.ItemCreate) error {
	requester := ctx.Value(common.CurrentUser).(common.Requester)
	data.User_id = requester.UserId()
	if err := s.db.Table(data.TableName()).Create(&data); err != nil {
		return err.Error
	}
	return nil
}
