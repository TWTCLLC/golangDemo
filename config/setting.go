package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func ConfigInit() viper.Viper {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.SetDefault("application.port", 8080)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper := viper.GetViper()
	return *viper
}
