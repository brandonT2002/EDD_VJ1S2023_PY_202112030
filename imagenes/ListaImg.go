package imagenes

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

type ListaImg struct {
	primero  *Nodo
	ultimo   *Nodo
	longitud int
}

func (l *ListaImg) Insertar(imagen *Imagen) {
	if l.primero != nil {
		l.ultimo.siguiente = &Nodo{imagen: imagen}
		l.ultimo.siguiente.anterior = l.ultimo
		l.ultimo = l.ultimo.siguiente
		l.longitud++
		return
	}
	l.primero = &Nodo{imagen: imagen}
	l.ultimo = l.primero
	l.longitud++
}

func (l *ListaImg) Buscar(imagen string) bool {
	actual := l.primero
	for actual != nil {
		if actual.imagen.Nombre == imagen {
			return true
		}
		actual = actual.siguiente
	}
	return false
}

func (nodo *ListaImg) Reporte() {
	dot := "digraph G {\n"
	dot += "label=\"Lista Doble\"\n"
	dot += "labelloc = t\n"
	dot += "rankdir=LR;\n"
	dot += "node[shape=\"box\" fontname=\"Arial\"];\n"

	dot += "null0 [label=\"null\"]\n"
	dot += "null1 [label=\"null\"]\n"

	actual := nodo.primero
	contador := 0
	for actual != nil {
		dot += "nodo_" + strconv.Itoa(contador) + " [label=\"" + actual.imagen.Nombre + "\"];\n"
		contador++
		actual = actual.siguiente
	}

	actual = nodo.primero
	contador = 0
	for actual != nil {
		dot += "nodo_" + strconv.Itoa(contador) + " -> "
		if actual.siguiente == nil {
			dot += "null0;\n"
		}
		contador++
		actual = actual.siguiente
	}

	actual = nodo.primero
	contador = 0
	dot += "null1"
	for actual != nil {
		dot += " -> " + "nodo_" + strconv.Itoa(contador)
		if actual.siguiente == nil {
			dot += "[dir=\"back\"] ;\n"
		}
		contador++
		actual = actual.siguiente
	}
	dot += "}"

	generarGrafo(dot)
	generarImg()
}

func generarGrafo(dot string) {
	nombreArchivo := "./Reportes/ListaDoble.dot"

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
	inputFile := "./Reportes/ListaDoble.dot"

	// Ruta del archivo de imagen de salida
	outputFile := "./Reportes/ListaDoble.pdf"

	// Comando para ejecutar Graphviz
	cmd := exec.Command("dot", "-Tpdf", "-o", outputFile, inputFile)

	// Ejecutar el comando
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (l *ListaImg) Mostrar() {
	actual := l.primero
	formato := "  ║ %-17s %-10s║\n"
	println("  ╔═════════════════════════════╗")
	fmt.Printf(formato, "Imagen", "Capas")
	println("  ╠═════════════════════════════╣")
	for actual != nil {
		fmt.Printf(formato, actual.imagen.Nombre, actual.imagen.Capas)
		actual = actual.siguiente
	}
	println("  ╚═════════════════════════════╝")
}
