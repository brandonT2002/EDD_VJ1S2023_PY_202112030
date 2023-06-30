package pedidos

import (
	"fmt"
	"strconv"
)

type ColaPedidos struct {
	Primero  *NodoC
	ultimo   *NodoC
	longitud int
}

func (c *ColaPedidos) Insertar(cliente *Pedido) {
	if c.Primero != nil {
		c.ultimo.siguiente = &NodoC{Pedido: cliente}
		c.ultimo = c.ultimo.siguiente
		c.longitud++
		return
	}
	c.Primero = &NodoC{Pedido: cliente}
	c.ultimo = c.Primero
	c.longitud++
}

func (c *ColaPedidos) Eliminar() *Pedido {
	if c.Primero != nil {
		temporal := c.Primero.Pedido
		c.Primero = c.Primero.siguiente
		c.longitud--
		return temporal
	}
	return nil
}

func (c *ColaPedidos) Cjson() []Pedido {
	actual := c.Primero
	var pedidos []Pedido
	for actual != nil {
		pedidos = append(pedidos, *actual.Pedido)
		actual = actual.siguiente
	}
	return pedidos
}

func (c *ColaPedidos) Mostrar() {
	actual := c.Primero
	formato := "  ║ %-10s %-20s║\n"
	println("  ╔════════════════════════════════╗")
	fmt.Printf(formato, "ID", "Nombre")
	println("  ╠════════════════════════════════╣")
	for actual != nil {
		fmt.Printf(formato, strconv.Itoa(actual.Pedido.IdCliente), actual.Pedido.Imagen)
		actual = actual.siguiente
	}
	println("  ╚════════════════════════════════╝")
}
