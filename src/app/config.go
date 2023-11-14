package app

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigFile(`.env`)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error load .env")
	}
	return nil
}
