package main

import (
	"fmt"

	"github.com/4kpros/go-api/common/middleware"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/di"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	// Setup logger
	utils.InitializeLogger()

	// Load env variables
	errAppEnv := config.LoadAppEnv(".")
	if errAppEnv != nil {
		utils.Logger.Warn(
			"Failed to load app ENV vars!",
			zap.String("Error", errAppEnv.Error()),
		)
		return
	}
	utils.Logger.Warn(
		"App ENV variables loaded!",
	)

	// Setup argon params for crypto
	_, errArgonCryptoParamsUtils := utils.EncryptWithArgon2id("")
	if errArgonCryptoParamsUtils != nil {
		utils.Logger.Warn(
			"Failed to setup argon2id params!",
			zap.String("Error", errArgonCryptoParamsUtils.Error()),
		)
		return
	}
	utils.Logger.Warn(
		"Argon2id crypto set ok!",
	)

	// Connect to postgres database
	errPostgresDB := config.ConnectToPostgresDB()
	if errPostgresDB != nil {
		utils.Logger.Warn(
			"Failed to connect to Postgres database!",
			zap.String("Error", errPostgresDB.Error()),
		)
		return
	}
	utils.Logger.Info(
		"Connected to Postgres!",
	)

	// Connect to memcache
	errMemcache := config.ConnectToMemcache()
	if errMemcache != nil {
		utils.Logger.Warn(
			"Failed to connect to Memcache!",
			zap.String("Error", errMemcache.Error()),
		)
		return
	}
	utils.Logger.Info(
		"Connected to Memcache!",
	)

	// Connect to redis
	errRedis := config.ConnectToRedis()
	if errRedis != nil {
		utils.Logger.Warn(
			"Failed to connect to Redis!",
			zap.String("Error", errRedis.Error()),
		)
		return
	}
	utils.Logger.Info(
		"Connected to Redis!",
	)

	// Load pem
	errPem := config.LoadPem()
	if errPem != nil {
		utils.Logger.Warn(
			"Failed to load pem files!",
			zap.String("Error", errRedis.Error()),
		)
		return
	}
	utils.Logger.Info(
		"All pem files loaded!",
	)
}

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
