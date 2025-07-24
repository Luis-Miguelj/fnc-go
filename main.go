package main

import (
	"backend-go/controller"
	"backend-go/db"
	"backend-go/infra"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db.Connect()
	db.DB.AutoMigrate(&infra.User{})

	app.Post("/register", controller.RegisterUser)
	app.Post("/login", controller.Login)
	app.Get("/user", controller.GetUser)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
