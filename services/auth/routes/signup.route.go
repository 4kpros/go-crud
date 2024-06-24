package routes

import (
	"github.com/4kpros/go-crud/common/router"
	"github.com/4kpros/go-crud/services/auth/controllers"
	"github.com/gin-gonic/gin"
)

func SetupSignUpRoutes(r *gin.Engine) {
	const path = "/auth"
	const requireAuth = false

	// Sign up
	router.POST(r, path+"/signup-email", controllers.SignUpWithEmail, requireAuth)
	router.POST(r, path+"/signup-phone", controllers.SignUpWithPhoneNumber, requireAuth)
	router.POST(r, path+"/signup-google", controllers.SignUpWithGoogle, requireAuth)
	router.POST(r, path+"/signup-facebook", controllers.SignUpWithFacebook, requireAuth)

	// Validate account
	router.POST(r, path+"/validate-account", controllers.ValidateAccount, requireAuth)
}
