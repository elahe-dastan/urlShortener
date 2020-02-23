package config

import (
	"github.com/spf13/viper"
	"log"
)

type Constants struct {
	DatabaseConfig Database
}

type Database struct {
	DBName string
	ConnectionString string
}


func ReadConfig() Constants {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig();err != nil {
		log.Fatalf("err: %s", err)
	}

	var constants Constants
	if err := viper.Unmarshal(&constants); err != nil {
		log.Fatalf("err: %s", err)
	}

	return constants
}