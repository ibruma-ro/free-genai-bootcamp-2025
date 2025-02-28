package models

import (
	"time"
)

type Word struct {
	ID      int    `json:"id"`
	French  string `json:"french"`
	English string `json:"english"`
}

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type WordGroup struct {
	ID      int `json:"id"`
	WordID  int `json:"word_id"`
	GroupID int `json:"group_id"`
}

type StudySession struct {
	ID              int       `json:"id"`
	GroupID         int       `json:"group_id"`
	CreatedAt       time.Time `json:"created_at"`
	StudyActivityID int       `json:"study_activity_id"`
}

type StudyActivity struct {
	ID             int       `json:"id"`
	StudySessionID int       `json:"study_session_id"`
	GroupID        int       `json:"group_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type WordReviewItem struct {
	WordID         int       `json:"word_id"`
	StudySessionID int       `json:"study_session_id"`
	Correct        bool      `json:"correct"`
	CreatedAt      time.Time `json:"created_at"`
}

// Response structs based on API specs
type PaginatedResponse struct {
	CurrentPage  int `json:"current_page"`
	TotalPages   int `json:"total_pages"`
	ItemsPerPage int `json:"items_per_page"`
	TotalItems   int `json:"total_items"`
}

type WordResponse struct {
	French       string                 `json:"french"`
	English      string                 `json:"english"`
	CorrectCount int                    `json:"correct_count"`
	WrongCount   int                    `json:"wrong_count"`
	Parts        map[string]interface{} `json:"parts,omitempty"`
	Groups       []Group                `json:"groups,omitempty"`
}

type GroupResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	WordCount int    `json:"word_count"`
}

type StudySessionResponse struct {
	ID            int       `json:"id"`
	ActivityName  string    `json:"activity_name"`
	GroupName     string    `json:"group_name"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	WordsReviewed int       `json:"words_reviewed"`
}

type StudyActivityResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnail_url"`
}

type DashboardStats struct {
	TotalWordsStudied   int `json:"total_words_studied"`
	TotalAvailableWords int `json:"total_available_words"`
	StudyStreakDays     int `json:"study_streak_days"`
}
