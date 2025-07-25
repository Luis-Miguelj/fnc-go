package controller

import (
	"backend-go/db"
	"backend-go/infra"
	"backend-go/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @title Registro de Usuário
// @description Endpoint para registrar um novo usuário
// @version 1.0
// @host localhost:3000
// @BasePath /register

func RegisterUser(c *fiber.Ctx) error {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	if input.Name == "" || input.Email == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	hashPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	user := infra.User{
		Id:        uuid.NewString(),
		Name:      input.Name,
		Email:     input.Email,
		Password:  hashPassword,
		CreatedAt: utils.GetCurrentTime(),
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"user": fiber.Map{
			"id":         user.Id,
			"name":       user.Name,
			"email":      user.Email,
			"created_at": user.CreatedAt,
		},
	})
}
