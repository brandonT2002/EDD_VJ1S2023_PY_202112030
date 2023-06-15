package clientes

import "fmt"

type ColaCliente struct {
	Primero  *Nodo
	ultimo   *Nodo
	longitud int
}

func (c *ColaCliente) Insertar(cliente *Cliente) {
	if c.Primero != nil {
		c.ultimo.siguiente = &Nodo{Cliente: cliente}
		c.ultimo = c.ultimo.siguiente
		c.longitud++
		return
	}
	c.Primero = &Nodo{Cliente: cliente}
	c.ultimo = c.Primero
	c.longitud++
}

func (c *ColaCliente) Eliminar() interface{} {
	if c.Primero != nil {
		temporal := c.Primero.Cliente
		c.Primero = c.Primero.siguiente
		c.longitud--
		return temporal
	}
	return nil
}

func (c *ColaCliente) Mostrar() {
	actual := c.Primero
	formato := "  ║ %-10s %-20s║\n"
	println("  ╔════════════════════════════════╗")
	fmt.Printf(formato, "ID", "Nombre")
	println("  ╠════════════════════════════════╣")
	for actual != nil {
		fmt.Printf(formato, actual.Cliente.Id, actual.Cliente.Nombre)
		actual = actual.siguiente
	}
	println("  ╚════════════════════════════════╝")
}
