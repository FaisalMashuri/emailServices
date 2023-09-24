package env

import (
	"log"
	"os"

	"github.com/FaisalMashuri/emailServices/models"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

var Config models.AppConfig

func init() {
	err := loadConfig()
	if err != nil {
		log.Fatal(err, "config/env : loadConfig")
	}
}

func loadConfig() (err error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}
	if err != nil {
		log.Fatal(err, "config/env : load godotenv")
	}

	err = env.Parse(&Config)
	if err != nil {
		return err
	}

	err = env.Parse(&Config.DatabaseConfig)
	if err != nil {
		return err
	}

	err = env.Parse(&Config.RabbitMQConfig)
	if err != nil {
		return err
	}

	err = env.Parse(&Config.EmailConfig)
	if err != nil {
		return err
	}
	return err
}
