package main

import "github.com/spf13/viper"

type applicationConfig struct {
	Mqtt struct {
		ConnectionString string
		ClientName       string
		Topic            string
		DefaultSender    string
	}
}

var Config applicationConfig

func LoadConfig() error {
	v := viper.New()
	v.SetConfigName("aperito")
	v.AddConfigPath("/etc/godot")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(&Config); err != nil {
		return err
	}
	return nil
}
