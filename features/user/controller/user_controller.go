package controller

import (
	"net/http"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/features/user/data/request"
	"github.com/4kpros/go-api/features/user/data/response"
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

// @Tags Users
// @Summary Create new user with email - [super-admin]
// @Accept  json
// @Produce  json
// @Param   email body string true "Enter your email"
// @Param   role body string true "Select role" Enums(super-admin, admin, manager, manager-assist, deliver, customer, customer-service)
// @Success 200 {object} response.CreateWithEmailResponse "OK"
// @Failure 400 {string} string "Invalid email or role!"
// @Failure 302 {string} string "User with this email already exists!"
// @Router /users/email [post]
func (controller *UserController) CreateWithEmail(c *gin.Context) {
	// Get data of req body
	var reqData = &request.CreateWithEmailRequest{}
	c.Bind(reqData)
	var user = &model.User{}
	user.Email = reqData.Email
	user.Role = reqData.Role

	// Execute the service
	password, errCode, err := controller.Service.CreateWithEmail(user)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.CreateWithEmailResponse{
			Email:    reqData.Email,
			Role:     reqData.Role,
			Password: password,
		},
	})
}

// @Tags Users
// @Summary Create new user with phone number - [super-admin]
// @Accept  json
// @Produce  json
// @Param   phoneNumber body int true "Enter your phone number"
// @Param   role body string true "Select role" Enums(super-admin, admin, manager, manager-assist, deliver, customer, customer-service)
// @Success 200 {object} response.CreateWithPhoneNumberResponse "OK"
// @Failure 400 {string} string "Invalid phone number or role!"
// @Failure 302 {string} string "User with this phone number already exists!"
// @Router /users/phone [post]
func (controller *UserController) CreateWithPhoneNumber(c *gin.Context) {
	// Get data of req body
	var reqData = &request.CreateWithPhoneNumberRequest{}
	c.Bind(reqData)
	var user = &model.User{}
	user.PhoneNumber = reqData.PhoneNumber
	user.Role = reqData.Role

	// Execute the service
	password, errCode, err := controller.Service.CreateWithPhoneNumber(user)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.CreateWithPhoneNumberResponse{
			PhoneNumber: reqData.PhoneNumber,
			Role:        reqData.Role,
			Password:    password,
		},
	})
}

// @Tags Users
// @Summary Update user
// @Accept  json
// @Produce  json
// @Param   email body string true "Enter your email"
// @Param   phoneNumber body int true "Enter your phone number"
// @Param   role body string true "Select role" Enums(super-admin, admin, manager, manager-assist, deliver, customer, customer-service)
// @Success 200 {object} model.User "OK"
// @Failure 400 {string} string "Invalid inputs!"
// @Failure 401 {string} string "Invalid user session!"
// @Failure 403 {string} string "Not permitted!"
// @Failure 404 {string} string "User not found!"
// @Router /users/{id} [put]
func (controller *UserController) UpdateUser(c *gin.Context) {
	// Get data of req body
	var user = &model.User{}
	c.Bind(user)

	// Execute the service
	errCode, err := controller.Service.UpdateUser(user)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: user,
	})
}

// @Tags Users
// @Summary Update user info
// @Accept  json
// @Produce  json
// @Param   payload body model.UserInfo false "User info model"
// @Success 200 {object} model.UserInfo "OK"
// @Failure 400 {string} string "Invalid inputs!"
// @Failure 401 {string} string "Invalid user session!"
// @Failure 403 {string} string "Not permitted!"
// @Failure 404 {string} string "User not found!"
// @Router /users/info/{id} [put]
func (controller *UserController) UpdateUserInfo(c *gin.Context) {
	// Get data of req body
	var userInfo = &model.UserInfo{}
	c.Bind(userInfo)

	// Execute the service
	errCode, err := controller.Service.UpdateUserInfo(userInfo)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: userInfo,
	})
}

// @Tags Users
// @Summary Update user info
// @Accept  json
// @Produce  json
// @Param   id path string true "User id"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid inputs!"
// @Failure 401 {string} string "Invalid user session!"
// @Failure 403 {string} string "Not permitted!"
// @Failure 404 {string} string "User not found!"
// @Router /users/{id} [delete]
func (controller *UserController) Delete(c *gin.Context) {
	// Get data of req header
	id := c.Param("id")

	// Execute the service
	user, errCode, err := controller.Service.Delete(id)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: user,
	})
}

// @Tags Users
// @Summary Update user info
// @Accept  json
// @Produce  json
// @Param   id path string true "User id"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid inputs!"
// @Failure 401 {string} string "Invalid user session!"
// @Failure 403 {string} string "Not permitted!"
// @Failure 404 {string} string "User not found!"
// @Router /users/{id} [get]
func (controller *UserController) FindById(c *gin.Context) {
	// Get data of req header
	id := c.Param("id")

	// Execute the service
	user, errCode, err := controller.Service.FindById(id)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: user,
	})
}

// @Tags Users
// @Summary Get all users
// @Accept  json
// @Produce  json
// @Param search query string false "Search keyword"
// @Param page query int false "Page"
// @Param limit query int false "Page limit"
// @Param orderBy query string false "Filter order by"
// @Param sort query string false "Sort asc, desc"
// @Success 200 {object} string "OK"
// @Failure 401 {string} string "Invalid user session!"
// @Failure 403 {string} string "Not permitted!"
// @Router /users/ [get]
func (controller *UserController) FindAll(c *gin.Context) {
	// Get data of req body
	pagination, filter := utils.GetPaginationFiltersFromQuery(c)

	// Execute the service
	users, errCode, err := controller.Service.FindAll(filter, pagination)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, types.WebSuccessPaginatedResponse{
		Data:       users,
		Filter:     filter,
		Pagination: pagination,
	})
}
