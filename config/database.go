package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToPostgresDB() (err error) {
	dsn := "host=" + EnvConfig.PostGresHost + " user=" + EnvConfig.PostGresUserName + " password=" + EnvConfig.PostGresPassword + " dbname=" + EnvConfig.PostGresDatabase + " port=" + EnvConfig.PostGresPort + " sslmode=" + EnvConfig.PostGresSslMode + " TimeZone=" + EnvConfig.PostGresTimeZone
	DB, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN: dsn,
		}),
		&gorm.Config{},
	)
	return
}
