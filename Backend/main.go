package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Usuario struct {
	Usuario    string
	Contrasena string
}

var admin = "123"
var passA = "123"

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"msg": "api en go",
		})
	})

	app.Post("/login", Login)

	app.Listen(":8080")
}

func Login(c *fiber.Ctx) error {
	var usuario Usuario
	c.BodyParser(&usuario)
	if usuario.Usuario == admin && usuario.Contrasena == passA {
		return c.JSON(&fiber.Map{
			"msg": "ok",
		})
	}
	return c.JSON(&fiber.Map{
		"msg": "no",
	})
}
