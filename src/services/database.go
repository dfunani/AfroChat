package services

import (
	"fmt"
	"log"

	"github.com/dfunani/AfroChat/backend/pkg/config"
	"github.com/dfunani/AfroChat/backend/pkg/database"
	"github.com/dfunani/AfroChat/backend/pkg/database/models"
)

func CreateDatabaseClient(appConfig *config.ApplicationConfig) (*database.DatabaseConnection, error) {
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
		return nil, err
	}
	runMigrations(conn)
	log.Println("✅ Database connected successfully")
	return conn, nil
}

func runMigrations(dbConnection *database.DatabaseConnection) error {
	log.Println("Running migrations...")
	err := dbConnection.DB.AutoMigrate(&models.User{})
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	log.Println("✅ Migrations ran successfully")
	return nil
}
