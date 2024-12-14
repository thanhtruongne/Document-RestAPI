package storgare

import (
	"context"
	"strconv"

	"github.com/user/Practice_api/common"
	"github.com/user/Practice_api/modules/items/models"
)

func (s *sqlStruct) ListItem(ctx context.Context, filter *common.Filter, paging *common.Paging, moreInfos ...string) ([]models.Test, error) {
	var data []models.Test
	db := s.db.Table(models.Test{}.TableName())

	if fil := filter; fil != nil {
		if status := filter.Status; status != "" {
			v, err := strconv.Atoi(status)
			if err != nil {
				return nil, err
			}
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	//preload
	for v := range moreInfos {
		db = db.Preload(moreInfos[v])
	}

	if err := db.Order("id desc").Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.NewErrorReponse(err, "List item data error", err.Error(), "ENTITY_ERR")
	}
	return data, nil
}
