package config

import "github.com/spf13/viper"

// struct obj for env setting
type Configuration struct {
	// DataBase Setup
	DBHost string `mapstructure:"DB_URI"`
	DBName string `mapstructure:"DB_NAME"`
	// Redis Setup
	RedisUrl string `mapstructure:"REDIS_URL"`
	// jwt
	Secret string `mapstructure:"SECRET"`
	// Server Port
	ServerUrl  string `mapstructure:"SERVER_PORT"`
	ServerPort string `mapstructure:"SERVER_PORT"`
}

// Loads Env file content
func LoadConfig(path string) (config Configuration, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	// handle null
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
