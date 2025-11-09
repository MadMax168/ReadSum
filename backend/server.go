package main

import (
	"log"

	"github.com/MadMax168/Readsum/config"
	"github.com/MadMax168/Readsum/models"
	"github.com/MadMax168/Readsum/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found")
	}
	
	config.ConnectDB();

	config.DB.AutoMigrate(
		&models.User{},
	)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	routes.SetAllRoutes(app);

	log.Fatal(app.Listen(":8080"))
}
