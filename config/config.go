package config

import (
	"bytes"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseConfig Database `mapstructure:"dbconfig"`
	Log            LogFile  `mapstructure:"log"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	DBName   string `mapstructure:"dbname"`
	Password string `mapstructure:"password"`
	SSLmode  string `mapstructure:"sslmode"`
}

type LogFile struct {
	Address string `mapstructure:"address"`
}

func ReadConfig() Config {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")

	if err := viper.ReadConfig(bytes.NewBufferString(Default)); err != nil {
		log.Fatalf("err: %s", err)
	}

	viper.SetConfigName("config.example")
	if err := viper.MergeInConfig(); err != nil {
		log.Print("No config file found")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("err: %s", err)
	}

	return cfg
}

func (d Database) Cstring() string {
	return "host=" + d.Host + " port=" + d.Port + " user=" + d.User + " dbname=" + d.DBName + " password=" + d.Password + " sslmode=" + d.SSLmode
}
