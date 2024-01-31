package config

import (
	"os"
)

type Config struct {
	MongoURI string
	DBName   string
	Port     string
}

func LoadConfig() Config {
	return Config{
		MongoURI: os.Getenv("MONGODB_URI"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("PORT"),
	}
}
