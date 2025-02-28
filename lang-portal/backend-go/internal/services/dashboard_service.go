package services

import (
	"lang-portal-backend/internal/models" // Update with your actual project path
)

// GetDashboardStats retrieves statistics for the dashboard
func GetDashboardStats() models.DashboardStats {
	// Logic to calculate stats (this is a placeholder)
	return models.DashboardStats{
		TotalWordsStudied:   500,
		TotalAvailableWords: 1000,
		StudyStreakDays:     4,
	}
}
