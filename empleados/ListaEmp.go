package empleados

import (
	"fmt"
)

type ListaEmp struct {
	primero  *Nodo
	ultimo   *Nodo
	longitud int
}

func (nodo *ListaEmp) Insertar(empleado *Empleado) {
	if nodo.primero != nil {
		nodo.ultimo.siguiente = &Nodo{empleado: empleado}
		nodo.ultimo = nodo.ultimo.siguiente
		nodo.longitud++
		return
	}
	nodo.primero = &Nodo{empleado: empleado}
	nodo.ultimo = nodo.primero
	nodo.longitud++
}

func (nodo *ListaEmp) Mostrar() {
	actual := nodo.primero
	formato := "%-10s %-25s %-15s %-10s\n"
	fmt.Printf(formato, "ID", "Nombre", "Cargo", "Contraseña")
	println("------------------------------------------------------------------")
	for actual != nil {
		fmt.Printf(formato, actual.empleado.Id, actual.empleado.Nombre, actual.empleado.Cargo, actual.empleado.Contrasena)
		actual = actual.siguiente
	}
	fmt.Println()
}
