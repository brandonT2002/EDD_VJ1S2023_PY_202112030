package main

import "paquete/imagenes"

func main() {
	l1 := imagenes.ListaEnlazada{}
	l1.Insertar(&imagenes.Imagen{Nombre: "bmo", Capas: 3})
	l1.Insertar(&imagenes.Imagen{Nombre: "corrin", Capas: 3})
	l1.Insertar(&imagenes.Imagen{Nombre: "deadpool", Capas: 4})
	l1.Insertar(&imagenes.Imagen{Nombre: "hulk", Capas: 4})
	l1.Insertar(&imagenes.Imagen{Nombre: "mario", Capas: 6})
	l1.Insertar(&imagenes.Imagen{Nombre: "tux", Capas: 3})
	l1.Mostrar()
}
