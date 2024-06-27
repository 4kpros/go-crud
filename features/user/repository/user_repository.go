package repository

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/features/user/model"
)

type UserRepository interface {
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id string) (int64, error)
	FindById(id string) (*model.User, error)
	FindAll(filter *types.Filter, pagination *types.Pagination) ([]model.User, error)
}
