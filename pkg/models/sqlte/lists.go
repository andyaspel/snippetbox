package sqlte

import (
	"fmt"

	"github.com/andyaspel/snippetbox/pkg/models"
	"gorm.io/gorm"
)

type ListModel struct {
	DB *gorm.DB `gorm:"embedded"`
}

// data should have been sanitized, validated, verified, etc before this stage in pipe-line
func (l *ListModel) Insert(title, content string, done bool) (int, error) {
	list := &models.List{Title: title, Content: content, Done: done}
	l.DB.Create(&list) // pass a slice to insert multiple row
	l.DB.Save(&list)
	fmt.Printf("\nNew record created on %v\nID: %d\t\tTitle: %s\nContent: %s\nDone: %t\n",
		list.CreatedAt, list.ID, title, content, done)
	return int(list.ID), nil
}

func (l *ListModel) Get(id int) (*models.List, error) {
	var list models.List
	err := l.DB.Model(&list).Where("id = ?", id).First(&list).Error

	if err != nil {
		return nil, models.ErrorRecord
	}
	return &list, nil
}

func (l *ListModel) Latest() ([]*models.List, error) {
	var list []*models.List
	err := l.DB.Order("id desc").Limit(10).Find(&list).Error
	if err != nil {
		return nil, models.ErrorRecords
	}
	return list, nil
}

func (l *ListModel) Update(id int, title, content string, done bool) error {
	var list models.List
	err := l.DB.Model(&list).Where("id = ?", id).First(&list).Error
	if err != nil {
		return models.ErrorRecord
	}
	err = l.DB.Model(&list).Where("id = ?", id).Update("title", title).Error
	if err != nil {
		return models.ErrorRecord
	}
	err = l.DB.Model(&list).Where("id = ?", id).Update("content", content).Error
	if err != nil {
		return models.ErrorRecord
	}
	err = l.DB.Model(&list).Where("id = ?", id).Update("done", done).Error
	if err != nil {
		return models.ErrorRecord
	}
	return l.DB.Model(&list).Where("id = ?", id).Update("title", title).Error

}
