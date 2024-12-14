package biz

import (
	"context"

	"github.com/user/Practice_api/common"
	"github.com/user/Practice_api/modules/items/models"
)

type UpdateItemStorgate interface {
	GetItem(ctx context.Context, conds map[string]interface{}) (*models.Test, error)
	UpdateItem(ctx context.Context, conditions map[string]interface{}, data *models.ItemUpdate) error
}

type NewUpdateItemBiz struct {
	store     UpdateItemStorgate
	requester common.Requester
}

func UpdateItemBiz(store UpdateItemStorgate, requester common.Requester) *NewUpdateItemBiz {
	return &NewUpdateItemBiz{store: store, requester: requester}
}

func (biz *NewUpdateItemBiz) UpdateItemByBiz(ctx context.Context, id int, data *models.ItemUpdate) error {
	user, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if user.Status == 0 {
		return models.ErrHasBeenDeleted
	}

	if id != biz.requester.UserId() {
		return models.ErrInvalidRequest
	}

	if err := biz.store.UpdateItem(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}
	return nil
}
