package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GOOGLE_CLIENT_ID           string
	GOOGLE_CLIENT_SECRET       string
	GOOGLE_CLIENT_CALLBACK_URL string

	GEMINI_LLM_KEY string

	PORT string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error loading from .env")
	}
	AppConfig = Config{
		GOOGLE_CLIENT_ID:           os.Getenv("GOOGLE_CLIENT_ID"),
		GOOGLE_CLIENT_SECRET:       os.Getenv("GOOGLE_CLIENT_SECRET"),
		GOOGLE_CLIENT_CALLBACK_URL: os.Getenv("GOOGLE_CLIENT_CALLBACK_URL"),
		PORT:                       os.Getenv("PORT"),
	}
	log.Println("Configuration loaded.")
}
