package biz

import (
	"context"

	"github.com/user/Practice_api/common"
	"github.com/user/Practice_api/components/tokenProviders"
	"github.com/user/Practice_api/modules/users/models"
)

type LoginStorage interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfos ...string) (*models.User, error)
}

type LoginBusiness struct {
	StoreUser     LoginStorage
	TokenProvider tokenProviders.Provider
	Hasher        Hasher
	Exipry        int
}

func LoginStoreInstance(login LoginStorage, tokenProviders tokenProviders.Provider, hasher Hasher, exipry int) *LoginBusiness {
	return &LoginBusiness{StoreUser: login, TokenProvider: tokenProviders, Hasher: hasher, Exipry: exipry}
}

func (biz *LoginBusiness) LoginBusiness(ctx context.Context, data *models.UserLogin) (tokenProviders.Token, error) {
	user, err := biz.StoreUser.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, models.ErrEmailOrPasswordInvalid
	}

	passHashed := biz.Hasher.Hash(data.Password + user.Salt)

	if passHashed != user.Password {
		return nil, models.ErrEmailOrPasswordInvalid
	}

	payload := &common.TokenPayload{
		UId:   user.Id,
		URole: user.Role.String(),
	}

	accecssToken, err := biz.TokenProvider.Generate(payload, biz.Exipry)
	if err != nil {
		return nil, common.ErrorInvalidRequest(err)
	}

	return accecssToken, nil
}
