package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLastStudySession(c *gin.Context) {
	// Logic to fetch the last study session
	lastSession := gin.H{
		"id":                123,
		"group_id":          1,
		"group_name":        "Greetings",
		"start_time":        "2024-03-20T14:30:00Z",
		"study_activity_id": 1,
	}
	c.JSON(http.StatusOK, lastSession)
}
