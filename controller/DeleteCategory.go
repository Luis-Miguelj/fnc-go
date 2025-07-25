package controller

import (
	"backend-go/db"
	"backend-go/infra"

	"github.com/gofiber/fiber/v2"
)

func DeleteCategory(c *fiber.Ctx) error {
	categoryId := c.Params("id")
	if categoryId == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Category ID is required",
		})
	}

	if err := db.DB.Delete(&infra.Category{}, "id  = ?", categoryId).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete category",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Category deleted successfully",
	})
}
