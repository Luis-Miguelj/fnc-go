package main

import (
	"backend-go/controller"
	"backend-go/db"
	"backend-go/infra"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Meu Projeto Go com Swagger
// @version 1.0
// @description Esta é a documentação da minha API em Go.
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()

	db.Connect()
	db.DB.AutoMigrate(&infra.User{})
	db.DB.AutoMigrate(&infra.Finance{})
	db.DB.AutoMigrate(&infra.Category{})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Post("/finance/:userId", controller.CreateFinance)
	app.Post("/register", controller.RegisterUser)
	app.Post("/login", controller.Login)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	app.Get("/user", controller.GetUser)
	app.Get("/finances/:userId", controller.GetFinances)
	app.Get("/categories", controller.GetCategories)

	app.Put("/category/:id", controller.UpdateCategory)

	app.Delete("/category/:id", controller.DeleteCategory)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
