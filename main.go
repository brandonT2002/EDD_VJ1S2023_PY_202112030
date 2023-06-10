package main

import (
	"fmt"
	"paquete/consola"
	"paquete/empleados"
	"paquete/imagenes"
	paneladmin "paquete/panelAdmin"
	panelusuario "paquete/panelUsuario"
)

func Menu(lEmp *empleados.ListaEmp, lImg *imagenes.ListaImg) {
	opcion := 0
	for opcion != 2 {
		Opciones()
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			IniciarSesion(lEmp, lImg)
		case 2:
			fmt.Println()
			fmt.Println("  Hasta pronto")
		default:
			fmt.Println()
			fmt.Println("  Opción incorrecta")
		}
	}
}

func Opciones() {
	fmt.Println()
	fmt.Println("  ╔════════════════════════════════════════════════════╗")
	fmt.Println("  ║                                                    ║")
	fmt.Println("  ║                       LOGIN                        ║")
	fmt.Println("  ║                1. Iniciar Sesion                   ║")
	fmt.Println("  ║                2. Salir del Sistema                ║")
	fmt.Println("  ║                                                    ║")
	fmt.Println("  ╚════════════════════════════════════════════════════╝")
	fmt.Print("  Opción: ")
}

func IniciarSesion(lEmp *empleados.ListaEmp, lImg *imagenes.ListaImg) {
	var usuario string
	var pass string
	fmt.Println()
	fmt.Print("  -> 👨 Usuario: ")
	fmt.Scanln(&usuario)
	fmt.Print("  -> 🔒 Contraseña: ")
	fmt.Scanln(&pass)
	if usuario == "admin" && pass == "123" {
		consola.LimpiarConsola()
		paneladmin.MenuAdmin("PAKO", lEmp, lImg)
	} else {
		consola.LimpiarConsola()
		panelusuario.MenuUsuario("F")
	}
}

func main() {
	lEmp := &empleados.ListaEmp{}
	lImg := &imagenes.ListaImg{}

	consola.LimpiarConsola()
	Menu(lEmp, lImg)

	/*
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

		l3 := clientes.ListaCliente{}
		l3.Insertar(&clientes.Cliente{Id: "4435", Nombre: "Cristian Suy"})
		l3.Insertar(&clientes.Cliente{Id: "5268", Nombre: "Hector Jimenez"})
		l3.Insertar(&clientes.Cliente{Id: "2503", Nombre: "Pablo Velasquez"})
		l3.Insertar(&clientes.Cliente{Id: "4975", Nombre: "Jaquelin Gomez"})
		l3.Insertar(&clientes.Cliente{Id: "1280", Nombre: "Yadira Ruiz"})
		l3.Insertar(&clientes.Cliente{Id: "3101", Nombre: "Paula Fuentes"})
		l3.Mostrar()
	*/
}
