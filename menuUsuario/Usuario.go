package menuusuario

import "fmt"

func MenuUsuario(usuario string) {
	opcion := 0
	for opcion != 3 {
		opciones(usuario)
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Print("  Algo")
		case 2:
			fmt.Println("  Añgp")
		case 3:
			fmt.Println()
		default:
			fmt.Println()
			fmt.Println("  Opción incorrecta")
		}
	}
}

func opciones(usuario string) {
	usuario = "¡Bienvenido " + usuario + "!"
	anchoRecuadro := 55
	espacios := anchoRecuadro - len(usuario) - 2

	fmt.Println()
	fmt.Println("  ╔════════════════════════════════════════════════════╗")
	fmt.Println("  ║                                                    ║")
	fmt.Printf("  ║%*s%s%*s║\n", espacios/2, "", usuario, (espacios+1)/2, "")
	fmt.Println("  ║              1. Ver Imagenes cargadas              ║")
	fmt.Println("  ║              2. Realizar Pedido                    ║")
	fmt.Println("  ║              3. Cerrar Sesion                      ║")
	fmt.Println("  ║                                                    ║")
	fmt.Println("  ╚════════════════════════════════════════════════════╝")
	fmt.Print("  Opción: ")
}
