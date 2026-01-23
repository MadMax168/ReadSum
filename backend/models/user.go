package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"name" gorm:"not null"`
	Email        string `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string `json:"-" gorm:"not null"`
	//Relationships
	Collection []Collection `json:"collection,omitempty" gorm:"foreignKey:UserID"`
	Docs       []Document   `json:"documents,omitempty" gorm:"foreignKey:UserID"`
}
