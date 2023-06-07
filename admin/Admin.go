package admin

import (
	"fmt"
)

func MenuAdmin() {
	opcion := 0
	for opcion != 6 {
		opciones()
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Print("  Algo")
		case 2:
			fmt.Println()
			fmt.Println("  Hasta pronto")
		case 3:
			fmt.Println()
			fmt.Println("  Hasta pronto")
		case 4:
			fmt.Println()
			fmt.Println("  Hasta pronto")
		case 5:
			fmt.Println()
			fmt.Println("  Hasta pronto")
		case 6:
			fmt.Println()
		default:
			fmt.Println()
			fmt.Println("  Opción incorrecta")
		}
	}
}

func opciones() {
	fmt.Println()
	fmt.Println("  ╔═══════════════════════════════════════════════════╗")
	fmt.Println("  ║                                                   ║")
	fmt.Println("  ║              ADMINISTRADOR 202112030              ║")
	fmt.Println("  ║                1. Cargar Empleados                ║")
	fmt.Println("  ║                2. Cargar Imagenes                 ║")
	fmt.Println("  ║                3. Cargar Usuarios                 ║")
	fmt.Println("  ║                4. Actualizar Cola                 ║")
	fmt.Println("  ║                5. Reportes Estructuras            ║")
	fmt.Println("  ║                6. Cerrar Sesion                   ║")
	fmt.Println("  ║                                                   ║")
	fmt.Println("  ╚═══════════════════════════════════════════════════╝")
	fmt.Print("  Opción: ")
}
