package utils

import "go.uber.org/zap"

var Logger *zap.Logger

func InitializeLogger() {
	Logger, _ = zap.NewProduction()
	defer Logger.Sync()
}

func PrintMigrationLogs(err error, modelName string) {
	if err != nil {
		Logger.Warn(
			"Failed to create << "+modelName+" >> table !",
			zap.String("Error", err.Error()),
		)
		return
	}
	Logger.Info(
		"Migration for << " + modelName + " >> Done !",
	)
}
