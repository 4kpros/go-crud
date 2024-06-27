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
	errAppEnv := config.LoadAppEnvConfig(".")
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
	errCryptoEnv := config.LoadCryptoEnvConfig(".")
	if errCryptoEnv != nil {
		utils.Logger.Warn(
			"Failed to load crypto ENV vars!",
			zap.String("Error", errCryptoEnv.Error()),
		)
		return
	}
	utils.Logger.Warn(
		"Crypto ENV variables loaded!",
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
}

func main() {
	// Setup gin for your API
	gin.SetMode(config.AppEnvConfig.GinMode)
	gin.ForceConsoleColor()
	engine := gin.Default()
	engine.Use(middleware.ErrorsHandler())
	// engine.ForwardedByClientIP = true
	// engine.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.2", "10.0.0.0/8"})

	apiGroup := engine.Group(config.AppEnvConfig.ApiGroup)

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
	formattedPort := fmt.Sprintf(":%d", config.AppEnvConfig.ServerPort)
	engine.Run(formattedPort)
}
