package router

import (
	"github.com/4kpros/go-api/common/router"
	"github.com/4kpros/go-api/features/auth/controller"
	"github.com/gin-gonic/gin"
)

func SetupAuthEndpoints(
	routerGroup *gin.RouterGroup,
	controller *controller.AuthController,
) {
	group := routerGroup.Group("/auth")
	const requireAuth = false

	// Sign in
	router.POST(group, "/signin-email", controller.SignInWithEmail, requireAuth)
	router.POST(group, "/signin-phone", controller.SignInWithEmail, requireAuth)
	router.POST(group, "/signin-provider", controller.SignInWithProvider, requireAuth)

	// Sign up
	router.POST(group, "/signup-email", controller.SignUpWithEmail, requireAuth)
	router.POST(group, "/signup-phone", controller.SignUpWithPhoneNumber, requireAuth)

	// Activate account
	router.POST(group, "/activate", controller.ActivateAccount, requireAuth)

	// Reset password
	router.POST(group, "/reset/init-email", controller.ResetPasswordEmailInit, requireAuth)
	router.POST(group, "/reset/init-phone", controller.ResetPasswordPhoneNumberInit, requireAuth)
	router.POST(group, "/reset/code", controller.ResetPasswordCode, requireAuth)
	router.POST(group, "/reset/new-password", controller.ResetPasswordNewPassword, requireAuth)
}
