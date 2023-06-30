package imagencapas

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"paquete/empleados"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func LeerCSV(path, carpeta, nombre string) *MatrizDispersa {
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
		nombre:  nombre,
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
	return capa
}

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

func LeerConfig(carpeta, nombre string) []int {
	file, err := os.Open("./csv/" + carpeta + "/" + nombre)

	if err != nil {
		color.Red("\n  Error, no se pudo abrir el archivo\n\n")
		return nil
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true

	ordenValores := []string{"pixel_width", "image_width", "pixel_height", "image_height"}

	valores := make(map[string]int)

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

		nombre := linea[0]
		valor, err := strconv.Atoi(linea[1])
		if err != nil {
			color.Yellow("\n  No se pudo convertir el valor a entero\n\n")
			continue
		}
		valores[nombre] = valor
	}

	var config []int
	for _, nombre := range ordenValores {
		valor, ok := valores[nombre]
		if ok {
			config = append(config, valor)
		} else {
			color.Yellow("\n  No se encontró el valor para\n\n", nombre)
		}
	}

	return config
}

func CrearImg(img, cliente string, filtros *empleados.EnvioFiltros) string {
	array := leerArchivos(img)
	imagen := &ListaCapas{}
	nombre := ""
	// crear carpeta
	ruta := "../Filtros/" + img + "_" + cliente
	err := os.Mkdir(ruta, os.ModePerm)

	if err != nil {
		// La carpeta ya existe, intenta crear una con distinción
		i := 1
		for {
			nuevaRuta := fmt.Sprintf("%s_%d", ruta, i)
			err = os.Mkdir(nuevaRuta, os.ModePerm)
			if err == nil {
				ruta = nuevaRuta
				break
			}
			i++
		}
	}

	if err != nil {
		fmt.Println("Error al crear la carpeta:", err)
		return "Ocurrió un error al crear la imagen"
	}

	for _, fl := range array {
		nombre = strings.TrimSuffix(fl.Archivo, ".csv")
		if fl.Capa != 0 {
			imagen.Insertar(LeerCSV("./csv/"+img+"/"+fl.Archivo, img, nombre))
			continue
		}
		config := LeerConfig(img, fl.Archivo)
		imagen.GenerarImg(config[0], config[1], config[2], config[3], ruta, img, nil)
		imagen.GenerarImg(config[0], config[1], config[2], config[3], ruta, img, filtros)
	}
	return "Imagen Creada"
}
