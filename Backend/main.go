package main

import (
	"paquete/empleados"
	"paquete/pedidos"

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

type Pedido struct {
	Pedidos string
}

var admin = "123"
var passA = "123"
var LEmp *empleados.ListaEmp
var Arbol *pedidos.ArbolAVL

func main() {
	LEmp = &empleados.ListaEmp{}
	Arbol = &pedidos.ArbolAVL{}

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"msg": "api en go",
		})
	})

	app.Post("/login", Login)
	app.Post("/pedidos", cargarPedidos)
	app.Post("/empleado", cargarEmpleados)

	app.Listen(":8080")
}

func cargarPedidos(c *fiber.Ctx) error {
	var pedido Pedido
	err := c.BodyParser(&pedido)
	if err != nil {
		return err
	}

	err = pedidos.LeerJSON(Arbol, pedido.Pedidos)
	if err != nil {
		return c.JSON(&fiber.Map{
			"msg": "no",
		})
	}
	return c.JSON(&fiber.Map{
		"msg": "ok",
	})
}

func cargarEmpleados(c *fiber.Ctx) error {
	var emp Empleado
	err := c.BodyParser(&emp)
	if err != nil {
		return err
	}

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
	emp := LEmp.Buscar(usuario.Usuario, usuario.Contrasena)

	if usuario.Usuario == admin && usuario.Contrasena == passA {
		return c.JSON(&fiber.Map{
			"msg": "admin",
		})
	} else if emp != nil {
		return c.JSON(&fiber.Map{
			"msg": "emp",
		})
	}
	return c.JSON(&fiber.Map{
		"msg": "no",
	})
}
