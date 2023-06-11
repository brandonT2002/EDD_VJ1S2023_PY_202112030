package clientes

import (
	"encoding/csv"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func LeerCSV1(cola *ColaCliente) {
	file, err := os.Open("ArchivosEntrada/clientes_cola.csv")

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
		if linea[0] == "x" || linea[0] == "X" {
			id := generarID()
			cola.Insertar(&Cliente{Id: id, Nombre: linea[1]})
		}
	}
	color.Green("\n  Archivo cargado exitosamente")
}

func generarID() string {
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	numero := rng.Intn(9000) + 1000

	id := strconv.Itoa(numero)
	return id
}
