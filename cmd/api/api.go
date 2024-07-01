package api

import (
	"fmt"

	"github.com/4kpros/go-api/common/middleware"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/di"

	"github.com/gin-gonic/gin"
)

// @title API Documentation
// @version 1.0
// @description This is the API documentation

// @contact.name Prosper Abouar
// @contact.email prosper.abouar@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey X-API-Key
// @in header
// @name ApiKey
// @description Enter the API key to have access

// @securityDefinitions.apikey Bearer
// @in header
// @name Bearer
// @description Enter Bearer with space and your token
func Start() {
	// Setup gin for your API
	gin.SetMode(config.AppEnv.GinMode)
	gin.ForceConsoleColor()
	engine := gin.Default()
	engine.HandleMethodNotAllowed = true
	engine.ForwardedByClientIP = true
	engine.SetTrustedProxies([]string{"127.0.0.1"})
	engine.Use(middleware.ErrorsHandler())
	apiGroup := engine.Group(config.AppEnv.ApiGroup)

	// Inject Dependencies
	authRepo, userRepo :=
		di.InitRepositories() // Repositories
	authSer, userSer :=
		di.InitServices(
			authRepo, userRepo,
		) // Services
	authContr, userContr :=
		di.InitControllers(
			authSer, userSer,
		) // Controllers
	di.InitRouters(
		apiGroup, authContr, userContr,
	) // Routers

	// Run gin
	formattedPort := fmt.Sprintf(":%d", config.AppEnv.ApiPort)
	engine.Run(formattedPort)
}
