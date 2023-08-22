package config

import "github.com/spf13/viper"

type Configuration struct {
	// DataBase Setup
	DBHost string `mapstructure:"DB_URI"`
	DBName string `mapstructure:"DB_NAME"`
	// Redis Setup
	RedisUrl string `mapstructure:"REDIS_URL"`
	// jwt secret
	Secret string `mapstructure:"SECRET"`
	// Server Port
	Port string `mapstructure:"SERVER_PORT"`
}

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
