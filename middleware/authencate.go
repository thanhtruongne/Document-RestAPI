package middleware

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/user/Practice_api/common"
	"github.com/user/Practice_api/components/tokenProviders"
	"github.com/user/Practice_api/modules/users/models"

	"github.com/gin-gonic/gin"
)

type Authencate interface {
	FindUser(ctx context.Context, condition map[string]interface{}, modelInfo ...string) (*models.User, error)
}

func ErrAuthenHeader(err error) *common.AppError {
	return common.NewErrorCustomResponse(
		err,
		fmt.Sprintf("wrong authen token header"),
		fmt.Sprintf("ErrWrongAuthenTokenHeader"),
	)
}

func parseAuthenHeaderToken(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrAuthenHeader(nil)
	}

	return parts[1], nil
}

func AuthenCationRequried(authStore Authencate, tokenProviders tokenProviders.Provider) func(c *gin.Context) {
	return func(c *gin.Context) {
		//check token
		token, err := parseAuthenHeaderToken(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		payload, err := tokenProviders.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := authStore.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId()})
		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrorInvalidRequest(errors.New("User doesn't have permission")))
		}
		c.Set(common.CurrentUser, user)
		c.Next()

	}
}
