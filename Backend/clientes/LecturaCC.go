package clientes

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/fatih/color"
)

func LeerCSV1(cola *ColaCliente, lista *ListaCliente, archivo string) {
	file, err := os.Open("csv/" + archivo + ".csv")

	if err != nil {
		color.Red("\n  Error, no se pudo abrir el archivo\n\n")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			color.Yellow("\n  No se pudo leer la linea del csv\n\n")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		if linea[0] == "x" || linea[0] == "X" {
			cliente := &Cliente{Id: linea[0], Nombre: linea[1]}
			lista.Insertar(cliente)
			cola.Insertar(cliente)
		} else {
			cola.Insertar(&Cliente{Id: linea[0], Nombre: linea[1]})
		}
	}
	color.Green("\n  Archivo cargado exitosamente\n\n")
}