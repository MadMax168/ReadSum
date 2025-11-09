package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model //ID //CreatedAt //UpdatesAt //DeletedAt
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
}