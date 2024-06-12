package utils

import "github.com/spf13/viper"

func Init(path string) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	return viper.ReadInConfig()
}
