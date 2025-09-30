package services

import (
	"net/http"

	"github.com/dfunani/AfroChat/backend/pkg/database"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"service":   "AfroChat Backend",
		"timestamp": gin.H{},
	})
}

func DatabaseHealthCheck(c *gin.Context, dbConnection *database.DatabaseConnection) {
	if err := dbConnection.Health(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"database": gin.H{
			"host": dbConnection.Config.Host,
			"port": dbConnection.Config.Port,
			"name": dbConnection.Config.DBName,
		},
	})
}
