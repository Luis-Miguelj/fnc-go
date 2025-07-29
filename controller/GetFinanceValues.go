package controller

import (
	"backend-go/db"
	"backend-go/infra"
	"backend-go/utils"

	"github.com/gofiber/fiber/v2"
)

func GetFinanceValues(c *fiber.Ctx) error {
	userId := c.Params("userId")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).SendString("User ID is required")
	}

	finance := []infra.Finance{}
	if err := db.DB.Find(&finance).Where("user_id = ?", userId).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve finance values",
		})
	}

	var totalValueEntrada float64
	typeValueEntrada := "Entrada"
	for _, f := range finance {
		if f.Type == typeValueEntrada {
			totalValueEntrada += f.Value
		}

	}

	var totalValueSaida float64
	typeValueSaida := "Saída"

	for _, f := range finance {
		if f.Type == typeValueSaida {
			totalValueSaida += f.Value
		}
	}

	valorTotal := totalValueEntrada - totalValueSaida

	format := utils.ArredondarParaDuasCasas(valorTotal)

	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{
			"entrada":    totalValueEntrada,
			"saida":      totalValueSaida,
			"valorTotal": format,
		},
	)
}
