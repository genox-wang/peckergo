package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("./api/config/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err.Error())
	}
}

// GetString get string config by key
func GetString(key string) string {
	return viper.GetString(key)
}

// GetBool get bool config by key
func GetBool(key string) bool {
	return viper.GetBool(key)
}

// GetInt 配置中获取 int 类型
func GetInt(key string) int {
	return viper.GetInt(key)
}
