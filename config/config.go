package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	AppPort            string
	BasePath           string
	ImageKitPrivateKey string
	ImageKitPublicKey  string
	ImageKitEndpoint   string
}

var AppConfig *Config

func InitConfig() *Config {
	LoadEnv()
	AppConfig = &Config{
		DBHost:             getEnv("DB_HOST", "localhost"),
		DBPort:             getEnv("DB_PORT", "5432"),
		DBUser:             getEnv("DB_USER", "postgres"),
		DBPassword:         getEnv("DB_PASSWORD", "password"),
		DBName:             getEnv("DB_NAME", "ngoclam"),
		AppPort:            getEnv("APP_PORT", "8080"),
		BasePath:           getEnv("BASE_PATH", "./uploads"),
		ImageKitPrivateKey: getEnv("IMAGEKIT_PRIVATE_KEY", ""),
		ImageKitPublicKey:  getEnv("IMAGEKIT_PUBLIC_KEY", ""),
		ImageKitEndpoint:   getEnv("IMAGEKIT_ENDPOINT_URL", ""),
	}

	return AppConfig
}

// LoadEnv loads environment variables from .env file
func LoadEnv() error {
	// Find .env file in the project root or current directory
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env file not found: %v", err)
		// Continue execution even if .env is not found
		// Variables might be set in the environment directly
	}
	return nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
