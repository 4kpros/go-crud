package routes

import (
	"github.com/4kpros/go-crud/common/router"
	"github.com/4kpros/go-crud/services/auth/controllers"
	"github.com/gin-gonic/gin"
)

func SetupSignInRoutes(r *gin.Engine) {
	const path = "/auth"
	const requireAuth = false

	router.POST(r, path+"/signin-email", controllers.SignInWithEmail, requireAuth)
	router.POST(r, path+"/signin-phone", controllers.SignInWithPhoneNumber, requireAuth)
	router.POST(r, path+"/signin-google", controllers.SignInWithGoogle, requireAuth)
	router.POST(r, path+"/signin-facebook", controllers.SignInWithFacebook, requireAuth)
	router.POST(r, path+"/signin-2fa", controllers.SignInWith2fa, requireAuth)
}
