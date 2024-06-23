package initializers

import (
	"os"

	"github.com/4kpros/go-crud/utils"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var host = os.Getenv("DB_HOST")
	var port = os.Getenv("DB_PORT")
	var username = os.Getenv("DB_USERNAME")
	var userPassword = os.Getenv("DB_USER_PASSWORD")
	var dbName = os.Getenv("DB_NAME")
	var sslMode = os.Getenv("DB_SSL_MODE")
	var timeZone = os.Getenv("DB_TIME_ZONE")

	dsn := "host=" + host + " user=" + username + " password=" + userPassword + " dbname=" + dbName + " port=" + port + " sslmode=" + sslMode + " TimeZone=" + timeZone
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
