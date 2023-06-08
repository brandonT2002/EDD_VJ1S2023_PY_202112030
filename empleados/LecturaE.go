package empleados

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/fatih/color"
)

func LeerCSV(lista *ListaEmp) {
	file, err := os.Open("ArchivosEntrada/empleados.csv")

	if err != nil {
		color.Red("\n  Error, no se pudo abrir el archivo")
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
			color.Yellow("\n  No se pudo leer la linea del csv")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		lista.Insertar(&Empleado{Id: linea[0], Nombre: linea[1], Cargo: linea[2], Contrasena: linea[3]})
	}
	color.Green("\n  Archivo cargado exitosamente")
}
