package empleados

import (
	"encoding/csv"
	"io"
	"strings"

	"github.com/fatih/color"
)

func LeerCSV(lista *ListaEmp, contenido string) string {
	lectura := csv.NewReader(strings.NewReader(contenido))
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			return "ok"
		}
		if err != nil {
			color.Yellow("\n  No se pudo leer la linea del csv\n\n")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		if strings.ReplaceAll(linea[0], " ", "") != "" || strings.ReplaceAll(linea[1], " ", "") != "" || strings.ReplaceAll(linea[2], " ", "") != "" || strings.ReplaceAll(linea[3], " ", "") != "" {
			lista.Insertar(&Empleado{Id: linea[0], Nombre: linea[1], Cargo: linea[2], Contrasena: linea[3], Grafo: &Grafo{nil}})
		}
	}
}
