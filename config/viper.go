package config

import (
	"log"

	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	viper := viper.New()
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./../")
	viper.AddConfigPath("./../../")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("viper error configuration: %s", err)
	}
	return viper
}
