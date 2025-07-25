package controller

import (
	"backend-go/db"
	"backend-go/infra"
	"backend-go/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateFinanceRequest struct {
	Type       string  `json:"type"`
	Value      float64 `json:"value"`
	CategoryId string  `json:"categoryId"`
}

func CreateFinance(c *fiber.Ctx) error {
	var request CreateFinanceRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	if err := c.Params("userId"); err == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User ID is required",
		})
	}

	if request.Type == "" || request.Value <= 0 || request.CategoryId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Type, value (greater than 0), and category ID are required",
		})
	}

	var financeType string

	if request.Type == "e" {
		financeType = "Entrada"
	} else {
		financeType = "Saída"
	}

	finance := infra.Finance{
		Id:         uuid.NewString(),
		UserId:     c.Params("userId"),
		Type:       financeType,
		Value:      request.Value,
		CategoryId: request.CategoryId,
		CreatedAt:  utils.GetCurrentTime(),
	}

	if err := db.DB.Create(&finance).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create finance record",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Finance record created successfully",
		"finance": fiber.Map{
			"id":         finance.Id,
			"userId":     finance.UserId,
			"type":       finance.Type,
			"value":      finance.Value,
			"categoryId": finance.CategoryId,
			"createdAt":  finance.CreatedAt,
		},
	},
	)
}
