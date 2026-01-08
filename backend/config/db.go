package config

import (
	"fmt"
	"os"

	"github.com/MadMax168/Readsum/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
) 

var DB *gorm.DB 

func ConnectDB() { 
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", 
		os.Getenv("DB_HOST"), 
		os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"), 
		os.Getenv("DB_NAME"), 
	) 
	
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) 
	
	if err != nil { 
		panic("Can not connect DB") 
	}

	database.AutoMigrate(&models.Document{}, &models.User{})
	
	DB = database 
}