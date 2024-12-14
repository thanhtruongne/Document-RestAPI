package gindata

import (
	"net/http"

	"github.com/user/Practice_api/common"
	"github.com/user/Practice_api/modules/items/biz"
	"github.com/user/Practice_api/modules/items/storgare"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListItemData(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging
		var filter common.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		paging.Process()

		store := storgare.SqlStore(db)
		business := biz.NewListItem(store)
		rows, err := business.ListItemData(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		for i := range rows {
			rows[i].Mask()
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(rows, paging, filter))

	}
}
