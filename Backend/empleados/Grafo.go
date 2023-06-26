package empleados

import "strconv"

type Grafo struct {
	Principal *NodoMatrizA
}

func (g *Grafo) InsertarC(padre, hijo, filtro string) {
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
	if g.Principal == nil {
		g.insertarF(envio.Padre)
		g.InsertarC(envio.Padre, envio.Imagen, "")
		g.InsertarC(envio.Cliente, envio.Imagen, envio.Filtros)
	} else {
		g.InsertarC(envio.Padre, envio.Cliente, "")
		g.InsertarC(envio.Cliente, envio.Imagen, envio.Filtros)
	}
}

func (g *Grafo) Mjson() []EnvioMatriz {
	var listaEnvios []EnvioMatriz

	if g.Principal != nil {
		aux := g.Principal.Abajo
		aux1 := aux

		for aux != nil {
			for aux1 != nil {
				envio := &EnvioMatriz{Padre: g.Principal.valor, Cliente: aux.valor, Imagen: aux1.valor, Filtros: ""}
				listaEnvios = append(listaEnvios, *envio)
				aux1 = aux1.Siguiente
			}
			if aux != nil {
				aux = aux.Abajo
				aux1 = aux
			}
		}
	}
	return listaEnvios
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
			cadena += "nodo" + strconv.Itoa(x+2) + strconv.Itoa(y) + "[len=2.00]; \n"
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
		// dot += "graph grafoDirigido{ \n rankdir=LR; \n node [shape=box]; layout=neato; \n nodo00[label=\"" + g.Principal.valor + "\"]; \n"
		dot += "graph grafoDirigido{\n"
		dot += "rankdir=LR;\n"
		dot += "node [shape=box];\n"
		dot += "layout=neato;"
		dot += "nodo00[label=\"" + g.Principal.valor + "\"]; "
		dot += g.matriz()
		dot += "\n}"
	}
	return dot
}
