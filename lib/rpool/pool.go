package rpool

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("fail to read config, err: ", err)
	}
}
