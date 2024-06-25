package config

import (
	"github.com/4kpros/go-crud/common/utils"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := "host=" + EnvConfig.PostGresHost + " user=" + EnvConfig.PostGresUserName + " password=" + EnvConfig.PostGresPassword + " dbname=" + EnvConfig.PostGresDatabase + " port=" + EnvConfig.PostGresPort + " sslmode=" + EnvConfig.PostGresSslMode + " TimeZone=" + EnvConfig.PostGresTimeZone
	var err error
	DB, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN: dsn,
		}),
		&gorm.Config{},
	)

	if err != nil {
		utils.Logger.Warn(
			"Failed to connect to database !",
			zap.String("Error", err.Error()),
		)
		return
	}
	utils.Logger.Info(
		"Connected to database: ",
		zap.String("DB name", DB.Name()),
	)
}
