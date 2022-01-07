package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mayckol/api_test_go/database"
	"github.com/mayckol/api_test_go/routes"
)

func main() {
	database.Connect()
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":3000")
}
