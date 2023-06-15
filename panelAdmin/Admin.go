package paneladmin

import (
	"fmt"
	"io/ioutil"
	"paquete/clientes"
	"paquete/consola"
	"paquete/empleados"
	"paquete/imagenes"
	"paquete/pedidos"

	"github.com/fatih/color"
)

func MenuAdmin(admin string, lEmp *empleados.ListaEmp, lImg *imagenes.ListaImg, lCl *clientes.ListaCliente, cCl *clientes.ColaCliente, pCl *pedidos.Pila) {
	opcion := 0
	archivo := ""
	for opcion != 6 {
		opciones(admin)
		fmt.Scanln(&opcion)

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
			reportes(lImg, lEmp, lCl, cCl, pCl)
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
	consola.LimpiarConsola()
	empleados.LeerCSV(lEmp, archivo)
	fmt.Println("  -- EMPLEADOS --")
	lEmp.Mostrar()
}

func leerImagenes(lImg *imagenes.ListaImg, archivo string) {
	consola.LimpiarConsola()
	imagenes.LeerCSV(lImg, archivo)
	fmt.Println("  -- IMAGENES --")
	lImg.Mostrar()
}

func leerClientes(lCl *clientes.ListaCliente, archivo string) {
	consola.LimpiarConsola()
	clientes.LeerCSV(lCl, archivo)
	fmt.Println("  -- CLIENTES --")
	lCl.Mostrar()
}

func leerClientesCola(cCl *clientes.ColaCliente, lCl *clientes.ListaCliente, archivo string) {
	consola.LimpiarConsola()
	clientes.LeerCSV1(cCl, lCl, archivo)
	fmt.Println("  -- CLIENTES EN COLA --")
	cCl.Mostrar()
	fmt.Println("  -- CLIENTES EN EL SISTEMA --")
	lCl.Mostrar()
}

func reportes(lImg *imagenes.ListaImg, lEmp *empleados.ListaEmp, lCl *clientes.ListaCliente, cCl *clientes.ColaCliente, pCl *pedidos.Pila) {
	consola.LimpiarConsola()
	opcion := 0
	for opcion != 7 {
		_reportes()
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			lImg.Reporte()
			consola.LimpiarConsola()
		case 2:
			lEmp.Reporte()
			consola.LimpiarConsola()
		case 3:
			lCl.Reporte()
			consola.LimpiarConsola()
		case 4:
			cCl.Reporte()
			consola.LimpiarConsola()
		case 5:
			pCl.Reporte()
			consola.LimpiarConsola()
		case 6:
			fmt.Println()
			consola.LimpiarConsola()
		case 7:
			jsonD, err := pCl.Json()
			if err != nil {
				color.Red("\n  Error al convertir el archivo JSON\n\n")
				return
			}

			err = ioutil.WriteFile("./Reportes/ReporteJSON.json", jsonD, 0644)
			if err != nil {
				color.Red("\n  Error al escribir en el archivo\n\n")
				return
			}

			consola.LimpiarConsola()
		case 8:
			consola.LimpiarConsola()
			return
		default:
			fmt.Println()
			fmt.Println("  Opción incorrecta")
		}
	}
}

func _reportes() {
	fmt.Println()
	fmt.Println("  ╔════════════════════════════════════════════════════╗")
	fmt.Println("  ║                                                    ║")
	fmt.Println("  ║                  1. Lista Doble                    ║")
	fmt.Println("  ║                  2. Lista Simple                   ║")
	fmt.Println("  ║                  3. Lista Circular                 ║")
	fmt.Println("  ║                  4. Cola                           ║")
	fmt.Println("  ║                  5. Pila                           ║")
	fmt.Println("  ║                  6. Matriz Dispersa                ║")
	fmt.Println("  ║                  7. Reporte JSON                   ║")
	fmt.Println("  ║                  8. Regresar                       ║")
	fmt.Println("  ║                                                    ║")
	fmt.Println("  ╚════════════════════════════════════════════════════╝")
	fmt.Print("  Opción: ")
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
