package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string

	DbHost string
	DbPort int
	DbUser string
	DbPass string
	DbName string
}

func LoadConfig() *Config {

	// == Load .env file ==
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, using default environment variables")
	}

	env := os.Getenv("APP_ENV")
	envFile := ".env"
	if env == "uat" {
		envFile = ".env.uat"
	} else if env == "production" {
		envFile = ".env.prod"
	}

	// == Load Database Config from .env file ==
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("No %s file found, using default environment variables", envFile)
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		dbPort = 5432 // default port
	}

	return &Config{
		AppPort: Getenv("APP_PORT", "9090"),

		DbHost: Getenv("DB_HOST", "localhost"),
		DbPort: dbPort,
		DbUser: Getenv("DB_USER", "sa"),
		DbPass: Getenv("DB_PASSWORD", "GLGr0uplease"),
		DbName: Getenv("DB_NAME", "GLMOBILE_DB"),
	}
}

func Getenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func GetLocalPath() string {
	return Getenv("PATH_PHOTO_UAT", "")
}
