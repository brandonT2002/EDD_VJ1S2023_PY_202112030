package clientes

import "fmt"

type ColaCliente struct {
	primero  *Nodo
	ultimo   *Nodo
	longitud int
}

func (c *ColaCliente) Insertar(cliente *Cliente) {
	if c.primero != nil {
		c.ultimo.siguiente = &Nodo{cliente: cliente}
		c.ultimo = c.ultimo.siguiente
		c.longitud++
		return
	}
	c.primero = &Nodo{cliente: cliente}
	c.ultimo = c.primero
	c.longitud++
}

func (c *ColaCliente) Mostrar() {
	actual := c.primero
	formato := "  ║ %-10s %-20s║\n"
	println("  ╔════════════════════════════════╗")
	fmt.Printf(formato, "ID", "Nombre")
	println("  ╠════════════════════════════════╣")
	for actual != nil {
		fmt.Printf(formato, actual.cliente.Id, actual.cliente.Nombre)
		actual = actual.siguiente
	}
	println("  ╚════════════════════════════════╝")
}

/*
func (l *Cola) eliminar() interface{} {
	if l.primero != nil {
		// temporal := l.primero.cliente
		l.primero = l.primero.siguiente
		l.longitud--
	}
	return nil
}
*/
