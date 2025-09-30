package main

import (
	"log"

	"github.com/dfunani/AfroChat/backend/pkg/config"
	"github.com/dfunani/AfroChat/backend/pkg/database"
	"github.com/dfunani/AfroChat/backend/services"
	"github.com/gin-gonic/gin"
)

var (
	dbConnection *database.DatabaseConnection
	appConfig    *config.ApplicationConfig
)

func main() {
	// Load configuration
	appConfig = config.LoadApplicationConfig()

	// Initialize database connection
	if err := services.InitDatabase(appConfig, dbConnection); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer dbConnection.Close()

	// Set Gin mode based on environment
	if appConfig.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create router
	router := gin.Default()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(services.CorsMiddleware())

	// Health check endpoints
	router.GET("/api/v1/health", services.HealthCheck)
	router.GET("/api/v1/health/db", func(c *gin.Context) { services.DatabaseHealthCheck(c, dbConnection) })

	// Start server
	log.Printf("🚀 AfroChat Backend starting on port %s", appConfig.Port)
	log.Printf("📊 Database: %s:%s/%s", appConfig.DBHost, appConfig.DBPort, appConfig.DBName)

	if err := router.Run(":" + appConfig.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
