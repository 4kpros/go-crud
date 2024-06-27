package controller

import (
	"github.com/4kpros/go-api/features/user/model"
	"github.com/4kpros/go-api/features/user/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{Service: service}
}

func (controller *UserController) Create(c *gin.Context) {
	user := &model.User{}
	controller.Service.Create(user)
}

func (controller *UserController) Update(c *gin.Context) {

}

func (controller *UserController) Delete(c *gin.Context) {

}

func (controller *UserController) FindById(c *gin.Context) {

}

func (controller *UserController) FindAll(c *gin.Context) {

}
