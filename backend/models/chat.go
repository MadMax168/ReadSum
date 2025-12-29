package model

import (
	"gorm.io/gorm"
)

type chat struct {
	gorm.Model
	Title   string `json:"title"`
	UserID  uint
	Message []message
}
