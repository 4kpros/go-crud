package service

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/features/user/model"
)

type UserService interface {
	Create(user *model.User) (errCode int, err error)
	Update(user *model.User) (errCode int, err error)
	Delete(id string) (affectedRows int64, errCode int, err error)
	FindById(id string) (user *model.User, errCode int, err error)
	FindAll(filter *types.Filter, pagination *types.Pagination) (users []model.User, errCode int, err error)
}
