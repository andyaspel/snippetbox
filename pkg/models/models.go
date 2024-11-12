package models

import (
	"errors"

	"gorm.io/gorm"
)

var ErrorRecord = errors.New("models: No mathching record found")
var ErrorRecords = errors.New("models: No records found - table is empty")

type Snippet struct {
	gorm.Model
	Title   string
	Content string
	Expires string
}
