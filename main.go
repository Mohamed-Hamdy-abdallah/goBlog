package main

import (
	"log"

	"github.com/Mohamed-Hamdy-abdallah/blogbackend/database"
	"github.com/Mohamed-Hamdy-abdallah/blogbackend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}
	port := "3000"
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)
}
