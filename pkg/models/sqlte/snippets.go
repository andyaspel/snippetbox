package sqlte

import (
	"fmt"

	"github.com/andyaspel/snippetbox/pkg/models"
	"gorm.io/gorm"
)

type SnippetModel struct {
	DB *gorm.DB `gorm:"embedded"`
}

func (s *SnippetModel) Insert(title, content, expires string) (int, error) {
	snippet := &models.Snippet{Title: title, Content: content, Expires: expires}
	s.DB.Create(&snippet) // pass a slice to insert multiple row
	s.DB.Save(&snippet)
	fmt.Printf("\nNew record created on %v\nID: %d\t\tTitle: %s\nContent: %s\nExpires: %s\n", snippet.CreatedAt, snippet.ID, title, content, expires)
	return int(snippet.ID), nil
}

func (s *SnippetModel) Get(id int) (*models.Snippet, error) {
	var result models.Snippet
	snippet := &models.Snippet{}
	result.ID = uint(id)
	snippet.ID = result.ID
	notFound := models.ErrorRecord
	err := s.DB.Model(&snippet).First(&result, id)
	if err != nil {
		return &result, notFound
	}
	fmt.Printf("\nID: %d\nTitle: %s\nContent: %s\nExpires: %s\n", result.ID, result.Title, result.Content, result.Expires)

	return &result, nil
}

func (s *SnippetModel) Latest() ([]*models.Snippet, error) {
	results := []*models.Snippet{}
	snippet := &models.Snippet{}
	s.DB.Model(&snippet).Limit(10).Order("id desc").Find(&results)
	if len(results) < 1 {
		return results, models.ErrorRecords
	}
	results = append(results, &models.Snippet{})
	return results, nil
}
