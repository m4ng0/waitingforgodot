package main

import "github.com/spf13/viper"

type applicationConfig struct {
	Button struct {
		Pin            string
		PollInterval   int
		DefaultState   int
		PressSoundFile string
		Command        string
	}
}

var Config applicationConfig

func LoadConfig() error {
	v := viper.New()
	v.SetConfigName("config")
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
