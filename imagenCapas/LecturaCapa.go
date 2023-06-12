package imagencapas

import (
	"encoding/csv"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
)

func LeerCSV(path, nombre string) *MatrizDispersa {
	file, err := os.Open(path)
	if err != nil {
		color.Red("\n  Error, no se pudo abrir el archivo")
		return nil
	}
	lectura := csv.NewReader(file)
	lectura.Comma = ','
	capa := &MatrizDispersa{
		accesoF: &ListaCabeza{},
		accesoC: &ListaCabeza{},
	}
	i := 0
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			color.Yellow("\n  No se pudo leer la linea del csv")
			continue
		}
		for j := 0; j < len(linea); j++ {
			if linea[j] != "x" {
				rgb := strings.Split(linea[j], "-")
				capa.insertar(i, j, &Color{R: rgb[0], G: rgb[1], B: rgb[2]})
			}
		}
		i++
	}
	generarGrafo(nombre, capa.dot())
	return capa
}
