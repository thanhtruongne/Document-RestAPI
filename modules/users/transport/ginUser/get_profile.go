package ginUser

import (
	"net/http"

	"github.com/user/Practice_api/common"

	"github.com/gin-gonic/gin"
)

func Profile() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(user))
	}
}
