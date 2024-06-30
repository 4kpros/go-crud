package main

import (
	"fmt"

	"github.com/4kpros/go-api/common/helpers"
	"github.com/4kpros/go-api/common/middleware"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/di"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	// Setup logger
	helpers.SetupLogger()

	// Load env variables
	errAppEnv := config.LoadAppEnv(".")
	if errAppEnv != nil {
		helpers.Logger.Warn(
			"Failed to load app ENV vars!",
			zap.String("Error", errAppEnv.Error()),
		)
	} else {
		helpers.Logger.Warn(
			"App ENV variables loaded!",
		)
	}

	// Setup argon2id params for crypto
	_, errArgonCryptoParamsUtils := utils.EncryptWithArgon2id("")
	if errArgonCryptoParamsUtils != nil {
		helpers.Logger.Warn(
			"Failed to setup argon2id params!",
			zap.String("Error", errArgonCryptoParamsUtils.Error()),
		)
	} else {
		helpers.Logger.Warn(
			"Argon2id crypto set ok!",
		)
	}

	// Connect to postgres database
	errPostgresDB := config.ConnectToPostgresDB()
	if errPostgresDB != nil {
		helpers.Logger.Warn(
			"Failed to connect to Postgres database!",
			zap.String("Error", errPostgresDB.Error()),
		)
	} else {
		helpers.Logger.Info(
			"Connected to Postgres database!",
		)
	}

	// Connect to memcache
	errMemcache := config.ConnectToMemcache()
	if errMemcache != nil {
		helpers.Logger.Warn(
			"Failed to connect to Memcache!",
			zap.String("Error", errMemcache.Error()),
		)
	} else {
		helpers.Logger.Info(
			"Connected to Memcache!",
		)
	}

	// Connect to redis
	errRedis := config.ConnectToRedis()
	if errRedis != nil {
		helpers.Logger.Warn(
			"Failed to connect to Redis!",
			zap.String("Error", errRedis.Error()),
		)
	} else {
		helpers.Logger.Info(
			"Connected to Redis!",
		)
	}

	// Load pem
	errPem := config.LoadPem()
	if errPem != nil {
		helpers.Logger.Warn(
			"Failed to load all pem files!",
			zap.String("Error", errRedis.Error()),
		)
	} else {
		helpers.Logger.Info(
			"All pem files loaded!",
		)
	}
}

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
func main() {
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
