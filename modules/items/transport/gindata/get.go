package gindata

import (
	"net/http"
	"strconv"

	"github.com/user/Practice_api/common"
	"github.com/user/Practice_api/modules/items/biz"
	"github.com/user/Practice_api/modules/items/storgare"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDataItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id")) // convert string to int
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		store := storgare.SqlStore(db)
		business := biz.NewGetItemBiz(store)
		data, err := business.GetNewItem(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, nil, nil))
	}
}
