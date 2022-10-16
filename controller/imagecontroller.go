package controller

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

var letters = []rune("abcdefghijklmnopqrstuvwx")

func randletters(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Uploadimage(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files, _ := form.File["image"]
	fileName := ""
	for _, file := range files {
		fileName = randletters(5) + "-" + file.Filename
		if err := c.SaveFile(file, "./uploads/"+fileName); err != nil {
			return err
		}
	}

	return c.JSON(fiber.Map{
		"url": "http://localhost:5000/api/upload/" + fileName,
	})
}
