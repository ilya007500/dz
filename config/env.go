package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DataBaseConfig struct {
	url string
}

func NewDataBaseConfig() *DataBaseConfig {
	return &DataBaseConfig{
		url: os.Getenv("DATABASE_URL"),
	}
}

func Init() {
	if err := godotenv.Load(); err != nil{
		log.Println("No .env file")
		return
	}
	log.Println(".env file loaded")
}

