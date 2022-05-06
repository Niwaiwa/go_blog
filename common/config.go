package common

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env       string
	GinMode   string
	Port      string
	SecretKey string
	DB        *DBConfig
	RedisDB   *RedisConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	Charset  string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

var configSet *Config

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	configSet = &Config{
		Env:       os.Getenv("ENV"),
		GinMode:   os.Getenv("GIN_MODE"),
		Port:      os.Getenv("PORT"),
		SecretKey: os.Getenv("SECRET_KEY"),
		DB: &DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Charset:  "utf8mb4",
		},
		RedisDB: &RedisConfig{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
	}
}

func GetConfig() *Config {
	return configSet
}
