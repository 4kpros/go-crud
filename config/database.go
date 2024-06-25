package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToPostgresDB() (err error) {
	dsn := "host=" + AppEnvConfig.PostGresHost + " user=" + AppEnvConfig.PostGresUserName + " password=" + AppEnvConfig.PostGresPassword + " dbname=" + AppEnvConfig.PostGresDatabase + " port=" + AppEnvConfig.PostGresPort + " sslmode=" + AppEnvConfig.PostGresSslMode + " TimeZone=" + AppEnvConfig.PostGresTimeZone
	DB, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN: dsn,
		}),
		&gorm.Config{},
	)
	return
}
