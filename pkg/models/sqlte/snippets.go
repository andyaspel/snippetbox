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
	var snippet models.Snippet
	err := s.DB.Model(&snippet).Where("id = ?", id).First(&snippet).Error
	if err != nil {
		return nil, models.ErrorRecord
	}
	return &snippet, nil
}

func (s *SnippetModel) Latest() ([]*models.Snippet, error) {
	var snippets []*models.Snippet
	err := s.DB.Order("id desc").Limit(10).Find(&snippets).Error
	if err != nil {
		return nil, models.ErrorRecords
	}
	return snippets, nil
}

func (s *SnippetModel) Update(id int, title, content, expires string) error {
	var snippet models.Snippet
	err := s.DB.Model(&snippet).Where("id = ?", id).First(&snippet).Error
	if err != nil {
		return models.ErrorRecord
	}
	err = s.DB.Model(&snippet).Where("id = ?", id).Update("title", title).Error
	if err != nil {
		return models.ErrorRecord
	}
	err = s.DB.Model(&snippet).Where("id = ?", id).Update("content", content).Error
	if err != nil {
		return models.ErrorRecord
	}
	err = s.DB.Where("id = ?", id).Update("expires", expires).Error
	if err != nil {
		return models.ErrorRecord
	}
	return s.DB.Model(&snippet).Where("id = ?", id).Update("title", title).Error

}
