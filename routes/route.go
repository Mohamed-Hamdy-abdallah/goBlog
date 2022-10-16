package routes

import (
	"github.com/Mohamed-Hamdy-abdallah/blogbackend/controller"
	"github.com/Mohamed-Hamdy-abdallah/blogbackend/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	app.Use(middleware.IsAuthenticate)

	app.Post("/api/blog", controller.CreateBlog)
	app.Get("/api/blog", controller.AllBlog)
	app.Get("/api/blog/:id", controller.DetailBlog)
	app.Put("/api/blog/:id", controller.UpdateBlog)
	app.Delete("/api/blog/:id", controller.DeleteBlog)

	app.Get("/api/uniqueblog", controller.UniqueBlog)

	app.Post("/api/blog/comment/:id", controller.CreateComment)
	app.Put("/api/blog/:blogid/comment/:id", controller.UpdateComment)
	app.Delete("/api/blog/:blogid/comment/:id", controller.DeleteComment)
	app.Get("/api/blog/comment/:id", controller.AllComments)
	app.Get("/api/blog/:blogid/comment/:id", controller.OneComment)

	app.Post("/api/uplaodimage", controller.Uploadimage)
	app.Static("/api/upload", "./uploads")

}
