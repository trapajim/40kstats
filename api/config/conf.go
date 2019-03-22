package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DB    *DBConfig
	OAuth *OAuth
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	Charset  string
}
type OAuth struct {
	Id       string
	Domain   string
	Callback string
	Secret   string
	Audience string
	ClientId string
}

func GetConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	pre := "development"
	return &Config{
		DB: &DBConfig{
			Host:     viper.GetString(pre + ".host"),
			Port:     viper.GetString(pre + ".port"),
			Dialect:  "postgres",
			Username: viper.GetString(pre + ".user"),
			Password: viper.GetString(pre + ".pass"),
			Name:     viper.GetString(pre + ".dbname"),
			Charset:  "utf8",
		},
		OAuth: &OAuth{
			Id:       viper.GetString(pre + ".oauthId"),
			Domain:   viper.GetString(pre + ".oauthDomain"),
			Callback: viper.GetString(pre + ".oauthCallback"),
			Secret:   viper.GetString(pre + ".oauthSecret"),
			Audience: viper.GetString(pre + ".oauthSecret"),
			ClientId: viper.GetString(pre + ".oauthClientId"),
		},
	}
}
