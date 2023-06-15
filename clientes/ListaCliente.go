package clientes

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

type ListaCliente struct {
	primero  *Nodo
	ultimo   *Nodo
	longitud int
}

func (l *ListaCliente) Insertar(cliente *Cliente) {
	if l.Existe(cliente.Id) {
		return
	}
	if l.primero != nil {
		l.ultimo.siguiente = &Nodo{Cliente: cliente}
		l.ultimo = l.ultimo.siguiente
		l.ultimo.siguiente = l.primero
		l.longitud++
		return
	}
	l.primero = &Nodo{Cliente: cliente}
	l.ultimo = l.primero
	l.ultimo.siguiente = l.primero
	l.longitud++
}

func (l *ListaCliente) Existe(id string) bool {
	actual := l.primero
	contador := 0
	for actual != nil {
		if contador == l.longitud {
			break
		}
		if actual.Cliente.Id == id && l.esNumero(id) {
			return true
		}
		actual = actual.siguiente
		contador++
	}
	return false
}

func (l *ListaCliente) esNumero(id string) bool {
	_, err := strconv.Atoi(id)
	if err == nil {
		return true
	}
	return false
}

func (l *ListaCliente) GuardarId(nombre, id string) {
	actual := l.primero
	for actual != nil {
		if actual.Cliente.Nombre == nombre && (actual.Cliente.Id == "x" || actual.Cliente.Id == "X") {
			actual.Cliente.Id = id
			return
		}
		actual = actual.siguiente
	}
}

func (nodo *ListaCliente) Reporte() {
	dot := "digraph G {\n"
	dot += "fontname=\"Arial\""
	dot += "label=\"Lista Circular\"\n"
	dot += "labelloc = t\n"
	dot += "rankdir=LR;\n"
	dot += "node[shape=\"box\" fontname=\"Arial\"];\n"

	actual := nodo.primero
	contador := 0
	for actual != nil {
		dot += "nodo_" + strconv.Itoa(contador) + " [label=\"" + actual.Cliente.Nombre + "\"];\n"
		contador++
		if contador == nodo.longitud {
			break
		}
		actual = actual.siguiente
	}

	actual = nodo.primero
	contador = 0
	for actual != nil {
		dot += "nodo_" + strconv.Itoa(contador) + " -> "
		contador++
		if nodo.longitud == contador {
			dot += "nodo_0;\n"
		}
		if contador == nodo.longitud {
			break
		}
		actual = actual.siguiente
	}
	dot += "}"

	generarGrafo(dot)
	generarImg()
}

func generarGrafo(dot string) {
	nombreArchivo := "./Reportes/ListaCircular.dot"

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
	inputFile := "./Reportes/ListaCircular.dot"

	// Ruta del archivo de imagen de salida
	outputFile := "./Reportes/ListaCircular.pdf"

	// Comando para ejecutar Graphviz
	cmd := exec.Command("dot", "-Tpdf", "-o", outputFile, inputFile)

	// Ejecutar el comando
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (l *ListaCliente) Mostrar() {
	actual := l.primero
	contador := 0
	formato := "  ║ %-10s %-20s║\n"
	println("  ╔════════════════════════════════╗")
	fmt.Printf(formato, "ID", "Nombre")
	println("  ╠════════════════════════════════╣")
	for actual != nil {
		fmt.Printf(formato, actual.Cliente.Id, actual.Cliente.Nombre)
		contador++
		if contador == l.longitud {
			break
		}
		actual = actual.siguiente
	}
	println("  ╚════════════════════════════════╝")
}
