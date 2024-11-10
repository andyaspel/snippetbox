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
	fmt.Println("ID:", snippet.ID)
	return int(snippet.ID), nil
}

func (s *SnippetModel) Get(id int) (*models.Snippet, error) {
	var result models.Snippet
	// var err error
	snippet := &models.Snippet{}
	result.ID = uint(id)
	snippet.ID = result.ID
	err := s.DB.Model(&snippet).First(&result)
	// err := models.ErrorRecord
	if err != nil {
		fmt.Println("hello", result)
		return &result, gorm.ErrRecordNotFound
	}
	fmt.Printf("\nTEST:\n%v\n%v\n%v\n", result.ID, snippet.Title, snippet.Content)
	return snippet, nil

}

func (s *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
