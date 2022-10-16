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

func CreateComment(c *fiber.Ctx) error {
	blogid, _ := strconv.Atoi(c.Params("id"))

	cookie := c.Cookies("jwt")
	userid, _ := util.ParseJwt(cookie)

	comment := models.Comment{
		BlogID: strconv.Itoa(blogid),
		UserID: string(userid),
	}
	// var comment models.Comment
	if err := c.BodyParser(&comment); err != nil {
		fmt.Println("unable to parse body ")
	}
	if err := database.DB.Create(&comment).Error; err != nil {
		return c.JSON(fiber.Map{
			"message": "Invalid Payload",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Comment created successfully",
	})
}

func AllComments(c *fiber.Ctx) error {
	blogid, _ := strconv.Atoi(c.Params("id"))
	var getcomments []models.Comment

	database.DB.Model(&getcomments).Where("blog_id=?", blogid).Preload("User").Preload("Blog").Find(&getcomments)
	return c.JSON(
		fiber.Map{
			"data": getcomments,
		})
}

func OneComment(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blogid, _ := strconv.Atoi(c.Params("blogid"))
	var comment models.Comment
	database.DB.Where("blog_id=?", blogid).Where("id=?", id).Preload("User").Preload("Blog").First(&comment)

	return c.JSON(fiber.Map{
		"data": comment,
	})
}

func UpdateComment(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blogid, _ := strconv.Atoi(c.Params("blogid"))

	var oneComment models.Comment
	database.DB.Where("blog_id=?", blogid).Where("id=?", id).Preload("User").Preload("Blog").First(&oneComment)

	cookie := c.Cookies("jwt")
	userid, _ := util.ParseJwt(cookie)

	if oneComment.UserID == userid {
		comment := models.Comment{
			Id:     uint(id),
			BlogID: strconv.Itoa(blogid),
		}
		if err := c.BodyParser(&comment); err != nil {
			fmt.Println("unable to parse body ")
		}

		database.DB.Model(&comment).Updates(comment)
		return c.JSON(comment)
	} else {
		return c.JSON(fiber.Map{
			"message": "you can't update this COMMENT",
		})
	}

}

func DeleteComment(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blogid, _ := strconv.Atoi(c.Params("blogid"))

	var oneComment models.Comment
	database.DB.Where("blog_id=?", blogid).Where("id=?", id).Preload("User").Preload("Blog").First(&oneComment)

	cookie := c.Cookies("jwt")
	userid, _ := util.ParseJwt(cookie)

	if oneComment.UserID == userid {
		comment := models.Comment{
			Id:     uint(id),
			BlogID: strconv.Itoa(blogid),
		}
		deleteQuery := database.DB.Delete(&comment)
		if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
			c.Status(400)
			return c.JSON(fiber.Map{
				"message": "comment NOT found",
			})
		}
		return c.JSON(fiber.Map{
			"message": "COMMENT deleted successfully",
		})
	} else {
		return c.JSON(fiber.Map{
			"message": "You can't delete this COMMENT",
		})
	}

}
