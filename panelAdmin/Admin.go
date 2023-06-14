package paneladmin

import (
	"fmt"
	"paquete/clientes"
	"paquete/consola"
	"paquete/empleados"
	"paquete/imagenes"
)

func MenuAdmin(admin string, lEmp *empleados.ListaEmp, lImg *imagenes.ListaImg, lCl *clientes.ListaCliente, cCl *clientes.ColaCliente) {
	opcion := 0
	for opcion != 6 {
		opciones(admin)
		fmt.Scanln(&opcion)
		archivo := ""

		switch opcion {
		case 1:
			fmt.Print("\n  Nombre del archivo: ")
			fmt.Scanln(&archivo)
			leerEmpleados(lEmp, archivo)
		case 2:
			fmt.Print("\n  Nombre del archivo: ")
			fmt.Scanln(&archivo)
			leerImagenes(lImg, archivo)
		case 3:
			fmt.Print("\n  Nombre del archivo: ")
			fmt.Scanln(&archivo)
			leerClientes(lCl, archivo)
		case 4:
			fmt.Print("\n  Nombre del archivo: ")
			fmt.Scanln(&archivo)
			leerClientesCola(cCl, lCl, archivo)
		case 5:
			fmt.Println()
			fmt.Println("  Hasta pronto")
		case 6:
			fmt.Println()
			consola.LimpiarConsola()
		default:
			fmt.Println()
			fmt.Println("  Opción incorrecta")
		}
	}
}

func leerEmpleados(lEmp *empleados.ListaEmp, archivo string) {
	empleados.LeerCSV(lEmp, archivo)
	fmt.Println("  -- EMPLEADOS --")
	lEmp.Mostrar()
}

func leerImagenes(lImg *imagenes.ListaImg, archivo string) {
	imagenes.LeerCSV(lImg, archivo)
	fmt.Println("  -- IMAGENES --")
	lImg.Mostrar()
}

func leerClientes(lCl *clientes.ListaCliente, archivo string) {
	clientes.LeerCSV(lCl, archivo)
	fmt.Println("  -- CLIENTES --")
	lCl.Mostrar()
}

func leerClientesCola(cCl *clientes.ColaCliente, lCl *clientes.ListaCliente, archivo string) {
	clientes.LeerCSV1(cCl, lCl, archivo)
	fmt.Println("  -- CLIENTES EN COLA --")
	cCl.Mostrar()
	fmt.Println("  -- CLIENTES EN EL SISTEMA --")
	lCl.Mostrar()
}

func opciones(admin string) {
	admin = "¡Bienvenido " + admin + "!"
	anchoRecuadro := 55
	espacios := anchoRecuadro - len(admin) - 2

	fmt.Println()
	fmt.Println("  ╔════════════════════════════════════════════════════╗")
	fmt.Println("  ║                                                    ║")
	fmt.Printf("  ║%*s%s%*s║\n", espacios/2, "", admin, (espacios+1)/2, "")
	fmt.Println("  ║                1. Cargar Empleados                 ║")
	fmt.Println("  ║                2. Cargar Imagenes                  ║")
	fmt.Println("  ║                3. Cargar Usuarios                  ║")
	fmt.Println("  ║                4. Actualizar Cola                  ║")
	fmt.Println("  ║                5. Reportes Estructuras             ║")
	fmt.Println("  ║                6. Cerrar Sesion                    ║")
	fmt.Println("  ║                                                    ║")
	fmt.Println("  ╚════════════════════════════════════════════════════╝")
	fmt.Print("  Opción: ")
}
