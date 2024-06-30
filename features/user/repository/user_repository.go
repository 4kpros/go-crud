package repository

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/features/user/model"
)

type UserRepository interface {
	Create(user *model.User) error
	UpdateUser(user *model.User) error
	UpdateUserInfo(userInfo *model.UserInfo) error
	Delete(id string) (int64, error)
	FindById(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByPhoneNumber(phoneNumber int) (*model.User, error)
	FindByProvider(provider string, providerUserId string) (*model.User, error)
	FindAll(filter *types.Filter, pagination *types.Pagination) ([]model.User, error)
}
