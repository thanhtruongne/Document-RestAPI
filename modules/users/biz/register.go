package biz

import (
	"context"
	"log"

	"github.com/user/Practice_api/common"
	"github.com/user/Practice_api/modules/users/models"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, condition map[string]interface{}, modelInfo ...string) (*models.User, error)
	CreateUser(ctx context.Context, data *models.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerUserBiz struct {
	store  RegisterStorage
	hasher Hasher
}

func NewRegisterUserBiz(store RegisterStorage, hasher Hasher) *registerUserBiz {
	return &registerUserBiz{store: store, hasher: hasher}
}

func (biz *registerUserBiz) RegiserUserByBuniesness(ctx context.Context, data *models.UserCreate) error {
	user, _ := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if user != nil {
		log.Print("user: ", user)
		return models.ErrEmailExisted
	}

	salt := common.GenSalt(30) //  tạo salt
	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt //  lưu salt
	data.Role = 1
	if err := biz.store.CreateUser(ctx, data); err != nil {
		return common.ErrCanNotCreateEntity(models.EnityModel, err)
	}
	return nil
}
