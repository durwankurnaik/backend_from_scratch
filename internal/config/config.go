package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func Config(key string) string {
	return viper.GetString(key)
}