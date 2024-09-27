package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	API_PORT string

	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	TWIT_HOST string
	TWIT_PORT string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.API_PORT = cast.ToString(coalesce("AUTH_PORT", ":8097"))

	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "postgres-twitter"))
	config.DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "1111"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "twitter"))

	config.TWIT_HOST = cast.ToString(coalesce("TWIT_HOST", "twit_service"))
	config.TWIT_PORT = cast.ToString(coalesce("TWIT_PORT", ":50069"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
