package database

import (
	"log"
	"os"

	"github.com/Mohamed-Hamdy-abdallah/blogbackend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading env file")
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to DB")
	} else {
		log.Println("Connecetd succesfully to DB")
	}
	DB = database
	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
		&models.Comment{},
	)
}
