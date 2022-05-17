package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

/**
* Функция для получение env значений из .env файла
 */
func ConfigInit() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

type DatabaseConfig struct {
	Login        string
	Password     string
	DatabaseName string
}

type Config struct {
	Database DatabaseConfig
}

func New() *Config {
	return &Config{
		Database: DatabaseConfig{
			Login:        getEnv("DATABASE_USER", "postgres"),
			Password:     getEnv("DATABASE_PASSWORD", "qwerty123"),
			DatabaseName: getEnv("DATABASE_NAME", "test"),
		},
	}
}

func getEnv(key string, defaultValue string) string {

	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valStr); err == nil {
		return value
	}

	return defaultValue
}
