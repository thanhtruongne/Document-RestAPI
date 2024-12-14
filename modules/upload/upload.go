package upload

import (
	"fmt"
	"net/http"

	"github.com/user/Practice_api/common"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UploadImage(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		fileResponse, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		dsn := fmt.Sprintf("static/%s", fileResponse.Filename)
		if err := c.SaveUploadedFile(fileResponse, dsn); err != nil {

		}
		img := common.ImageStruct{
			Id:        0,
			Url:       dsn,
			Width:     300,
			Height:    300,
			Cloudname: "*",
			Extension: "*",
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
