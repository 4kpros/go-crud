package repository

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/features/auth/model"
	userModel "github.com/4kpros/go-api/features/user/model"
)

type AuthRepository interface {
	Create(user *model.NewUser) error
	CreateActivatedUser(user *userModel.User) error
	Update(user *model.NewUser) error
	UpdatePasswordById(id string, password string) (*model.NewUser, error)
	Delete(id string) (int64, error)
	FindById(id string) (*model.NewUser, error)
	FindByEmail(email string) (*model.NewUser, error)
	FindByPhoneNumber(phoneNumber int) (*model.NewUser, error)
	FindByProvider(provider string, providerUserId string) (*model.NewUser, error)
	FindAll(filter *types.Filter, pagination *types.Pagination) ([]model.NewUser, error)
}
