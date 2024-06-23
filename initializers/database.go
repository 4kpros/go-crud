package initializers

import (
	"fmt"
	"log"
	"os"

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
	DB, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
		}),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	fmt.Println("Connected to database: ", DB.Name())
}
