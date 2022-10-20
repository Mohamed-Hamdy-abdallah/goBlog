package controller

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Mohamed-Hamdy-abdallah/blogbackend/database"
	"github.com/Mohamed-Hamdy-abdallah/blogbackend/models"
	"github.com/Mohamed-Hamdy-abdallah/blogbackend/util"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse the body")
	}

	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password must be more than 6 characters",
		})
	}

	if !ValidateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid Email",
		})
	}

	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)

	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email already exist",
		})
	}

	user := models.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Phone:     data["phone"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
	}

	user.SetPassword(data["password"].(string))

	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "User created successfully",
	})

}

var ctx *gin.Context

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse the body")
	}
	var user models.User
	database.DB.Where("email=?", strings.TrimSpace(data["email"])).First(&user)

	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Email Address doesn't exist you must sign up",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	cookie := fiber.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
	}
	c.Cookie(&cookie)
	// ctx.SetCookie("jwt", token, 60*60*24, "/", "localhost", false, true)
	return c.JSON(fiber.Map{
		"message": "you have succesfully login",
		"user":    user,
		"token":   token,
	})
}

func GetUser(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	userid, _ := util.ParseJwt(cookie)

	var user models.User
	database.DB.Where("id=?", strings.TrimSpace(userid)).First(&user)

	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "No log in User",
		})
	}

	return c.JSON(fiber.Map{
		"user": user,
	})
}

func GetToken(c *fiber.Ctx) error {
	// cookie := c.Cookies("jwt")
	return c.JSON(fiber.Map{
		"token": "123",
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Expires:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		Value:    "",
		HTTPOnly: true,
		// Path:     "/api/auth",
		Domain: "/localhost",
	}
	c.Cookie(&cookie)
	// c.ClearCookie("jwt")
	// ctx.SetCookie("jwt", "", -1, "/", "localhost", false, true)

	return c.JSON(fiber.Map{
		"message": "Logged out",
	})
}

type Claims struct {
	jwt.StandardClaims
}
