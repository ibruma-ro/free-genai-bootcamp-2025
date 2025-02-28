package handlers

import (
	"net/http"
	"strconv"

	"lang-portal-backend/internal/services" // Update with your actual project path

	"github.com/gin-gonic/gin"
)

func GetWords(c *gin.Context) {
	words, err := services.GetWords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch words"})
		return
	}
	c.JSON(http.StatusOK, words)
}

func GetWord(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	word, err := services.GetWord(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Word not found"})
		return
	}
	c.JSON(http.StatusOK, word)
}
