package config

import (
	"github.com/4kpros/go-crud/common/utils"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	ServerPort string `mapstructure:"PORT"`

	ApiKey string `mapstructure:"API_KEY"`

	PostGresHost     string `mapstructure:"POSTGRES_HOST"`
	PostGresPort     string `mapstructure:"POSTGRES_PORT"`
	PostGresUserName string `mapstructure:"POSTGRES_USERNAME"`
	PostGresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostGresDatabase string `mapstructure:"POSTGRES_DATABASE"`
	PostGresSslMode  string `mapstructure:"POSTGRES_SSL_MODE"`
	PostGresTimeZone string `mapstructure:"POSTGRES_TIME_ZONE"`

	JwtTokenKey       string `mapstructure:"JWT_TOKEN_SECRET_"`
	JwtTokenMaxAge    string `mapstructure:"JWT_TOKEN_MAX_AGE_"`
	JwtTokenExpiredIn string `mapstructure:"JWT_TOKEN_EXPIRED_IN"`
}

var EnvConfig Config

func LoadEnvironmentVariables(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		utils.Logger.Warn(
			"Failed to load ENV vars !",
			zap.String("Error", err.Error()),
		)
	}
	err = viper.Unmarshal(&EnvConfig)
	if err != nil {
		utils.Logger.Warn(
			"Failed to load ENV vars !",
			zap.String("Error", err.Error()),
		)
	}
}
