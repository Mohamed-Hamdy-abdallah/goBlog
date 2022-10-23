package main

import (
	"log"
	"os"

	"github.com/Mohamed-Hamdy-abdallah/blogbackend/database"
	"github.com/Mohamed-Hamdy-abdallah/blogbackend/routes"
	"github.com/gofiber/fiber/v2"

	// "github.com/rs/cors"

	// "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/joho/godotenv"
)

// func EnableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// }

// func handleArticles(w http.ResponseWriter, r *http.Request) {
// 	EnableCors(&w)
// 	js, err := json.Marshal(Articles)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(js)
// }

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}
	port := os.Getenv("PORT")
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin,Authorization,withCredentials",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// app.Use(cors.New)

	routes.Setup(app)
	app.Listen(":" + port)
}
