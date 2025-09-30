package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/dfunani/AfroChat/backend/lib/utils"
	_ "github.com/lib/pq" // PostgreSQL driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// Connection holds database connection and configuration
type DatabaseConnection struct {
	DB     *gorm.DB
	Config *DatabaseConfig
	SQLDB  *sql.DB
}

// NewConfig creates a new database configuration from environment variables
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     utils.GetEnv("DB_HOST"),
		Port:     utils.GetEnv("DB_PORT"),
		User:     utils.GetEnv("DB_USER"),
		Password: utils.GetEnv("DB_PASSWORD"),
		DBName:   utils.GetEnv("DB_NAME"),
		SSLMode:  utils.GetEnv("DB_SSLMODE"),
	}
}

// NewConnection creates a new database connection
func NewDatabaseConnection(config *DatabaseConfig) (*DatabaseConnection, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql db: %w", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.SetConnMaxIdleTime(1 * time.Minute)

	return &DatabaseConnection{
		DB:     db,
		Config: config,
		SQLDB:  sqlDB,
	}, nil
}

func (c *DatabaseConnection) Close() error {
	if c.SQLDB == nil {
		return nil
	}
	log.Println("Closing database connection...")
	if err := c.SQLDB.Close(); err != nil {
		return fmt.Errorf("failed to close sql db: %w", err)
	}
	log.Println("Database connection closed successfully")
	return nil
}

func (c *DatabaseConnection) Health() error {
	log.Println("Checking database connection health...")
	if c.SQLDB == nil {
		return fmt.Errorf("no database connection found")
	}

	if err := c.SQLDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}
	log.Println("Database connection ping successful")
	return nil
}
