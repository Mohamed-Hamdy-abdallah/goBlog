package controller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Mohamed-Hamdy-abdallah/blogbackend/database"
	"github.com/Mohamed-Hamdy-abdallah/blogbackend/models"
	"github.com/Mohamed-Hamdy-abdallah/blogbackend/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateBlog(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	userid, _ := util.ParseJwt(cookie)
	fmt.Println(userid)
	blogpost := models.Blog{
		UserID: string(userid),
	}
	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("unable to parse body ")
	}
	// fmt.Print(blogpost)
	if err := database.DB.Create(&blogpost).Error; err != nil {
		return c.JSON(fiber.Map{
			"message": "Invalid Payload",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Post created successfully",
	})
}

func AllBlog(c *fiber.Ctx) error {
	// page, _ := strconv.Atoi(c.Query("page", "1"))

	var getblog []models.Blog
	database.DB.Preload("User").Find(&getblog)
	// database.DB.Model(&models.Blog{}).Count(&total)

	return c.JSON(
		fiber.Map{
			"data": getblog,
		})
}

func DetailBlog(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blogpost models.Blog
	database.DB.Where("id=?", id).Preload("User").First(&blogpost)

	return c.JSON(fiber.Map{
		"data": blogpost,
	})
}

func UpdateBlog(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var blogpost models.Blog
	database.DB.Where("id=?", id).Preload("User").First(&blogpost)

	cookie := c.Cookies("jwt")
	userid, _ := util.ParseJwt(cookie)

	if blogpost.UserID == userid {

		blog := models.Blog{
			Id: uint(id),
		}
		if err := c.BodyParser(&blog); err != nil {
			fmt.Println("unable to parse body ")
		}
		database.DB.Model(&blog).Updates(blog)
		return c.JSON(blog)
	} else {
		return c.JSON(fiber.Map{
			"message": "You are not authorized to Update this Post",
		})
	}

}

func UniqueBlog(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.ParseJwt(cookie)

	var blog []models.Blog
	database.DB.Model(&blog).Where("user_id=?", id).Preload("User").Find(&blog)
	return c.JSON(blog)
}

func DeleteBlog(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))
	var blogpost models.Blog
	database.DB.Where("id=?", id).Preload("User").First(&blogpost)

	cookie := c.Cookies("jwt")
	userid, _ := util.ParseJwt(cookie)

	if blogpost.UserID == userid {
		blog := models.Blog{
			Id: uint(id),
		}
		deleteQuery := database.DB.Delete(&blog)
		if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
			c.Status(400)
			return c.JSON(fiber.Map{
				"message": "record NOT found",
			})
		}
		return c.JSON(fiber.Map{
			"message": "POST deleted successfully",
		})
	} else {
		return c.JSON(fiber.Map{
			"message": "You are not authorized to delete this Post",
		})
	}

}
