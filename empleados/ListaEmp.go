package empleados

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

type ListaEmp struct {
	primero  *Nodo
	ultimo   *Nodo
	longitud int
}

func (l *ListaEmp) Insertar(empleado *Empleado) {
	if l.primero != nil {
		l.ultimo.siguiente = &Nodo{empleado: empleado}
		l.ultimo = l.ultimo.siguiente
		l.longitud++
		return
	}
	l.primero = &Nodo{empleado: empleado}
	l.ultimo = l.primero
	l.longitud++
}

func (nodo *ListaEmp) Reporte() {
	dot := "digraph G {\n"
	dot += "fontname=\"Arial\"\n"
	dot += "label=\"Lista Simple - Empleados\"\n"
	dot += "labelloc = t\n"
	dot += "rankdir=LR;\n"
	dot += "node[shape=\"box\" fontname=\"Arial\"];\n"

	actual := nodo.primero
	contador := 0
	for actual != nil {
		dot += "nodo_" + strconv.Itoa(contador) + " [label=\"ID: " + actual.empleado.Id + "\\nNombre: " + actual.empleado.Nombre + "\"];\n"
		contador++
		actual = actual.siguiente
	}

	actual = nodo.primero
	contador = 0
	for actual != nil {
		dot += "nodo_" + strconv.Itoa(contador) + " -> "
		if actual.siguiente == nil {
			dot += "null;\n"
		}
		contador++
		actual = actual.siguiente
	}
	dot += "}"

	generarGrafo(dot)
	generarImg()
}

func generarGrafo(dot string) {
	nombreArchivo := "./Reportes/ListaSimple.dot"

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

func generarImg() {
	// Ruta del archivo .dot de entrada
	inputFile := "./Reportes/ListaSimple.dot"

	// Ruta del archivo de imagen de salida
	outputFile := "./Reportes/ListaSimple.pdf"

	// Comando para ejecutar Graphviz
	cmd := exec.Command("dot", "-Tpdf", "-o", outputFile, inputFile)

	// Ejecutar el comando
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (l *ListaEmp) Mostrar() {
	actual := l.primero
	formato := "  ║ %-10s %-25s %-15s %-15s║\n"
	println("  ╔═════════════════════════════════════════════════════════════════════╗")
	fmt.Printf(formato, "ID", "Nombre", "Cargo", "Contraseña")
	println("  ╠═════════════════════════════════════════════════════════════════════╣")
	for actual != nil {
		fmt.Printf(formato, actual.empleado.Id, actual.empleado.Nombre, actual.empleado.Cargo, actual.empleado.Contrasena)
		actual = actual.siguiente
	}
	println("  ╚═════════════════════════════════════════════════════════════════════╝")
}
