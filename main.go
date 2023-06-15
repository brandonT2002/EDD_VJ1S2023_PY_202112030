package main

import (
	"fmt"
	"paquete/clientes"
	"paquete/consola"
	"paquete/empleados"
	"paquete/imagenes"
	paneladmin "paquete/panelAdmin"
	panelusuario "paquete/panelUsuario"
	"paquete/pedidos"
)

func Menu(lEmp *empleados.ListaEmp, lImg *imagenes.ListaImg, lCl *clientes.ListaCliente, cCl *clientes.ColaCliente, pCl *pedidos.Pila) {
	opcion := 0
	for opcion != 2 {
		Opciones()
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			IniciarSesion(lEmp, lImg, lCl, cCl, pCl)
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

func IniciarSesion(lEmp *empleados.ListaEmp, lImg *imagenes.ListaImg, lCl *clientes.ListaCliente, cCl *clientes.ColaCliente, pCl *pedidos.Pila) {
	emp := &empleados.Empleado{"3060", "pako", "Desarrollador", "123"}
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
		paneladmin.MenuAdmin(admin, lEmp, lImg, lCl, cCl, pCl)
	} else if usuario == emp.Nombre && pass == emp.Contrasena {
		consola.LimpiarConsola()
		panelusuario.MenuUsuario(emp, lImg, cCl, pCl, lCl)
	} else {
		fmt.Println("  \nVerifique sus credenciales")
	}
}

func main() {
	lEmp := &empleados.ListaEmp{}
	lImg := &imagenes.ListaImg{}
	lCl := &clientes.ListaCliente{}
	cCl := &clientes.ColaCliente{}
	pCl := &pedidos.Pila{}

	consola.LimpiarConsola()
	Menu(lEmp, lImg, lCl, cCl, pCl)

}
