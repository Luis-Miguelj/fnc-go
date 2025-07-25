package controller

//@title Criação de Categoria
//@description Endpoint para criar uma nova categoria de finanças
//@version 1.0
//@host localhost:3000
//@BasePath /category

import (
	"backend-go/db"
	"backend-go/infra"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

func CreateCategory(c *fiber.Ctx) error {
	var userId = c.Params("userId")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User ID is required",
		})
	}

	var request CreateCategoryRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	if request.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Category name is required",
		})
	}

	category := infra.Category{
		Id:     uuid.NewString(),
		Name:   request.Name,
		UserId: userId,
	}
	if err := db.DB.Create(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create category",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Category created successfully",
		"category": fiber.Map{
			"id":   category.Id,
			"name": category.Name,
		},
	})
}
