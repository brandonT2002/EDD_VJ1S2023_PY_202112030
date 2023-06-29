package empleados

import "strconv"

type Grafo struct {
	Principal *NodoMatrizA
}

func (g *Grafo) InsertarC(padre, hijo string, filtro string) {
	nuevoNodo := &NodoMatrizA{valor: hijo}
	if g.Principal != nil && padre == g.Principal.valor {
		aux := g.Principal
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = nuevoNodo
	} else {
		g.insertarF(padre)
		aux := g.Principal
		for aux != nil {
			if aux.valor == padre {
				break
			}
			aux = aux.Abajo
		}
		if aux != nil {
			nuevoNodo.Siguiente = &NodoMatrizA{valor: filtro}
			aux.Siguiente = nuevoNodo
		}
	}
}

func (g *Grafo) insertarF(padre string) {
	nuevoNodo := &NodoMatrizA{valor: padre}
	if g.Principal == nil {
		g.Principal = nuevoNodo
	} else {
		aux := g.Principal
		for aux.Abajo != nil {
			if aux.valor == padre {
				return
			}
			aux = aux.Abajo
		}
		aux.Abajo = nuevoNodo
	}
}

func (g *Grafo) InsertarValores(envio *EnvioMatriz) {
	fs := g.obtenerFiltros(envio.Filtros)
	if g.Principal == nil {
		g.insertarF(envio.Padre)
		g.InsertarC(envio.Padre, envio.Imagen, "")
		g.InsertarC(envio.Cliente, envio.Imagen, fs)
	} else {
		g.InsertarC(envio.Padre, envio.Cliente, "")
		g.InsertarC(envio.Cliente, envio.Imagen, fs)
	}
}

func (g *Grafo) obtenerFiltros(filtros *EnvioFiltros) string {
	fs := ""
	if filtros.N {
		fs += "Negativo"
	}
	if filtros.G {
		if fs != "" {
			fs += " - "
		}
		fs += "Grises"
	}
	if filtros.EX {
		if fs != "" {
			fs += " - "
		}
		fs += "Espejo X"
	}
	if filtros.EY {
		if fs != "" {
			fs += " - "
		}
		fs += "Espejo Y"
	}
	if filtros.DE {
		if fs != "" {
			fs += " - "
		}
		fs += "Doble Espejo"
	}
	if fs == "" {
		fs += "Sin Filtros"
	}
	return fs
}

func (g *Grafo) Mjson() []Solicitud {
	var solicitudes []Solicitud
	x := 0
	y := 1

	if g.Principal == nil || g.Principal.Abajo == nil {
		return solicitudes
	}

	aux := g.Principal.Abajo
	aux1 := aux

	for aux != nil {
		sol := &Solicitud{}
		for aux1 != nil {
			if x == 0 {
				sol.Cliente = aux1.valor
			} else if x == 1 {
				sol.Imagen = aux1.valor
			} else {
				sol.Filtros = aux1.valor
			}
			aux1 = aux1.Siguiente
			x++
		}
		if aux != nil {
			aux = aux.Abajo
			aux1 = aux
		}
		solicitudes = append(solicitudes, *sol)
		x = 0
		y++
	}
	return solicitudes
}

func (g *Grafo) matriz() string {
	cadena := ""
	x := 0
	y := 1

	aux := g.Principal.Abajo
	aux1 := aux

	for aux != nil {
		for aux1 != nil {
			cadena += "nodo" + strconv.Itoa(x) + strconv.Itoa(y) + "[label=\"" + aux1.valor + "\" ]; \n"
			aux1 = aux1.Siguiente
			x++
		}
		if aux != nil {
			aux = aux.Abajo
			aux1 = aux
		}
		x = 0
		y++
	}

	x = 0
	y = 1
	aux = g.Principal.Abajo
	aux1 = aux

	for aux != nil {
		if aux1 != nil {
			cadena += "nodo00 -- "
			cadena += "nodo" + strconv.Itoa(x) + strconv.Itoa(y) + " -- "
			cadena += "nodo" + strconv.Itoa(x+1) + strconv.Itoa(y) + " -- "
			cadena += "nodo" + strconv.Itoa(x+2) + strconv.Itoa(y) + "[len=2.25]; \n"
		}
		if aux != nil {
			aux = aux.Abajo
			aux1 = aux
		}
		x = 0
		y++
	}

	return cadena
}

func (g *Grafo) Dot() string {
	dot := ""
	if g.Principal != nil {
		dot += "graph grafoDirigido{\n"
		dot += "rankdir=LR;\n"
		dot += "bgcolor=\"#282A37\";\n"
		dot += "edge[color=\"white\"]"
		dot += "node [shape=box color=\"white\" fontcolor=\"white\" fillcolor=\"#282A37\"];\n"
		dot += "layout=neato;"
		dot += "nodo00[label=\"Empleado\\n" + g.Principal.valor + "\"]; "
		dot += g.matriz()
		dot += "}"
	}
	return dot
}
