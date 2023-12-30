package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type values struct {
	BotToken    string `envconfig:"BOT_TOKEN" required:"true"`
	LoggerLevel string `envconfig:"LOGGER_LEVEL"`
	BotDebug    string `envconfig:"BOT_DEBUG"`
	YandexPath  string `envconfig:"YANDEX_PATH"`
}

var Values values

func LoadFromFile(fPath string) error {
	_ = godotenv.Load(fPath)

	err := envconfig.Process("", &Values)
	if err != nil {
		log.Println("ERROR: envconfig.Process(): ", err.Error())
	}
	return err
}
