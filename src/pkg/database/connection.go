package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/dfunani/AfroChat/backend/lib/utils"
	_ "github.com/lib/pq" // PostgreSQL driver
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
	DB     *sql.DB
	Config *DatabaseConfig
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

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DatabaseConnection{
		DB:     db,
		Config: config,
	}, nil
}

// Close closes the database connection
func (c *DatabaseConnection) Close() error {
	return c.DB.Close()
}

// Health checks if the database connection is healthy
func (c *DatabaseConnection) Health() error {
	return c.DB.Ping()
}
