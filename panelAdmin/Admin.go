package paneladmin

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"paquete/clientes"
	"paquete/consola"
	"paquete/empleados"
	imagencapas "paquete/imagenCapas"
	"paquete/imagenes"
	"paquete/pedidos"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type ArchivoCapa struct {
	Capa    int
	Archivo string
}

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
			consola.LimpiarConsola()
			lImg.Mostrar()
			img := ""
			fmt.Print("  Imagen: ")
			fmt.Scanln(&img)
			numero := 0
			fmt.Print("  No. Capa: ")
			fmt.Scanln(&numero)
			reporte(img, numero)
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

func leerArchivos(nombre string) []ArchivoCapa {
	ruta := "./csv/" + nombre + "/inicial.csv"
	file, err := os.Open(ruta)
	if err != nil {
		color.Red("\n  Error, no se pudo abrir el archivo")
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read()
	if err != nil {
		color.Red("\n  Error, no se pudieron leer los encabezados: ")
		return nil
	}

	var archivoCapa []ArchivoCapa
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			color.Red("\n  Error, no se pudo leer el archivo: ")
			return nil
		}

		layer, err := strconv.Atoi(row[0])
		if err != nil {
			color.Red("\n  Error, no se pudo convertir el número de capa: ")
			return nil
		}

		archivoCapa = append(archivoCapa, ArchivoCapa{
			Capa:    layer,
			Archivo: row[1],
		})
	}

	sort.Slice(archivoCapa, func(i, j int) bool {
		if archivoCapa[i].Capa == 0 {
			return false
		} else if archivoCapa[j].Capa == 0 {
			return true
		}
		return archivoCapa[i].Capa < archivoCapa[j].Capa
	})

	return archivoCapa
}

func reporte(img string, capa int) {
	array := leerArchivos(img)
	nombre := ""

	for _, fl := range array {
		nombre = strings.TrimSuffix(fl.Archivo, ".csv")
		if capa == 0 {
			color.Yellow("\n  La capa 0 está reservado para las configuraciones")
			break
		}
		if fl.Archivo != "config.csv" {
			if fl.Capa == capa {
				imagencapas.LeerCSV("./csv/"+img+"/"+fl.Archivo, img, nombre).GenerarGrafo("Reportes")
				break
			}
		}
	}
	color.Yellow("\n  La capa no existe")
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
