package config

import (
	"fmt"

	"github.com/dfunani/AfroChat/backend/lib/utils"
)

// AppConfig holds application configuration
type ApplicationConfig struct {
	Port   string
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
	DBSSL  string
	Env    string
}

// LoadConfig loads configuration from environment variables
func LoadApplicationConfig() *ApplicationConfig {
	fmt.Println("Loading Application config...")
	return &ApplicationConfig{
		Port:   utils.GetEnv("PORT"),
		DBHost: utils.GetEnv("DB_HOST"),
		DBPort: utils.GetEnv("DB_PORT"),
		DBUser: utils.GetEnv("DB_USER"),
		DBPass: utils.GetEnv("DB_PASSWORD"),
		DBName: utils.GetEnv("DB_NAME"),
		DBSSL:  utils.GetEnv("DB_SSLMODE"),
		Env:    utils.GetEnv("ENVIRONMENT"),
	}
}
