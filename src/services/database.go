package services

import (
	"log"

	"github.com/dfunani/AfroChat/backend/pkg/config"
	"github.com/dfunani/AfroChat/backend/pkg/database"
)

func InitDatabase(appConfig *config.ApplicationConfig, dbConnection *database.DatabaseConnection) error {
	dbConfig := &database.DatabaseConfig{
		Host:     appConfig.DBHost,
		Port:     appConfig.DBPort,
		User:     appConfig.DBUser,
		Password: appConfig.DBPass,
		DBName:   appConfig.DBName,
		SSLMode:  appConfig.DBSSL,
	}

	conn, err := database.NewDatabaseConnection(dbConfig)
	if err != nil {
		return err
	}
	*dbConnection = *conn
	log.Println("✅ Database connected successfully")
	return nil
}
