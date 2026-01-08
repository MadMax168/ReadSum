package models

import (
	"time"

	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	UserID    uint      `json:"user_id"`
	Title     string    `json:"title"      gorm:"not null"`
	FileType  string    `json:"file_type"`
	FileUrl   string    `json:"file_url"   gorm:"type:text"`
	Summary   string    `json:"summary"    gorm:"type:jsonb"`
	UpdatedAt time.Time `json:"upload_date"`
	Status    string    `json:"status"`
}