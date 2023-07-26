package viper

import (
	"cfd/myapp/config"
	"github.com/spf13/viper"
)

func LoadConfig() (config.Config, error) {
	vp := viper.New()
	vp.SetConfigType("json")
	vp.AddConfigPath("./")
	err := vp.ReadInConfig()
	if err != nil {
		return config.Config{}, err
	}

	var configuration config.Config
	err = vp.Unmarshal(&configuration)
	if err != nil {
		return config.Config{}, err
	}

	return configuration, nil
}
