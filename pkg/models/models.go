package models

import (
	"errors"

	"gorm.io/gorm"
)

var ErrorRecord = errors.New("models: No matching record found")
var ErrorRecords = errors.New("models: No records found - table is empty")

type Snippet struct {
	gorm.Model
	Title   string
	Content string
	Expires string
}

type List struct {
	gorm.Model
	Title   string
	Content string
	Done    bool
}

// Log model for DB
// Add to models.go
type Log struct {
	gorm.Model
	Time      string
	Method    string
	URL       string
	Status    int
	Duration  string
	Remote    string
	UserAgent string
}
