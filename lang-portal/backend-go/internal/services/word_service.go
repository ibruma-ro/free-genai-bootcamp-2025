package services

import (
	"lang-portal-backend/internal/models"
	"lang-portal-backend/internal/models/database"
)

// GetWords retrieves all words from the database
func GetWords() ([]models.Word, error) {
	rows, err := database.DB.Query("SELECT id, french, english FROM words")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var words []models.Word
	for rows.Next() {
		var word models.Word
		if err := rows.Scan(&word.ID, &word.French, &word.English); err != nil {
			return nil, err
		}
		words = append(words, word)
	}
	return words, nil
}

// GetWord retrieves a specific word by ID from the database
func GetWord(id int) (models.Word, error) {
	var word models.Word
	err := database.DB.QueryRow("SELECT id, french, english FROM words WHERE id = ?", id).Scan(&word.ID, &word.French, &word.English)
	if err != nil {
		return word, err
	}
	return word, nil
}
