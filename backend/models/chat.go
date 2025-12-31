package models

import (
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Title   string `json:"title"`
	UserID  uint
	Message []Message
}
