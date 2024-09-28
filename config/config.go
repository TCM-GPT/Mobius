package config

import (
	"gen-SFT-Dataset/internal/model"
	"github.com/spf13/viper"
	"log"
)

var config *model.App

func GetInstance() *model.App {
	return config
}

func InitConfigFile() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AddConfigPath("./config")
	if err := v.ReadInConfig(); err != nil {
		log.Println("config file not found, please check config file path")
		panic(err)
	}

	err := v.Unmarshal(&config)
	if err != nil {
		log.Println("viper unmarshal err")
		return
	}
}
