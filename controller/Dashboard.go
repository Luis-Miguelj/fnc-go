package controller

import (
	"backend-go/db"
	"backend-go/infra"
	"time"

	"github.com/gofiber/fiber/v2"
)

type DashboardType struct {
	Month    string
	Entradas float64
	Saidas   float64
}

func Dashboard(c *fiber.Ctx) error {
	userId := c.Params("userId")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).SendString("User ID is required")
	}

	finance := []infra.Finance{}

	if err := db.DB.Find(&finance).Where("user_id = ?", userId).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving finance data")
	}

	anoFiltrado := time.Now().Year()
	mesAtual := int(time.Now().Month())
	dashboard := []DashboardType{}

	for mes := 1; mes <= mesAtual; mes++ {
		var totalEntradasMes float64
		var totalSaidasMes float64
		for _, f := range finance {
			data := f.CreatedAt
			if data.Year() == anoFiltrado && int(data.Month()) == mes {
				if f.Type == "Entrada" {
					totalEntradasMes += f.Value
				}

				if f.Type == "Saída" {
					totalSaidasMes += f.Value
				}
			}

		}

		nomeMesPtBr := map[int]string{
			1: "Janeiro", 2: "Fevereiro", 3: "Março", 4: "Abril", 5: "Maio", 6: "Junho",
			7: "Julho", 8: "Agosto", 9: "Setembro", 10: "Outubro", 11: "Novembro", 12: "Dezembro",
		}[mes]
		dashboard = append(dashboard, DashboardType{
			Month:    nomeMesPtBr,
			Entradas: totalEntradasMes,
			Saidas:   totalSaidasMes,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dashboard)
}
