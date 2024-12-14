package jwt

import (
	"fmt"
	"time"

	"github.com/user/Practice_api/common"
	"github.com/user/Practice_api/components/tokenProviders"

	"github.com/dgrijalva/jwt-go"
)

type jwtProvider struct {
	secrect string
	prefix  string
}

func NewTokenJWTProvider(prefix, secret string) *jwtProvider {
	return &jwtProvider{prefix: prefix, secrect: secret}
}

type myClaims struct {
	Payload common.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

func (t *Token) GetToken() string {
	return t.Token
}

// create token
func (j *jwtProvider) Generate(data tokenProviders.TokenPayload, expiry int) (tokenProviders.Token, error) {
	now := time.Now()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		common.TokenPayload{
			UId:   data.UserId(),
			URole: data.Role(),
		},
		jwt.StandardClaims{
			ExpiresAt: now.Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  now.Local().Unix(),
			Id:        fmt.Sprintf("%d", now.UnixNano()),
		},
	})

	myToken, err := t.SignedString([]byte(j.secrect))
	if err != nil {
		return nil, err
	}
	token := &Token{
		Token:   myToken,
		Created: now,
		Expiry:  expiry,
	}
	return token, nil
}

//check token

func (j *jwtProvider) Validate(myToken string) (tokenProviders.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secrect), nil
	})

	if err != nil {
		return nil, tokenProviders.ErrInvalidToken
	}

	if !res.Valid {
		return nil, tokenProviders.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)

	if !ok {
		return nil, tokenProviders.ErrInvalidToken
	}

	return claims.Payload, nil

}
func (j *jwtProvider) SecretKey() string {
	return j.secrect
}
