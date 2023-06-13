package imagencapas

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

type NodoCabeza struct {
	indice    int
	anterior  *NodoCabeza
	siguiente *NodoCabeza
	acceso    *NodoInterno
	ultimo    *NodoInterno
}

type NodoInterno struct {
	color     *Color
	fila      int
	columna   int
	derecha   *NodoInterno
	izquierda *NodoInterno
	arriba    *NodoInterno
	abajo     *NodoInterno
}

type ListaCabeza struct {
	primero *NodoCabeza
	ultimo  *NodoCabeza
}

func (nodo *ListaCabeza) insertar(indice int) {
	if nodo.primero != nil {
		if indice < nodo.primero.indice {
			nodo.primero.anterior = &NodoCabeza{indice: indice}
			nodo.primero.anterior.siguiente = nodo.primero
			nodo.primero = nodo.primero.anterior
		} else if indice > nodo.ultimo.indice {
			nodo.ultimo.siguiente = &NodoCabeza{indice: indice}
			nodo.ultimo.siguiente.anterior = nodo.ultimo
			nodo.ultimo = nodo.ultimo.siguiente
		} else {
			actual := nodo.primero
			for actual.siguiente != nil {
				if indice > actual.indice && indice < actual.siguiente.indice {
					tmp := &NodoCabeza{indice: indice}
					tmp.anterior = actual
					tmp.siguiente = actual.siguiente

					actual.siguiente.anterior = tmp
					actual.siguiente = tmp
					return
				}
				actual = actual.siguiente
			}
		}
		return
	}
	nodo.primero = &NodoCabeza{indice: indice}
	nodo.ultimo = nodo.primero
}

func (nodo *ListaCabeza) existeNodo(indice int) bool {
	if nodo.primero != nil {
		actual := nodo.primero
		for actual != nil {
			if actual.indice == indice {
				return true
			}
			actual = actual.siguiente
		}
	}
	return false
}

type MatrizDispersa struct {
	accesoF *ListaCabeza
	accesoC *ListaCabeza
}

func (nodo *MatrizDispersa) insertar(fila, columna int, color *Color) {
	if !nodo.accesoF.existeNodo(fila) {
		nodo.accesoF.insertar(fila)
	}
	if !nodo.accesoC.existeNodo(columna) {
		nodo.accesoC.insertar(columna)
	}
	nodoI := &NodoInterno{fila: fila, columna: columna, color: color}
	nodo.agregarFila(fila, nodoI)
	nodo.agregarColumna(columna, nodoI)
}

func (nodo *MatrizDispersa) agregarFila(fila int, nodoI *NodoInterno) {
	actualF := nodo.accesoF.primero
	for actualF != nil {
		if actualF.indice == fila {
			if actualF.acceso != nil {
				if nodoI.columna < actualF.acceso.columna {
					actualF.acceso.izquierda = nodoI
					actualF.acceso.izquierda.derecha = actualF.acceso
					actualF.acceso = actualF.acceso.izquierda
				} else if nodoI.columna > actualF.ultimo.columna {
					actualF.ultimo.derecha = nodoI
					actualF.ultimo.derecha.izquierda = actualF.ultimo
					actualF.ultimo = actualF.ultimo.derecha
				} else {
					actualC := actualF.acceso
					for actualC.derecha != nil {
						if nodoI.columna > actualC.columna && nodoI.columna < actualC.derecha.columna {
							nodoI.izquierda = actualC
							nodoI.derecha = actualC.derecha

							actualC.derecha.izquierda = nodoI
							actualC.derecha = nodoI
							return
						}
						actualC = actualC.derecha
					}
				}
				return
			}
			actualF.acceso = nodoI
			actualF.ultimo = actualF.acceso
			return
		}
		actualF = actualF.siguiente
	}
}

func (nodo *MatrizDispersa) agregarColumna(columna int, nodoI *NodoInterno) {
	actualC := nodo.accesoC.primero
	for actualC != nil {
		if actualC.indice == columna {
			if actualC.acceso != nil {
				if nodoI.fila < actualC.acceso.fila {
					actualC.acceso.arriba = nodoI
					actualC.acceso.arriba.abajo = actualC.acceso
					actualC.acceso = actualC.acceso.arriba
				} else if nodoI.fila > actualC.ultimo.fila {
					actualC.ultimo.abajo = nodoI
					actualC.ultimo.abajo.arriba = actualC.ultimo
					actualC.ultimo = actualC.ultimo.abajo
				} else {
					actualF := actualC.acceso
					for actualF.abajo != nil {
						if nodoI.fila > actualF.fila && nodoI.fila < actualF.abajo.fila {
							nodoI.arriba = actualF
							nodoI.abajo = actualF.abajo

							actualF.abajo.arriba = nodoI
							actualF.abajo = nodoI
							return
						}
						actualF = actualF.abajo
					}
				}
				return
			}
			actualC.acceso = nodoI
			actualC.ultimo = actualC.acceso
			return
		}
		actualC = actualC.siguiente
	}
}

