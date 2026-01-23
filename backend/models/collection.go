package models

import (
	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null"`
	Desc   string `json:"desc" gorm:"type:text"`
	UserID uint   `json:"user_id" gorm:"not null;index"`
	User   User   `json:"user,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	//Relationships
	Documents []Document `json:"documents,omitempty" gorm:"foreignKey:ChatID"`
}
