package biz

import (
	"context"

	"github.com/user/Practice_api/modules/items/models"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, conds map[string]interface{}) (*models.Test, error)
}

type getItemBiz struct {
	store GetItemStorage
}

// tạo biz
func NewGetItemBiz(store GetItemStorage) *getItemBiz {
	return &getItemBiz{store: store}
}

func (biz *getItemBiz) GetNewItem(ctx context.Context, id int) (*models.Test, error) {

	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}
	return data, nil
}
