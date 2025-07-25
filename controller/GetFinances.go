package controller

import (
	"backend-go/db"
	"backend-go/infra"

	"github.com/gofiber/fiber/v2"
)

func GetFinances(c *fiber.Ctx) error {
	userId := c.Params("userId")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User ID is required",
		})
	}

	var finances []infra.Finance
	if err := db.DB.Where("user_id = ?", userId).Find(&finances).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve finances",
		})
	}

	return c.JSON(finances)
}
