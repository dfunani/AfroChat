package main

import "github.com/gin-gonic/gin"

func main() {
	// Entry point of the backend application
	println("Backend service is running...")

	router := gin.Default()

	router.GET("/api/v1/health", healthCheck)

	router.Run(":8080")
}

func healthCheck(context *gin.Context) {
	context.JSON(200, gin.H{"status": "ok"})
}
