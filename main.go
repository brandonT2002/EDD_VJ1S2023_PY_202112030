package main

import (
	"paquete/empleados"
	"paquete/imagenes"
)

func main() {
	l1 := imagenes.ListaImg{}
	l1.Insertar(&imagenes.Imagen{Nombre: "bmo", Capas: 3})
	l1.Insertar(&imagenes.Imagen{Nombre: "corrin", Capas: 3})
	l1.Insertar(&imagenes.Imagen{Nombre: "deadpool", Capas: 4})
	l1.Insertar(&imagenes.Imagen{Nombre: "hulk", Capas: 4})
	l1.Insertar(&imagenes.Imagen{Nombre: "mario", Capas: 6})
	l1.Insertar(&imagenes.Imagen{Nombre: "tux", Capas: 3})
	l1.Mostrar()

	l2 := empleados.ListaEmp{}
	l2.Insertar(&empleados.Empleado{Id: "6385", Nombre: "Cristian Suy", Cargo: "Ventas", Contrasena: "6385_Ventas"})
	l2.Insertar(&empleados.Empleado{Id: "5372", Nombre: "Hector Jimenez", Cargo: "Ventas", Contrasena: "5372_Ventas"})
	l2.Insertar(&empleados.Empleado{Id: "4399", Nombre: "Pablo Velasquez", Cargo: "Diseño", Contrasena: "4399_Diseño"})
	l2.Insertar(&empleados.Empleado{Id: "1229", Nombre: "Jaquelin Gomez", Cargo: "Diseño", Contrasena: "1229_Diseño"})
	l2.Insertar(&empleados.Empleado{Id: "3607", Nombre: "Yadira Ruiz", Cargo: "Diseño", Contrasena: "3607_Diseño"})
	l2.Insertar(&empleados.Empleado{Id: "3518", Nombre: "Paula Fuentes", Cargo: "Ventas", Contrasena: "3518_Ventas"})
	l2.Mostrar()
}
