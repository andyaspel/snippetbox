package models

import (
	"errors"
	"time"
)

var ErrorRecord = errors.New("models: No mathching record found")

type Snippet struct {
	ID        uint
	Title     string
	Content   string
	CreatedAt time.Time
	Expires   time.Time
}
