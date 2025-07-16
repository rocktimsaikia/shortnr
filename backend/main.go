package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Create a Gin router
	router := gin.Default()

	// Test server
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong!")
	})

	// Start server on port 8080
	log.Println("Server is live on 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
