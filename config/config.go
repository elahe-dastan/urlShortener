package config

import (
	"log"

	"github.com/spf13/viper"
)

type Constants struct {
	DatabaseConfig Database
	Log            LogFile
}

type Database struct {
	DBName           string
	ConnectionString string
}

type LogFile struct {
	Address string
}

func ReadConfig() Constants {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("err: %s", err)
	}

	var cfg Constants
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("err: %s", err)
	}

	return cfg
}
