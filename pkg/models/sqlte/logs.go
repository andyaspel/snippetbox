package sqlte

import (
	"fmt"

	"github.com/andyaspel/snippetbox/pkg/models"
	"gorm.io/gorm"
)

type LogModel struct {
	DB *gorm.DB `gorm:"embedded"`
}

// data should have been sanitized, validated, verified, etc before this stage in pipe-line
func (l *LogModel) Insert(time, method, url string, status int, duration, remote, userAgent string) (int, error) {
	log := &models.Log{Time: time, Method: method, URL: url, Status: status, Duration: duration, Remote: remote, UserAgent: userAgent}
	l.DB.Create(&log) // pass a slice to insert multiple row
	l.DB.Save(&log)
	fmt.Printf("\nNew record created on %v\nID: %d\t\tTime: %s\nMethod: %s\nStatus: %d\nUrl: %s\nDuration: %s\n",
		log.CreatedAt, log.ID, time, method, status, url, duration)
	return int(log.ID), nil
}

func (l *LogModel) Get(id int) (*models.Log, error) {
	var log models.Log
	err := l.DB.Model(&log).Where("id = ?", id).First(&log).Error

	if err != nil {
		return nil, models.ErrorRecord
	}
	return &log, nil
}

func (l *LogModel) Latest() ([]*models.Log, error) {
	var log []*models.Log
	err := l.DB.Order("id desc").Limit(10).Find(&log).Error
	if err != nil {
		return nil, models.ErrorRecords
	}
	return log, nil
}

func (l *LogModel) Update(id int, title, content string, done bool) error {
	var log models.Log
	err := l.DB.Model(&log).Where("id = ?", id).First(&log).Error
	if err != nil {
		return models.ErrorRecord
	}
	err = l.DB.Model(&log).Where("id = ?", id).Update("title", title).Error
	if err != nil {
		return models.ErrorRecord
	}
	err = l.DB.Model(&log).Where("id = ?", id).Update("content", content).Error
	if err != nil {
		return models.ErrorRecord
	}
	err = l.DB.Model(&log).Where("id = ?", id).Update("done", done).Error
	if err != nil {
		return models.ErrorRecord
	}
	return l.DB.Model(&log).Where("id = ?", id).Update("title", title).Error

}
