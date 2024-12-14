package gindata

import (
	"net/http"
	"strconv"

	"github.com/user/Practice_api/common"
	"github.com/user/Practice_api/modules/items/biz"
	"github.com/user/Practice_api/modules/items/models"
	"github.com/user/Practice_api/modules/items/storgare"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateItemTransport(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.ItemUpdate

		id, err := strconv.Atoi(c.Param("id")) // convert string to int
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester) //bắt profile user khi tạo item

		store := storgare.SqlStore(db)
		business := biz.UpdateItemBiz(store, requester)
		if err := business.UpdateItemByBiz(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, nil, nil))
	}
}
