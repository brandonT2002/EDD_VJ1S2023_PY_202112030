package clientes

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

type ColaCliente struct {
	Primero  *Nodo
	ultimo   *Nodo
	longitud int
}

func (c *ColaCliente) Insertar(cliente *Cliente) {
	if c.Primero != nil {
		c.ultimo.siguiente = &Nodo{Cliente: cliente}
		c.ultimo = c.ultimo.siguiente
		c.longitud++
		return
	}
	c.Primero = &Nodo{Cliente: cliente}
	c.ultimo = c.Primero
	c.longitud++
}

func (c *ColaCliente) Eliminar() interface{} {
	if c.Primero != nil {
		temporal := c.Primero.Cliente
		c.Primero = c.Primero.siguiente
		c.longitud--
		return temporal
	}
	return nil
}

func (nodo *ColaCliente) Reporte() {
	dot := "digraph G {\n"
	dot += "fontname=\"Arial\"\n"
	dot += "label=\"Cola - Clientes en Cola\"\n"
	dot += "labelloc = t\n"
	dot += "rankdir=LR;\n"
	dot += "node[shape=\"box\" fontname=\"Arial\"];\n"

	actual := nodo.Primero
	contador := 0
	for actual != nil {
		dot += "nodo_" + strconv.Itoa(contador) + " [label=\"ID: " + actual.Cliente.Id + "\\nNombre: " + actual.Cliente.Nombre + "\"];\n"
		contador++
		actual = actual.siguiente
	}

	actual = nodo.Primero
	contador = 0
	long := 1
	for actual != nil {
		if long < nodo.longitud {
			dot += "nodo_" + strconv.Itoa(contador) + " -> "
			long++
		} else {
			dot += "nodo_" + strconv.Itoa(contador)
		}
		if actual.siguiente == nil {
			dot += ";\n"
		}
		contador++
		actual = actual.siguiente
	}
	dot += "}\n"

	_generarGrafo(dot)
	_generarImg()
}

func _generarGrafo(dot string) {
	nombreArchivo := "./Reportes/Cola.dot"

	// Eliminar el archivo existente
	err := os.Remove(nombreArchivo)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error al eliminar el archivo:", err)
		return
	}

	// Crear un nuevo archivo
	file, err := os.Create(nombreArchivo)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	// Escribir el nuevo contenido en el archivo
	_, err = file.WriteString(dot)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}
}

func _generarImg() {
	// Ruta del archivo .dot de entrada
	inputFile := "./Reportes/Cola.dot"

	// Ruta del archivo de imagen de salida
	outputFile := "./Reportes/Cola.pdf"

	// Comando para ejecutar Graphviz
	cmd := exec.Command("dot", "-Tpdf", "-o", outputFile, inputFile)

	// Ejecutar el comando
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (c *ColaCliente) Mostrar() {
	actual := c.Primero
	formato := "  ║ %-10s %-20s║\n"
	println("  ╔════════════════════════════════╗")
	fmt.Printf(formato, "ID", "Nombre")
	println("  ╠════════════════════════════════╣")
	for actual != nil {
		fmt.Printf(formato, actual.Cliente.Id, actual.Cliente.Nombre)
		actual = actual.siguiente
	}
	println("  ╚════════════════════════════════╝")
}
