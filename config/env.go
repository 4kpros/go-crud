package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	ServerPort string `mapstructure:"PORT"`

	ApiKey string `mapstructure:"API_KEY"`

	PostGresHost     string `mapstructure:"POSTGRES_HOST"`
	PostGresPort     string `mapstructure:"POSTGRES_PORT"`
	PostGresUserName string `mapstructure:"POSTGRES_USERNAME"`
	PostGresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostGresDatabase string `mapstructure:"POSTGRES_DATABASE"`
	PostGresSslMode  string `mapstructure:"POSTGRES_SSL_MODE"`
	PostGresTimeZone string `mapstructure:"POSTGRES_TIME_ZONE"`

	JwtTokenSecret    string `mapstructure:"JWT_TOKEN_SECRET_"`
	JwtTokenMaxAge    string `mapstructure:"JWT_TOKEN_MAX_AGE_"`
	JwtTokenExpiredIn string `mapstructure:"JWT_TOKEN_EXPIRED_IN"`
}
type CryptoConfig struct {
	ArgonMemoryLeft  int `mapstructure:"ARGON_PARAM_MEMORY_L"`
	ArgonMemoryRight int `mapstructure:"ARGON_PARAM_MEMORY_R"`
	ArgonIterations  int `mapstructure:"ARGON_PARAM_ITERATIONS"`
	ArgonSaltLength  int `mapstructure:"ARGON_PARAM_SALT_LENGTH"`
	ArgonKeyLength   int `mapstructure:"ARGON_PARAM_KEY_LENGTH"`
}

var AppEnvConfig AppConfig
var CryptoEnvConfig CryptoConfig

func LoadAppEnvConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err == nil {
		err = viper.Unmarshal(&AppEnvConfig)
	}
	return
}

func LoadCryptoEnvConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("crypto")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err == nil {
		err = viper.Unmarshal(&CryptoEnvConfig)
	}
	return
}
