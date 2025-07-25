package controller

import (
	"backend-go/db"
	"backend-go/infra"

	"github.com/gofiber/fiber/v2"
)

type UpdateCategoryRequest struct {
	Name string `json:"name"`
}

func UpdateCategory(c *fiber.Ctx) error {

	var request UpdateCategoryRequest
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

	categoryId := c.Params("id")
	if categoryId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Category ID is required",
		})
	}

	if err := db.DB.Model(&infra.Category{}).Where("id = ?", categoryId).Update("Name", request.Name).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update category",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Category updated successfully",
	})

}
