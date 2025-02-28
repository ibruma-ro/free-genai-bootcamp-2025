package routes

import (
	"lang-portal-backend/internal/handlers" // Correct import path

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/api/words", handlers.GetWords)
}
