package ginUser

import (
	"net/http"

	"github.com/user/Practice_api/common"
	"github.com/user/Practice_api/components/tokenProviders"
	"github.com/user/Practice_api/modules/users/biz"
	"github.com/user/Practice_api/modules/users/models"
	"github.com/user/Practice_api/modules/users/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginUserData(db *gorm.DB, tokenProvider tokenProviders.Provider) func(*gin.Context) {
	return func(c *gin.Context) {
		var loginUserData models.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		// tokenProvider := jwt.NewTokenJWTProvider("jwt", "22123123")
		store := storage.SqlInstance(db)
		md5 := common.NewMd5Hash()

		business := biz.LoginStoreInstance(store, tokenProvider, md5, 60*60*24*30)

		account, err := business.LoginBusiness(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))

	}
}
