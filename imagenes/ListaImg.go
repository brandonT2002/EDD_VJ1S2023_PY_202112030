package imagenes

import (
	"fmt"
)

type ListaImg struct {
	primero  *Nodo
	ultimo   *Nodo
	longitud int
}

func (nodo *ListaImg) Insertar(imagen *Imagen) {
	if nodo.primero != nil {
		nodo.ultimo.siguiente = &Nodo{imagen: imagen}
		nodo.ultimo.siguiente.anterior = nodo.ultimo
		nodo.ultimo = nodo.ultimo.siguiente
		nodo.longitud++
		return
	}
	nodo.primero = &Nodo{imagen: imagen}
	nodo.ultimo = nodo.primero
	nodo.longitud++
}

func (nodo *ListaImg) Mostrar() {
	actual := nodo.primero
	formato := "%-15s %-10s\n"
	fmt.Printf(formato, "Imagen", "Capas")
	println("-----------------------")
	for actual != nil {
		fmt.Printf(formato, actual.imagen.Nombre, actual.imagen.Capas)
		actual = actual.siguiente
	}
	fmt.Println()
}
