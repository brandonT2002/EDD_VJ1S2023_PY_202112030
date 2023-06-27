package main

import (
	"fmt"
	"paquete/empleados"
	"paquete/pedidos"
	"strconv"

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

type PedidoJSON struct {
	IdCliente string `json:"IdCliente"`
	Imagen    string `json:"Imagen"`
}

var admin = "123"
var passA = "123"
var Emp *empleados.Empleado
var LEmp *empleados.ListaEmp
var Arbol *pedidos.ArbolAVL
var Cola *pedidos.ColaPedidos

func main() {
	LEmp = &empleados.ListaEmp{}
	Arbol = &pedidos.ArbolAVL{}
	Cola = &pedidos.ColaPedidos{}

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"msg": "api en go",
		})
	})

	app.Post("/login", Login)
	app.Post("/pedidos", cargarPedidos)
	app.Get("/pedidos", verPedidos)
	app.Post("/empleado", cargarEmpleados)
	app.Post("/ventas", registrarVentas)
	app.Get("/ventas", solicitudes)

	app.Listen(":8080")
}

func cargarPedidos(c *fiber.Ctx) error {
	type PedidoRequest struct {
		Pedidos []PedidoJSON `json:"Pedidos"`
	}

	var request PedidoRequest
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	for _, pedido := range request.Pedidos {
		id, _ := strconv.Atoi(pedido.IdCliente)
		Arbol.Insertar(&pedidos.Pedido{IdCliente: id, Imagen: pedido.Imagen})
	}
	Arbol.Inorder(Cola)
	Cola.Mostrar()
	return c.JSON(&fiber.Map{
		"msg": "ok",
	})
}

func verPedidos(c *fiber.Ctx) error {
	return c.JSON(Cola.Cjson())
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

func registrarVentas(c *fiber.Ctx) error {
	var nuevoNodo empleados.EnvioMatriz
	c.BodyParser(&nuevoNodo)
	Emp.Grafo.InsertarValores(&nuevoNodo)
	Cola.Eliminar()
	fmt.Println(Emp.Grafo.Dot())
	return c.JSON(&fiber.Map{
		"msg": "Venta Registrada",
	})
}

func solicitudes(c *fiber.Ctx) error {
	return c.JSON(Emp.Grafo.Mjson())
}

func Login(c *fiber.Ctx) error {
	var usuario Usuario
	c.BodyParser(&usuario)
	Emp = LEmp.Buscar(usuario.Usuario, usuario.Contrasena)

	if usuario.Usuario == admin && usuario.Contrasena == passA {
		return c.JSON(&fiber.Map{
			"msg":     "admin",
			"usuario": admin,
		})
	} else if Emp != nil {
		return c.JSON(&fiber.Map{
			"msg":     "emp",
			"usuario": Emp,
		})
	}
	return c.JSON(&fiber.Map{
		"msg": "no",
	})
}
