package clientes

import (
	"fmt"
	"strconv"
)

type ListaCliente struct {
	primero  *Nodo
	ultimo   *Nodo
	longitud int
}

func (l *ListaCliente) Insertar(cliente *Cliente) {
	if l.existe(cliente.Id) {
		return
	}
	if l.primero != nil {
		l.ultimo.siguiente = &Nodo{Cliente: cliente}
		l.ultimo = l.ultimo.siguiente
		l.ultimo.siguiente = l.primero
		l.longitud++
		return
	}
	l.primero = &Nodo{Cliente: cliente}
	l.ultimo = l.primero
	l.ultimo.siguiente = l.primero
	l.longitud++
}

func (l *ListaCliente) existe(id string) bool {
	actual := l.primero
	contador := 0
	for actual != nil {
		if contador == l.longitud {
			break
		}
		if actual.Cliente.Id == id && l.esNumero(id) {
			return true
		}
		actual = actual.siguiente
		contador++
	}
	return false
}

func (l *ListaCliente) esNumero(id string) bool {
	_, err := strconv.Atoi(id)
	if err == nil {
		return true
	}
	return false
}

func (l *ListaCliente) Mostrar() {
	actual := l.primero
	contador := 0
	formato := "  ║ %-10s %-20s║\n"
	println("  ╔════════════════════════════════╗")
	fmt.Printf(formato, "ID", "Nombre")
	println("  ╠════════════════════════════════╣")
	for actual != nil {
		fmt.Printf(formato, actual.Cliente.Id, actual.Cliente.Nombre)
		contador++
		if contador == l.longitud {
			break
		}
		actual = actual.siguiente
	}
	println("  ╚════════════════════════════════╝")
}
