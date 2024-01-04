package config

import (
	"log"

	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	v := viper.New()

	v.SetConfigFile(".env")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("error while read file")
	}

	return v
}
