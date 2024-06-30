package service

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/features/user/model"
)

type UserService interface {
	CreateWithEmail(user *model.User) (password string, errCode int, err error)
	CreateWithPhoneNumber(user *model.User) (password string, errCode int, err error)
	UpdateUser(user *model.User) (errCode int, err error)
	UpdateUserInfo(userInfo *model.UserInfo) (errCode int, err error)
	Delete(id string) (affectedRows int64, errCode int, err error)
	FindById(id string) (user *model.User, errCode int, err error)
	FindAll(filter *types.Filter, pagination *types.Pagination) (users []model.User, errCode int, err error)
}
