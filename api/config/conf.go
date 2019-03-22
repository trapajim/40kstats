package config

import (
	"fmt"
	"log"
	"os"

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
	env := os.Getenv("ENV")

	if env == "" {
		log.Fatal("$ENV must be set")
	}

	return &Config{
		DB: &DBConfig{
			Host:     viper.GetString(env + ".host"),
			Port:     viper.GetString(env + ".port"),
			Dialect:  "postgres",
			Username: viper.GetString(env + ".user"),
			Password: viper.GetString(env + ".pass"),
			Name:     viper.GetString(env + ".dbname"),
			Charset:  "utf8",
		},
		OAuth: &OAuth{
			Id:       viper.GetString(env + ".oauthId"),
			Domain:   viper.GetString(env + ".oauthDomain"),
			Callback: viper.GetString(env + ".oauthCallback"),
			Secret:   viper.GetString(env + ".oauthSecret"),
			Audience: viper.GetString(env + ".oauthSecret"),
			ClientId: viper.GetString(env + ".oauthClientId"),
		},
	}
}
