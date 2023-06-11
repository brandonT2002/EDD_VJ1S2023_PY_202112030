package clientes

import "fmt"

type ListaCliente struct {
	primero  *Nodo
	ultimo   *Nodo
	longitud int
}

func (l *ListaCliente) Insertar(cliente *Cliente) {
	if l.primero != nil {
		l.ultimo.siguiente = &Nodo{cliente: cliente}
		l.ultimo = l.ultimo.siguiente
		l.ultimo.siguiente = l.primero
		l.longitud++
		return
	}
	l.primero = &Nodo{cliente: cliente}
	l.ultimo = l.primero
	l.ultimo.siguiente = l.primero
	l.longitud++
}

func (l *ListaCliente) Mostrar() {
	actual := l.primero
	contador := 0
	formato := "  ║ %-10s %-20s║\n"
	println("  ╔════════════════════════════════╗")
	fmt.Printf(formato, "ID", "Nombre")
	println("  ╠════════════════════════════════╣")
	for actual != nil {
		fmt.Printf(formato, actual.cliente.Id, actual.cliente.Nombre)
		contador++
		if contador == l.longitud {
			break
		}
		actual = actual.siguiente
	}
	println("  ╚════════════════════════════════╝")
}
