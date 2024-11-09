package sqlte

import (
	"github.com/andyaspel/snippetbox/pkg/models"
	"gorm.io/gorm"
)

type SnippetModel struct {
	DB *gorm.DB
}

func (s *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func (s *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (s *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
