package controller

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/Mohamed-Hamdy-abdallah/blogbackend/database"
	"github.com/Mohamed-Hamdy-abdallah/blogbackend/models"
	"github.com/Mohamed-Hamdy-abdallah/blogbackend/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateBlog(c *fiber.Ctx) error {
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("unable to parse body ")
	}
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
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var getblog []models.Blog
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getblog)
	database.DB.Model(&models.Blog{}).Count(&total)

	return c.JSON(
		fiber.Map{
			"data": getblog,
			"meta": fiber.Map{
				"total":     total,
				"page":      page,
				"last_page": math.Ceil(float64(int(total) / limit)),
			},
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
