package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	TWIT_PORT string

	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	MONGO_DB_HOST string
	MONGO_DB_PORT int
	MONGO_DB_NAME string
	MONGO_DB_USER string
	MONGO_DB_PASS string

	KAFKA_HOST string
	KAFKA_PORT string

	REDIS_HOST string
	REDIS_PORT string

	EMAIL_NAME string
	EMAIL_PASSWORD string

	SMS_TOKEN string

	BOT_TOKEN string
	GROUP_ID string
	GROUP_FOR_INFO_ID string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.TWIT_PORT = cast.ToString(coalesce("TWIT_PORT", ":8097"))

	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "postgres-twitter"))
	config.DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "1111"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "twitter"))

	config.KAFKA_HOST = cast.ToString(coalesce("KAFKA_HOST", "localhost"))
	config.KAFKA_PORT = cast.ToString(coalesce("KAFKA_PORT", ":9092"))

	config.REDIS_HOST = cast.ToString(coalesce("REDIS_HOST", "redis-book"))
	config.REDIS_PORT = cast.ToString(coalesce("REDIS_PORT", ":6379"))

	config.EMAIL_NAME = cast.ToString(coalesce("EMAIL_NAME", "string"))
	config.EMAIL_PASSWORD = cast.ToString(coalesce("EMAIL_PASSWORD", "string"))

	config.SMS_TOKEN = cast.ToString(coalesce("SMS_TOKEN", "string"))

	config.BOT_TOKEN = cast.ToString(coalesce("BOT_TOKEN", "7867941196:AAHrsNk-gdDRP9jHmma9lRtRLwFDrvCGZZM"))
	config.GROUP_ID = cast.ToString(coalesce("GROUP_ID", "-4593428130"))
	config.GROUP_FOR_INFO_ID = cast.ToString(coalesce("GROUP_FOR_INFO_ID", "-4501154940"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
