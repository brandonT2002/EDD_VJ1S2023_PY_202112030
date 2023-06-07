package admin

import (
	"fmt"
)

func MenuAdmin(admin string) {
	opcion := 0
	for opcion != 6 {
		opciones(admin)
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
