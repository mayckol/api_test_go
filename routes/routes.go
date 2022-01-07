package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mayckol/api_test_go/controllers"
)

func Register(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
}
