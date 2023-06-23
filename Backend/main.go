package main

import (
	"fmt"
	"paquete/empleados"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Usuario struct {
	Usuario    string
	Contrasena string
}

type Empleado struct {
	Credenciales string
}

var admin = "123"
var passA = "123"
var LEmp *empleados.ListaEmp

func main() {
	LEmp = &empleados.ListaEmp{}

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"msg": "api en go",
		})
	})

	app.Post("/login", Login)
	app.Post("/empleado", cargarEmpleados)

	app.Listen(":8080")
}

func cargarEmpleados(c *fiber.Ctx) error {
	var emp Empleado
	err := c.BodyParser(&emp)
	if err != nil {
		return err
	}
	fmt.Println(emp)
	resultado := empleados.LeerCSV(LEmp, emp.Credenciales)
	if resultado != "ok" {
		return c.JSON(&fiber.Map{
			"msg": "no",
		})
	}
	LEmp.Mostrar()
	return c.JSON(&fiber.Map{
		"msg": "ok",
	})

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
