package auth

import (
	"github.com/4kpros/go-crud/services/auth/routes"
	"github.com/gin-gonic/gin"
)

func SetupService(r *gin.Engine) {
	routes.SetupSignInRoutes(r)
	routes.SetupSignUpRoutes(r)
	routes.SetupResetPasswordRoutes(r)
}
