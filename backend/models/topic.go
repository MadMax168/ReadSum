package models

import (
	"time"

	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	Name      string     `json:"name" gorm:"uniqueIndex;not null"`
	Category  string     `json:"category"`
	Documents []Document `json:"documents,omitempty" gorm:"many2many:document_topics;"`
}

type DocumentTopic struct {
	DocumentID     uint    `gorm:"primaryKey"`
	TopicID        uint    `gorm:"primaryKey"`
	RelevanceScore float64 `json:"relevance_score" gorm:"type:decimal(3,2)"`
	CreatedAt      time.Time
}
