package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig() {
	// read the config.json file
	viper.SetConfigName("config")
	viper.AddConfigPath("env")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	log.Println("Config file loaded:", viper.ConfigFileUsed())
}

func GetString(key string) string {
	return viper.GetString(key)
}
