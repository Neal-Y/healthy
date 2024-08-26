package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	GPTApiKey     string
	WebURL        string
	AwsDbUsername string
	AwsDbPassword string
	AwsDbHost     string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	AppConfig = Config{
		GPTApiKey:     viper.GetString("CHATGPT_API_KEY"),
		WebURL:        viper.GetString("WEB_URL"),
		AwsDbUsername: viper.GetString("AWS_DATABASE_USERNAME"),
		AwsDbPassword: viper.GetString("AWS_DATABASE_PASSWORD"),
		AwsDbHost:     viper.GetString("AWS_DATABASE_HOST"),
	}
}
