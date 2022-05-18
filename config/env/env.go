package env

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	development string = "development"
	production  string = "production"
	test        string = "test"
)

func getEnvFile(modeENV string) string {
	var fileName string
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if modeENV == test {
		fileName = ".env.test"
	}
	fileName = ".env"

	res := fmt.Sprintf("%s/%s", rootDir, fileName)
	log.Printf("Path to enviroment file: %s", res)

	return res
}

/**
* Функция для получение env значений из .env файла
 */
func ConfigInit() {
	modeENV := getEnv("MODE_ENV", development)
	log.Printf("The application is running in the %s mode", modeENV)

	fileName := getEnvFile(modeENV)

	if err := godotenv.Load(fileName); err != nil {
		log.Println("No .env file found")
	}
}

type DatabaseConfig struct {
	User         string
	Password     string
	DatabaseName string
	Port         int
	Host         string
}

type Config struct {
	Database DatabaseConfig
}

func New() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:         getEnv("DATABASE_HOST", "localhost"),
			Port:         getEnvAsInt("DATABASE_PORT", 5432),
			User:         getEnv("DATABASE_USER", "postgres"),
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
