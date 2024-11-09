package models

import (
	"errors"

	"gorm.io/gorm"
)

var ErrorRecord = errors.New("models: No mathching record found")

type Snippet struct {
	gorm.Model
	Title   string
	Content string
	Expires string
}
