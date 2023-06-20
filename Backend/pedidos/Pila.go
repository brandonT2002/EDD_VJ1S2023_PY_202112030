package pedidos

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Pila struct {
	ultimo *Nodo
}

func (p *Pila) Insertar(pedido *Pedido) {
	if p.ultimo != nil {
		nuevoNodo := &Nodo{pedido: pedido}
		nuevoNodo.anterior = p.ultimo
		p.ultimo = nuevoNodo
		return
	}
	p.ultimo = &Nodo{pedido: pedido}
}

func (nodo *Pila) Reporte() {
	dot := "digraph G {\n"
	dot += "fontname=\"Arial\"\n"
	dot += "label=\"Pila - Pedidos\"\n"
	dot += "labelloc=t\nstack [shape=none, margin=0, label=<<TABLE BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"0\" CELLPADDING=\"2\">\n"

	actual := nodo.ultimo
	if actual == nil {
		dot += "<tr><td></td></tr>"
	}
	for actual != nil {
		dot += "<tr>\n"
		dot += "<td width=\"125\" height=\"50\"><font point-size=\"15\">ID Cliente: " + actual.pedido.IdCliente + "<BR/>Imagen: " + actual.pedido.Imagen + "</font></td>\n"
		dot += "</tr>\n"
		actual = actual.anterior
	}
	dot += "</TABLE>>];\n"
	dot += "}\n"

	generarGrafo(dot)
	generarImg()
}

func generarGrafo(dot string) {
	nombreArchivo := "./Reportes/Pila.dot"

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
	inputFile := "./Reportes/Pila.dot"

	// Ruta del archivo de imagen de salida
	outputFile := "./Reportes/Pila.pdf"

	// Comando para ejecutar Graphviz
	cmd := exec.Command("dot", "-Tpdf", "-o", outputFile, inputFile)

	// Ejecutar el comando
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (p *Pila) Json() ([]byte, error) {
	actual := p.ultimo
	var pedidos []map[string]interface{}
	for actual != nil {
		pedido := map[string]interface{}{
			"IdCliente": actual.pedido.IdCliente,
			"Imagen":    actual.pedido.Imagen,
		}
		pedidos = append(pedidos, pedido)
		actual = actual.anterior
	}
	return json.MarshalIndent(pedidos, "", "	")
}

func (p *Pila) Mostrar() {
	actual := p.ultimo
	formato := "  ║ %-10s %-20s║\n"
	println("  ╔════════════════════════════════╗")
	fmt.Printf(formato, "ID Cliente", "Imagen")
	println("  ╠════════════════════════════════╣")
	for actual != nil {
		fmt.Printf(formato, actual.pedido.IdCliente, actual.pedido.Imagen)
		actual = actual.anterior
	}
	println("  ╚════════════════════════════════╝")
}
