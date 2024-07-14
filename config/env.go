package config

import (
	"github.com/spf13/viper"
)

type Env struct {
	// API config
	ApiPort      int    `mapstructure:"API_PORT"`
	ApiKey       string `mapstructure:"API_KEY"`
	ApiGroup     string `mapstructure:"API_GROUP"`
	GinMode      string `mapstructure:"GIN_MODE"`
	AllowedHosts string `mapstructure:"ALLOWED_HOSTS"`

	// Postgres database
	PostGresHost     string `mapstructure:"POSTGRES_HOST"`
	PostGresPort     int    `mapstructure:"POSTGRES_PORT"`
	PostGresUserName string `mapstructure:"POSTGRES_USERNAME"`
	PostGresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostGresDatabase string `mapstructure:"POSTGRES_DATABASE"`
	PostGresSslMode  string `mapstructure:"POSTGRES_SSL_MODE"`
	PostGresTimeZone string `mapstructure:"POSTGRES_TIME_ZONE"`

	// JWT
	JwtExpiresSignIn              int `mapstructure:"JWT_EXPIRES_SIGN_IN"`
	JwtExpiresSignInStayConnected int `mapstructure:"JWT_EXPIRES_SIGN_IN_STAY_CONNECTED"`
	JwtExpiresDefault             int `mapstructure:"JWT_EXPIRES_DEFAULT"`

	// Redis for fast key-value database
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     int    `mapstructure:"REDIS_PORT"`
	RedisUserName string `mapstructure:"REDIS_USERNAME"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisDatabase int    `mapstructure:"REDIS_DATABASE"`

	// Memcache for fast key-value database
	MemcacheServersCount int    `mapstructure:"MEMCACHE_SERVERS_COUNT"`
	MemcacheHostRange    string `mapstructure:"MEMCACHE_HOST_RANGE"`
	MemcacheInitialPort  int    `mapstructure:"MEMCACHE_INITIAL_PORT"`

	// Crypto Argon2id for passwords
	ArgonMemoryLeft  int `mapstructure:"ARGON_PARAM_MEMORY_L"`
	ArgonMemoryRight int `mapstructure:"ARGON_PARAM_MEMORY_R"`
	ArgonIterations  int `mapstructure:"ARGON_PARAM_ITERATIONS"`
	ArgonSaltLength  int `mapstructure:"ARGON_PARAM_SALT_LENGTH"`
	ArgonKeyLength   int `mapstructure:"ARGON_PARAM_KEY_LENGTH"`
}

var AppEnv = &Env{}

func LoadAppEnv(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err == nil {
		err = viper.Unmarshal(AppEnv)
	}
	return
}
