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
		emp := &Empleado{Id: linea[0], Nombre: linea[1], Cargo: linea[2], Contrasena: linea[3]}
		lista.Insertar(emp)
	}
	return "ok"
}
