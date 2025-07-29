package controller

import (
	"backend-go/db"
	"backend-go/infra"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type FinanceUpdateRequest struct {
	Type       string  `json:"type"`
	CategoryId string  `json:"categoryId"`
	Value      float64 `json:"value"`
}

func UpdateFinance(c *fiber.Ctx) error {
	financeId := c.Params("financeId")
	if financeId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Finance ID is required",
		})
	}

	var request FinanceUpdateRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	var typeFinance string

	if request.Type == "e" {
		typeFinance = "Entrada"
	} else {
		typeFinance = "Saída"
	}

	fmt.Println(typeFinance)
	if err := db.DB.Model(&infra.Finance{}).Where("id = ?", financeId).Update("Type", typeFinance).Update("CategoryId", request.CategoryId).Update("Value", request.Value).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update finance record",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Finança atualizada com sucesso!",
		"type":       typeFinance,
		"categoryId": request.CategoryId,
		"value":      request.Value,
	})
}
