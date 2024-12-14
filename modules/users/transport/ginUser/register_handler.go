package ginUser

import (
	"net/http"

	"github.com/user/Practice_api/common"
	"github.com/user/Practice_api/modules/users/biz"
	"github.com/user/Practice_api/modules/users/models"
	"github.com/user/Practice_api/modules/users/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserData(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrorInternalServerError(err))
			return
		}

		store := storage.SqlInstance(db)
		md5 := common.NewMd5Hash()
		business := biz.NewRegisterUserBiz(store, md5)

		if err := business.RegiserUserByBuniesness(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))

	}
}
