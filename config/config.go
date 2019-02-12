package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func setupDefaults() {
	viper.SetDefault("baseurl", "msg-na.ecouser.net")
	viper.SetDefault("deviceId", "suceanne-1.0")
	viper.SetDefault("countryCode", "us")
	viper.SetDefault("continentCode", "na")
}

func InitializeConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("etc/")
	viper.SetConfigType("yaml")

	setupDefaults()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error while reading config file: %s", err))
	}
}