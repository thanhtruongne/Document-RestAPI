package middleware

import (
	"log"
	"net/http"

	"github.com/user/Practice_api/common"

	"github.com/gin-gonic/gin"
)

func Recovery() func(*gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					log.Print("error: ", appErr)
					c.AbortWithStatusJSON(http.StatusBadRequest, appErr)
					return
				}

				appErr := common.ErrorInternalServerError(err.(error)) // trả kiểu lỗi theo dạng khác k có trong thư viện
				log.Print("error 2: ", appErr)
				c.AbortWithStatusJSON(http.StatusBadRequest, appErr)
				return
			}
		}()

		c.Next()
	}
}
