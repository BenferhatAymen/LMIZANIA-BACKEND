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
)

func LoadENV() {
	log.Println("Loading environment variables")
	godotenv.Load("./.env")
	JWTSecret = os.Getenv("JWTSecret")
	MONGO_URI = os.Getenv("MONGO_URI")
	DB_NAME = os.Getenv("DB_NAME")
}
