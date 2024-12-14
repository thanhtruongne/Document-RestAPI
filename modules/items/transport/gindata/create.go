package gindata

import (
	"net/http"

	"github.com/user/Practice_api/common"
	"github.com/user/Practice_api/modules/items/biz"
	"github.com/user/Practice_api/modules/items/models"
	"github.com/user/Practice_api/modules/items/storgare"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateDataItemcompiler(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.ItemCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester) //bắt profile user khi tạo item
		// log.Print("check main: ", requester.GetEmail(), requester.GetRole(), requester.UserId())
		store := storgare.SqlStore(db)
		business := biz.NewCreateItemBiz(store, requester)
		if err := business.CreateNewItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, nil, nil))
	}
}
