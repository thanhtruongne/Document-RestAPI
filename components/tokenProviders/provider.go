package tokenProviders

import (
	"errors"

	"github.com/user/Practice_api/common"
)

type Provider interface { // khai báo dạng dùng chung cho các generate
	Generate(data TokenPayload, expiry int) (Token, error)
	Validate(token string) (TokenPayload, error)
	SecretKey() string
}

type TokenPayload interface {
	UserId() int
	Role() string
}

type Token interface {
	GetToken() string
}

var (
	ErrNotFound     = common.NewErrorCustomResponse(errors.New("Token not found"), "Token not found", "ErrNotFound")
	ErrEndcodeToken = common.NewErrorCustomResponse(errors.New("Error encode token"), "Error encode token", "ErrEncodeToken")
	ErrInvalidToken = common.NewErrorCustomResponse(errors.New("Invalid token"), "Invalid token", "ErrInvalidToken")
)
