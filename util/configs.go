package util

import (
	"fmt"

	"github.com/spf13/viper"
)

// Contains all configs for our app usinf viper
type Config struct {
	DbDriver      string `mapstructure:"DB_DRIVER"`
	DbSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// Reads configuration form file or env variable
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil{
		fmt.Printf("Error in ReadInConfig : %v\n", err.Error())
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil{
		fmt.Printf("Error in ReadInConfig : %v\n", err.Error())
		return
	}

	return
}