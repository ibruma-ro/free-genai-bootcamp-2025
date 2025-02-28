package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

type Word struct {
	French  string `json:"french"`
	English string `json:"english"`
}

type Group struct {
	Name string `json:"name"`
}

type StudySession struct {
	GroupID         int    `json:"group_id"`
	CreatedAt       string `json:"created_at"`
	StudyActivityID int    `json:"study_activity_id"`
}

type StudyActivity struct {
	StudySessionID int    `json:"study_session_id"`
	GroupID        int    `json:"group_id"`
	CreatedAt      string `json:"created_at"`
}

type WordGroup struct {
	WordID  int `json:"word_id"`
	GroupID int `json:"group_id"`
}

type WordReviewItem struct {
	WordID         int    `json:"word_id"`
	StudySessionID int    `json:"study_session_id"`
	Correct        bool   `json:"correct"`
	CreatedAt      string `json:"created_at"`
}

func main() {
	db, err := sql.Open("sqlite3", "./words.db") // Update with your actual database path
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	seedData(db, "words.json", "words")
	seedData(db, "groups.json", "groups")
	seedData(db, "study_sessions.json", "study_sessions")
	seedData(db, "study_activities.json", "study_activities")
	seedData(db, "word_groups.json", "word_groups")
	seedData(db, "word_review_items.json", "word_review_items")
}

func seedData(db *sql.DB, filename string, table string) {
	filePath := filepath.Join("db", "seeds", "json", filename)
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return
	}

	switch table {
	case "words":
		var words []Word
		json.Unmarshal(file, &words)
		for _, word := range words {
			_, err := db.Exec("INSERT INTO words (french, english) VALUES (?, ?)", word.French, word.English)
			if err != nil {
				fmt.Printf("Error inserting word: %v\n", err)
			}
		}
	case "groups":
		var groups []Group
		json.Unmarshal(file, &groups)
		for _, group := range groups {
			_, err := db.Exec("INSERT INTO groups (name) VALUES (?)", group.Name)
			if err != nil {
				fmt.Printf("Error inserting group: %v\n", err)
			}
		}
	case "study_sessions":
		var sessions []StudySession
		json.Unmarshal(file, &sessions)
		for _, session := range sessions {
			_, err := db.Exec("INSERT INTO study_sessions (group_id, created_at, study_activity_id) VALUES (?, ?, ?)", session.GroupID, session.CreatedAt, session.StudyActivityID)
			if err != nil {
				fmt.Printf("Error inserting study session: %v\n", err)
			}
		}
	case "study_activities":
		var activities []StudyActivity
		json.Unmarshal(file, &activities)
		for _, activity := range activities {
			_, err := db.Exec("INSERT INTO study_activities (study_session_id, group_id, created_at) VALUES (?, ?, ?)", activity.StudySessionID, activity.GroupID, activity.CreatedAt)
			if err != nil {
				fmt.Printf("Error inserting study activity: %v\n", err)
			}
		}
	case "word_groups":
		var wordGroups []WordGroup
		json.Unmarshal(file, &wordGroups)
		for _, wg := range wordGroups {
			_, err := db.Exec("INSERT INTO word_groups (word_id, group_id) VALUES (?, ?)", wg.WordID, wg.GroupID)
			if err != nil {
				fmt.Printf("Error inserting word group: %v\n", err)
			}
		}
	case "word_review_items":
		var reviewItems []WordReviewItem
		json.Unmarshal(file, &reviewItems)
		for _, item := range reviewItems {
			_, err := db.Exec("INSERT INTO word_review_items (word_id, study_session_id, correct, created_at) VALUES (?, ?, ?, ?)", item.WordID, item.StudySessionID, item.Correct, item.CreatedAt)
			if err != nil {
				fmt.Printf("Error inserting word review item: %v\n", err)
			}
		}
	}
}
