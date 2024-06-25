package main

import (
	"github.com/4kpros/go-crud/common/middlewares"
	"github.com/4kpros/go-crud/common/utils"
	"github.com/4kpros/go-crud/config"
	"github.com/4kpros/go-crud/services/auth"
	"github.com/4kpros/go-crud/services/post"
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
		"Crypto ENV variables loaded!",
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
		"Connected to Postgres database: ",
		zap.String("DB name", config.DB.Name()),
	)
}

func main() {
	// Setup gin for HTTP requests
	r := gin.Default()
	r.Use(middlewares.ErrorsHandler())

	// Setup services
	auth.SetupService(r) // Auth service
	post.SetupService(r) // Post service

	// Run gin with custom port
	r.Run(":" + config.AppEnvConfig.ServerPort)
}