func (nodo *MatrizDispersa) dot() string {
	dot := "digraph T{\n\tnode[shape=box fontname=\"Arial\" fillcolor=\"white\" style=filled];"
	dot += "\n\tRoot[label = \"Capa 0\", group=\"0\"];"

	actualF := nodo.accesoF.primero
	for actualF != nil {
		dot += "\n\tF" + strconv.Itoa(actualF.indice) + "[group=\"0\" fillcolor=\"plum\"];"
		actualF = actualF.siguiente
	}

	actualC := nodo.accesoC.primero
	for actualC != nil {
		dot += "\n\tC" + strconv.Itoa(actualC.indice) + "[group=" + strconv.Itoa(actualC.indice) + " fillcolor=\"powderblue\"];"
		actualC = actualC.siguiente
	}

	actualC = nodo.accesoC.primero
	for actualC != nil {
		actualF := actualC.acceso
		for actualF != nil {
			dot += "\n\tN" + strconv.Itoa(actualF.fila) + "_" + strconv.Itoa(actualF.columna) + "[group=" + strconv.Itoa(actualF.columna) + " label=\"" + actualF.color.R + "-" + actualF.color.G + "-" + actualF.color.B + "\"];"
			actualF = actualF.abajo
		}
		actualC = actualC.siguiente
	}

	dot += "\n\tsubgraph columnHeader {\n\t\trank = same;"
	enlace := "\n\t\tRoot -> "
	actualC = nodo.accesoC.primero
	for actualC != nil {
		enlace += "C" + strconv.Itoa(actualC.indice)
		actualC = actualC.siguiente
		if actualC != nil {
			enlace += " -> "
		}
	}
	dot += enlace + ";" + enlace + "[dir=back];\n\t}"

	actualF = nodo.accesoF.primero
	for actualF != nil {
		dot += "\n\tsubgraph row" + strconv.Itoa(actualF.indice) + " {\n\t\trank = same;"
		enlace := "\n\t\tF" + strconv.Itoa(actualF.indice) + " -> "
		actualC := actualF.acceso
		for actualC != nil {
			enlace += "N" + strconv.Itoa(actualC.fila) + "_" + strconv.Itoa(actualC.columna)
			if actualC.derecha != nil {
				actualC = actualC.derecha
			} else {
				break
			}
			if actualC != nil {
				enlace += " -> "
			}
		}
		dot += enlace + ";" + enlace + "[dir=back];\n\t}"
		actualF = actualF.siguiente
	}

	dot += "\n\tsubgraph rowHeader {"
	enlace = "\n\t\tRoot -> "
	actualF = nodo.accesoF.primero
	for actualF != nil {
		enlace += "F" + strconv.Itoa(actualF.indice)
		actualF = actualF.siguiente
		if actualF != nil {
			enlace += " -> "
		}
	}
	dot += enlace + ";" + enlace + "[dir=back];\n\t}"

	actualC = nodo.accesoC.primero
	for actualC != nil {
		dot += "\n\tsubgraph column" + strconv.Itoa(actualC.indice) + " {"
		enlace = "\n\t\tC" + strconv.Itoa(actualC.indice) + " -> "
		actualF := actualC.acceso
		for actualF != nil {
			enlace += "N" + strconv.Itoa(actualF.fila) + "_" + strconv.Itoa(actualF.columna)
			if actualF.abajo != nil {
				actualF = actualF.abajo
			} else {
				break
			}
			if actualF != nil {
				enlace += " -> "
			}
		}
		dot += enlace + ";" + enlace + "[dir=back];\n\t}"
		actualC = actualC.siguiente
	}
	dot += "\n}"

	return dot
}

func (nodo *MatrizDispersa) ObtenerCSS(ancho int) string {
	actualF := nodo.accesoF.primero
	css := ""
	for actualF != nil {
		actualC := actualF.acceso
		for actualC != nil {
			css += fmt.Sprintf("\n.pixel:nth-child(%d) {background: rgb(%s,%s,%s);}", actualF.indice*ancho+actualC.columna+1, actualC.color.R, actualC.color.G, actualC.color.B)
			actualC = actualC.derecha
		}
		actualF = actualF.siguiente
	}
	return css
}

func generarGrafo(nombreArchivo, dot string) {
	nombreArchivo += ".dot"
	nuevoContenido := dot

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
	_, err = file.WriteString(nuevoContenido)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

	fmt.Println("Se ha escrito el nuevo contenido en el archivo.")
}

func generarImg() {
	// Ruta del archivo .dot de entrada
	inputFile := "grafo.dot"

	// Ruta del archivo de imagen de salida
	outputFile := "grafo.png"

	// Comando para ejecutar Graphviz
	cmd := exec.Command("dot", "-Tpng", "-o", outputFile, inputFile)

	// Ejecutar el comando
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Archivo de imagen generado:", outputFile)
}
