package main

import (
	"github.com/gin-gonic/gin"
	"log"
)
type Content struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func main() {
	// Create a Gin router
	router := gin.Default()

	// Test server
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong!")
	})

	// API routes
	router.POST("/api/urls/shorten", func(c *gin.Context) {
		// TODO: Shorten the long url
		var content Content
		if err := c.ShouldBindJSON(&content); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, "content okay !")
	})

	router.GET("/api/urls/:id", func(c *gin.Context) {
		id := c.Param("id")
		log.Println(id)
		// TODO: Retrieve the long url
	})

	// Start server on port 8080
	log.Println("Server is live on 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
