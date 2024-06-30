package repository

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/features/user/model"
)

type AuthRepository interface {
	Create(user *model.User) error
	CreateUserInfo(userInfo *model.UserInfo) error
	Update(user *model.User) error
	UpdatePasswordById(id string, password string) (*model.User, error)
	Delete(id string) (int64, error)
	FindById(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByPhoneNumber(phoneNumber int) (*model.User, error)
	FindByProvider(provider string, providerUserId string) (*model.User, error)
	FindAll(filter *types.Filter, pagination *types.Pagination) ([]model.User, error)
}
