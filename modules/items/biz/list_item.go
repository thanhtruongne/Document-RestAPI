package biz

import (
	"context"

	"github.com/user/Practice_api/common"
	"github.com/user/Practice_api/modules/items/models"
)

type ListItemStorage interface {
	ListItem(ctx context.Context, filter *common.Filter, paging *common.Paging, moreInfos ...string) ([]models.Test, error)
}

type ListItemBiz struct {
	store ListItemStorage
}

func NewListItem(store ListItemStorage) *ListItemBiz {
	return &ListItemBiz{store: store}
}

func (biz *ListItemBiz) ListItemData(ctx context.Context, filter *common.Filter, paging *common.Paging) ([]models.Test, error) {

	data, err := biz.store.ListItem(ctx, filter, paging, "Owner")

	if err != nil {
		return nil, err
	}

	return data, nil
}
