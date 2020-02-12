package configuration

import (
	"github.com/spf13/viper"
)

type Constants struct {
	DatabaseConfig Database
}

type Database struct {
	DBName string
	ConnectionString string
}


func InitViper() (Constants, error){
	viper.SetConfigName("config")
	viper.AddConfigPath("./configuration/")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig();err != nil {
		return Constants{}, err
	}

	var constants Constants
	err := viper.Unmarshal(&constants)
	return constants, err
}