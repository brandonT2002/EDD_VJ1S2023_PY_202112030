package main

import (
	"fmt"
	"paquete/clientes"
	"paquete/consola"
	"paquete/empleados"
	"paquete/imagenes"
	paneladmin "paquete/panelAdmin"
	panelusuario "paquete/panelUsuario"
)

func Menu(lEmp *empleados.ListaEmp, lImg *imagenes.ListaImg, lCl *clientes.ListaCliente, cCl *clientes.ColaCliente) {
	opcion := 0
	for opcion != 2 {
		Opciones()
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			IniciarSesion(lEmp, lImg, lCl, cCl)
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

func IniciarSesion(lEmp *empleados.ListaEmp, lImg *imagenes.ListaImg, lCl *clientes.ListaCliente, cCl *clientes.ColaCliente) {
	var usuario string
	var pass string
	admin := "br"
	passA := "1"
	fmt.Println()
	fmt.Print("  -> 👨 Usuario: ")
	fmt.Scanln(&usuario)
	fmt.Print("  -> 🔒 Contraseña: ")
	fmt.Scanln(&pass)
	if usuario == admin && pass == passA {
		consola.LimpiarConsola()
		paneladmin.MenuAdmin(admin, lEmp, lImg, lCl, cCl)
	} else {
		consola.LimpiarConsola()
		panelusuario.MenuUsuario("F", lImg)
	}
}

func main() {
	lEmp := &empleados.ListaEmp{}
	lImg := &imagenes.ListaImg{}
	lCl := &clientes.ListaCliente{}
	cCl := &clientes.ColaCliente{}

	consola.LimpiarConsola()
	Menu(lEmp, lImg, lCl, cCl)

}
