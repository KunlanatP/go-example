package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	DB_URL string `env:"DB_URL,required"`
}

var Default = Config{}

func init() {
	log.Println("Load init config ...")
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file : ", err)
	}
	if err := env.Parse(&Default); err != nil {
		log.Fatal("Error Parsing config .env file : ", err)
	}
}
