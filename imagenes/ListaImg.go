package imagenes

import (
	"fmt"
)

type ListaImg struct {
	primero  *Nodo
	ultimo   *Nodo
	longitud int
}

func (l *ListaImg) Insertar(imagen *Imagen) {
	if l.primero != nil {
		l.ultimo.siguiente = &Nodo{imagen: imagen}
		l.ultimo.siguiente.anterior = l.ultimo
		l.ultimo = l.ultimo.siguiente
		l.longitud++
		return
	}
	l.primero = &Nodo{imagen: imagen}
	l.ultimo = l.primero
	l.longitud++
}

func (l *ListaImg) Buscar(imagen string) bool {
	actual := l.primero
	for actual != nil {
		if actual.imagen.Nombre == imagen {
			return true
		}
		actual = actual.siguiente
	}
	return false
}

func (l *ListaImg) Mostrar() {
	actual := l.primero
	formato := "  ║ %-17s %-10s║\n"
	println("  ╔═════════════════════════════╗")
	fmt.Printf(formato, "Imagen", "Capas")
	println("  ╠═════════════════════════════╣")
	for actual != nil {
		fmt.Printf(formato, actual.imagen.Nombre, actual.imagen.Capas)
		actual = actual.siguiente
	}
	println("  ╚═════════════════════════════╝")
}
