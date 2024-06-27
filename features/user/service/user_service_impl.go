package service

import (
	"fmt"
	"net/http"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/features/user/model"
	"github.com/4kpros/go-api/features/user/repository"
)

type UserServiceImpl struct {
	Repository repository.UserRepository
}

// Create implements UserService.
func (service *UserServiceImpl) Create(user *model.User) (errCode int, err error) {
	err = service.Repository.Create(user)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	return
}

// Delete implements UserService.
func (service *UserServiceImpl) Delete(id string) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id)
	if err != nil {
		errCode = http.StatusInternalServerError
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		message := "Could not delete user that doesn't exists! Please enter valid id."
		err = fmt.Errorf("%s", message)
		return
	}
	return
}

// FindAll implements UserService.
func (service *UserServiceImpl) FindAll(filter *types.Filter, pagination *types.Pagination) (users []model.User, errCode int, err error) {
	users, err = service.Repository.FindAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
	}
	return
}

// FindById implements UserService.
func (service *UserServiceImpl) FindById(id string) (user *model.User, errCode int, err error) {
	user, err = service.Repository.FindById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	if user == nil {
		errCode = http.StatusNotFound
		message := "User not found! Please enter valid id."
		err = fmt.Errorf("%s", message)
	}
	return
}

// Update implements UserService.
func (service *UserServiceImpl) Update(user *model.User) (errCode int, err error) {
	err = service.Repository.Update(user)
	return
}

func NewUserServiceImpl(repository repository.UserRepository) UserService {
	return &UserServiceImpl{Repository: repository}
}
