package biz

import (
	"context"
	"strings"

	"github.com/user/Practice_api/common"
	"github.com/user/Practice_api/modules/items/models"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *models.ItemCreate) error
}

type createItemBiz struct {
	store     CreateItemStorage
	requester common.Requester
}

// tạo biz
func NewCreateItemBiz(store CreateItemStorage, requester common.Requester) *createItemBiz {
	return &createItemBiz{store: store, requester: requester}
}

func (biz *createItemBiz) CreateNewItem(ctx context.Context, data *models.ItemCreate) error {
	name := strings.TrimSpace(data.Name)
	if name == "" {
		return models.ErrTitle
	}
	//get request_user
	ctx_value := context.WithValue(ctx, common.CurrentUser, biz.requester)
	if err := biz.store.CreateItem(ctx_value, data); err != nil {
		return err
	}
	return nil
}
