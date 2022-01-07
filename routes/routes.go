package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mayckol/api_test_go/controllers"
)

func Register(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
}
