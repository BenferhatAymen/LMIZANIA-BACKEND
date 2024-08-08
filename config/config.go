package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	JWTSecret string
	MONGO_URI string
	DB_NAME   string
	REDIS_ADDR string
	REDIS_PASSWORD string
	REDIS_USERNAME string
	MAILER_EMAIL string
	MAILER_PASSWORD string

)

func LoadENV() {
	log.Println("Loading environment variables")
	godotenv.Load("./.env")
	JWTSecret = os.Getenv("JWTSecret")
	MONGO_URI = os.Getenv("MONGO_URI")
	DB_NAME = os.Getenv("DB_NAME")
	REDIS_ADDR = os.Getenv("REDIS_ADDR")
	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	REDIS_USERNAME = os.Getenv("REDIS_USERNAME")
	MAILER_EMAIL = os.Getenv("MAILER_EMAIL")
	MAILER_PASSWORD = os.Getenv("MAILER_PASSWORD")
	log.Println("Environment variables loaded successfully")
}
