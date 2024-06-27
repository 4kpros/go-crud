package main

import (
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/features/auth"
	"github.com/4kpros/go-api/features/user"
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
	auth.SetupMigrations() // Auth migrations
	user.SetupMigrations() // User migrations
}
