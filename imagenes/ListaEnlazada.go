package imagenes

import (
	"fmt"
	"strconv"
)

type ListaEnlazada struct {
	primero  *Nodo
	ultimo   *Nodo
	longitud int
}

func (nodo *ListaEnlazada) Insertar(pelicula *Imagen) {
	if nodo.primero != nil {
		nodo.ultimo.siguiente = &Nodo{imagen: pelicula}
		nodo.ultimo.siguiente.anterior = nodo.ultimo
		nodo.ultimo = nodo.ultimo.siguiente
		nodo.longitud++
		return
	}
	nodo.primero = &Nodo{imagen: pelicula}
	nodo.ultimo = nodo.primero
	nodo.longitud++
}

func (nodo *ListaEnlazada) Mostrar() {
	actual := nodo.primero
	formato := "%-15s %-10s\n"
	fmt.Printf(formato, "Imagen", "Capas")
	println("-----------------------")
	for actual != nil {
		fmt.Printf(formato, actual.imagen.Nombre, strconv.Itoa(actual.imagen.Capas))
		actual = actual.siguiente
	}
}
