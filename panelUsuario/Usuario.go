package panelusuario

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"paquete/clientes"
	"paquete/consola"
	"paquete/empleados"
	imagencapas "paquete/imagenCapas"
	"paquete/imagenes"
	"paquete/pedidos"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type ArchivoCapa struct {
	Capa    int
	Archivo string
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

	// sort.Slice(archivoCapa, func(i, j int) bool {
	// 	return archivoCapa[i].Capa < archivoCapa[j].Capa
	// })

	return archivoCapa
}

func leerConfig(carpeta, nombre string) []int {
	file, err := os.Open("./csv/" + carpeta + "/" + nombre)

	if err != nil {
		color.Red("\n  Error, no se pudo abrir el archivo")
		return nil
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true

	var config []int
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			color.Yellow("\n  No se pudo leer la linea del csv")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		// fmt.Println("  -> ", linea[1])
		valor, _ := strconv.Atoi(linea[1])
		config = append(config, valor)
	}
	return config
}

func generarImagen(img string) {
	array := leerArchivos(img)
	imagen := &imagencapas.ListaCapas{}
	nombre := ""
	// aquí se crea carpeta
	ruta := "./" + img
	err := os.Mkdir(ruta, os.ModePerm)
	if err != nil {
		fmt.Println("Error al crear la carpeta:", err)
		return
	}

	for _, fl := range array {
		nombre = strings.TrimSuffix(fl.Archivo, ".csv")
		if fl.Archivo != "config.csv" {
			imagen.Insertar(imagencapas.LeerCSV("./csv/"+img+"/"+fl.Archivo, img, nombre))
			continue
		}
		config := leerConfig(img, fl.Archivo)
		imagen.GenerarImg(config[0], config[1], config[2], config[3], ruta, img)
	}
}

func MenuUsuario(usuario *empleados.Empleado, lImg *imagenes.ListaImg, cCl *clientes.ColaCliente, pCl *pedidos.Pila) {
	opcion := 0
	img := ""

	for opcion != 3 {
		opciones(usuario.Nombre)
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			lImg.Mostrar()
			fmt.Print("\n  Imagen: ")
			fmt.Scanln(&img)
			generarImagen(img)
		case 2:
			pedido(usuario.Id, cCl, pCl, lImg)
		case 3:
			fmt.Println()
			consola.LimpiarConsola()
		default:
			fmt.Println()
			fmt.Println("  Opción incorrecta")
		}
	}
}

func pedido(idEmp string, cCl *clientes.ColaCliente, pCl *pedidos.Pila, lImg *imagenes.ListaImg) {
	idCliente := ""
	imagen := ""

	fmt.Println("\n  -- CLIENTES EN COLA --")
	cCl.Mostrar()

	if cCl.Primero.Cliente.Id == "X" || cCl.Primero.Cliente.Id == "x" {
		fmt.Println("algo")
	} else {
		for {
			fmt.Print("\n  ID Cliente: ")
			fmt.Scanln(&idCliente)
			if cCl.Primero.Cliente.Id == idCliente {
				break
			}
			color.Yellow("  Verifique el ID")
		}
	}
	fmt.Printf("\n  ID Empleado: %s\n", idEmp)
	fmt.Println("\n  -- IMAGENES EXISTENTES --")
	lImg.Mostrar()
	for {
		fmt.Print("\n  Imagen: ")
		fmt.Scanln(&imagen)
		if lImg.Buscar(imagen) {
			break
		}
		color.Yellow("  Verifique el nombre de la imagen")
	}
	pCl.Insertar(&pedidos.Pedido{IdCliente: idCliente, IdEmpleado: idEmp, Imagen: imagen})
	cCl.Eliminar()
	fmt.Println("\n  -- PEDIDOS --")
	pCl.Mostrar()

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
