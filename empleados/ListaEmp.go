package empleados

import (
	"fmt"
)

type ListaEmp struct {
	primero  *Nodo
	ultimo   *Nodo
	longitud int
}

func (l *ListaEmp) Insertar(empleado *Empleado) {
	if l.primero != nil {
		l.ultimo.siguiente = &Nodo{empleado: empleado}
		l.ultimo = l.ultimo.siguiente
		l.longitud++
		return
	}
	l.primero = &Nodo{empleado: empleado}
	l.ultimo = l.primero
	l.longitud++
}

func (l *ListaEmp) Mostrar() {
	actual := l.primero
	formato := "  ║ %-10s %-25s %-15s %-15s║\n"
	println("  ╔═════════════════════════════════════════════════════════════════════╗")
	fmt.Printf(formato, "ID", "Nombre", "Cargo", "Contraseña")
	println("  ╠═════════════════════════════════════════════════════════════════════╣")
	for actual != nil {
		fmt.Printf(formato, actual.empleado.Id, actual.empleado.Nombre, actual.empleado.Cargo, actual.empleado.Contrasena)
		actual = actual.siguiente
	}
	println("  ╚═════════════════════════════════════════════════════════════════════╝")
}
