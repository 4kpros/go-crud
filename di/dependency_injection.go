package di

import (
	"github.com/4kpros/go-api/config"
	authController "github.com/4kpros/go-api/features/auth/controller"
	authRepository "github.com/4kpros/go-api/features/auth/repository"
	authRouter "github.com/4kpros/go-api/features/auth/router"
	authService "github.com/4kpros/go-api/features/auth/service"
	userController "github.com/4kpros/go-api/features/user/controller"
	userRepository "github.com/4kpros/go-api/features/user/repository"
	userRouter "github.com/4kpros/go-api/features/user/router"
	userService "github.com/4kpros/go-api/features/user/service"
	"github.com/gin-gonic/gin"
)

func InitRepositories() (
	authRepo *authRepository.AuthRepository,
	userRepo *userRepository.UserRepository,
) {
	tmpAuth := authRepository.NewAuthRepositoryImpl(config.DB)
	tmpUser := userRepository.NewUserRepositoryImpl(config.DB)
	authRepo = &tmpAuth
	userRepo = &tmpUser
	return
}

func InitServices(
	authRepo *authRepository.AuthRepository,
	userRepo *userRepository.UserRepository,
) (
	authSer *authService.AuthService,
	userSer *userService.UserService,
) {
	tmpAuth := authService.NewAuthServiceImpl(*authRepo)
	tmpUser := userService.NewUserServiceImpl(*userRepo)
	authSer = &tmpAuth
	userSer = &tmpUser
	return
}

func InitControllers(
	authSer *authService.AuthService,
	userSer *userService.UserService,
) (
	authContr *authController.AuthController,
	userContr *userController.UserController,
) {
	tmpAuth := *authController.NewAuthController(*authSer)
	tmpUser := *userController.NewUserController(*userSer)
	authContr = &tmpAuth
	userContr = &tmpUser
	return
}

func InitRouters(
	routerGroup *gin.RouterGroup,
	authContr *authController.AuthController,
	userContr *userController.UserController,
) {
	authRouter.SetupAuthRouter(routerGroup, authContr)
	userRouter.SetupUserRouter(routerGroup, userContr)
}
