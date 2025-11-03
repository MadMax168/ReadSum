package main

import (
	"log"
	"backend/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func server() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Not found .env file")
	}

	config.ConnectDB();

	config.DB.AutoMigrate(
		//models
	)

	app := fiber.New();

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowCredentials: true,
	}))

	log.Println("Server Starting on: 8000")
	app.Listen(":8000")
}
