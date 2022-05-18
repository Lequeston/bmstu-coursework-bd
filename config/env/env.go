package env

import (
	"log"
	"os"
	"path"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	DEVELOPMENT_MODE string = "development"
	PRODUCTION_MODE  string = "production"
	TEST_MODE        string = "test"
)

func getEnvFile(modeENV string) string {
	var fileName string
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fileName = ".env"

	if modeENV == TEST_MODE {
		fileName = ".env.test"
		rootDir = getEnv("ROOT_DIR", rootDir)
	}

	res := path.Join(rootDir, fileName)

	return res
}

/**
* Функция для получение env значений из .env файла
 */
func ConfigInit() {
	modeENV := getEnv("MODE_ENV", DEVELOPMENT_MODE)

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

type ApplicationConfig struct {
	Mode    string
	RootDir string
}
type Config struct {
	Database    DatabaseConfig
	Application ApplicationConfig
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
		Application: ApplicationConfig{
			Mode:    getEnv("MODE_ENV", DEVELOPMENT_MODE),
			RootDir: getEnv("ROOT_DIR", ""),
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
