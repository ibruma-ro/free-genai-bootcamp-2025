package main

import (
	"lang-portal-backend/config"
	"lang-portal-backend/internal/handlers"
	"lang-portal-backend/internal/models/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize the database
	database.InitDB(cfg.DatabaseURL)

	// Set up the Gin router
	router := gin.Default()

	// Enable CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Set up API routes
	api := router.Group("/api")
	{
		// Dashboard routes
		api.GET("/dashboard/last_study_session", handlers.GetLastStudySession)

		// Study routes
		api.GET("/study_activities/:id", handlers.GetStudyActivity)
		api.GET("/study_activities/:id/sessions", handlers.GetStudyActivitySessions)

		// Word routes
		api.GET("/words", handlers.GetWords)
		api.GET("/words/:id", handlers.GetWord)
	}

	// Start the server
	router.Run(":" + cfg.Port)
}
