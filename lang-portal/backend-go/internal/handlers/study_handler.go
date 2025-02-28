package handlers

import (
	"lang-portal-backend/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetStudyActivity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	activity := services.GetStudyActivity(id)
	c.JSON(http.StatusOK, activity)
}

func GetStudyActivitySessions(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	sessions := services.GetStudyActivitySessions(id)
	c.JSON(http.StatusOK, gin.H{"items": sessions})
}
