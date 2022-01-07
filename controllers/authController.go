package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mayckol/api_test_go/database"
	"github.com/mayckol/api_test_go/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	// password precisa ser convertido para um array de bites para GenerateFromPassword
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	database.DB.Create(&user)
	return c.JSON(user)
}
