package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		_ = fmt.Errorf("Fatal Error in config File %s", err.Error())
		panic(err)
	}
	return viper.GetString(key)
}
