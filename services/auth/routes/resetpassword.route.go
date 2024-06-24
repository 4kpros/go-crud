package routes

import (
	"github.com/4kpros/go-crud/common/router"
	"github.com/4kpros/go-crud/services/auth/controllers"
	"github.com/gin-gonic/gin"
)

func SetupResetPasswordRoutes(r *gin.Engine) {
	const path = "/auth"
	const requireAuth = false

	router.POST(r, path+"/reset-password/init", controllers.ResetPasswordInit, requireAuth)
	router.POST(r, path+"/reset-password/code", controllers.ResetPasswordCode, requireAuth)
	router.POST(r, path+"/reset-password/secret-question", controllers.ResetPasswordSecretQuestion, requireAuth)
	router.POST(r, path+"/reset-password/new-password", controllers.ResetPasswordNewPassword, requireAuth)
}
