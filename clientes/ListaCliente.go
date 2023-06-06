package clientes

import "fmt"

type ListaCliente struct {
	primero  *Nodo
	ultimo   *Nodo
	longitud int
}

func (nodo *ListaCliente) Insertar(cliente *Cliente) {
	if nodo.primero != nil {
		nodo.ultimo.siguiente = &Nodo{cliente: cliente}
		nodo.ultimo = nodo.ultimo.siguiente
		nodo.ultimo.siguiente = nodo.primero
		nodo.longitud++
		return
	}
	nodo.primero = &Nodo{cliente: cliente}
	nodo.ultimo = nodo.primero
	nodo.ultimo.siguiente = nodo.primero
	nodo.longitud++
}

func (nodo *ListaCliente) Mostrar() {
	actual := nodo.primero
	contador := 0
	formato := "%-10s %-25s\n"
	fmt.Printf(formato, "ID", "Nombre")
	println("------------------------------")
	for actual != nil {
		fmt.Printf(formato, actual.cliente.Id, actual.cliente.Nombre)
		contador++
		if contador == nodo.longitud {
			break
		}
		actual = actual.siguiente
	}
	fmt.Println()
}
