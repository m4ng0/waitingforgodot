package main

import "github.com/spf13/viper"

type applicationConfig struct {
	Rfid struct {
		Command struct {
			AccessGranted string
		}
		Sound struct {
			AccessGranted string
			AccessDenied  string
		}
		Access []struct {
			Key string
		}
	}
}

var Config applicationConfig

func LoadConfig() error {
	v := viper.New()
	v.SetConfigName("rfid")
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
