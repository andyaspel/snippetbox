package sqlte

import (
	"fmt"
	"log"

	"github.com/andyaspel/snippetbox/pkg/models"
	"gorm.io/gorm"
)

type SnippetModel struct {
	DB *gorm.DB `gorm:"embedded"`
}

func (s *SnippetModel) Insert(title, content, expires string) (int, error) {
	var Snippet models.Snippet
	err := s.DB.AutoMigrate(&Snippet)
	if err != nil {
		log.Fatal(err)
	}
	snippet := &models.Snippet{Title: title, Content: content, Expires: expires}

	s.DB.Create(&snippet) // pass a slice to insert multiple row

	fmt.Println("ID:", snippet.ID)
	return int(snippet.ID), nil
}

func (s *SnippetModel) Get(id int) (*models.Snippet, error) {
	var result models.Snippet
	// result.ID = int(id)
	snippet := &models.Snippet{}
	// snippet := &models.Snippet{id}
	s.DB.Model(&snippet).First(&result)
	//	Db.models.snippet(&result)

	// if err == sql.ErrNoRows {
	// 	return nil, models.ErrNoRecord
	// } else if err != nil {
	// 	return nil, err
	// }

	return &result, nil
}

func (s *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
