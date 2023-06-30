package pedidos

import (
	"fmt"
	"strconv"
)

type ArbolAVL struct {
	raiz *Nodo
}

func (a *ArbolAVL) Insertar(pedido *Pedido) {
	a.raiz = a.Insertar_(pedido, a.raiz)
}

func (a *ArbolAVL) Insertar_(pedido *Pedido, nodo *Nodo) *Nodo {
	if nodo == nil {
		return &Nodo{
			pedido: pedido,
			izq:    nil,
			der:    nil,
		}
	}
	if pedido.IdCliente < nodo.pedido.IdCliente {
		nodo.izq = a.Insertar_(pedido, nodo.izq)
		if a.getAltura(nodo.izq)-a.getAltura(nodo.der) == 2 {
			if pedido.IdCliente < nodo.izq.pedido.IdCliente {
				nodo = a.rotarIzq(nodo)
			} else {
				nodo = a.rotarDobleIzq(nodo)
			}
		}
	} else if pedido.IdCliente > nodo.pedido.IdCliente {
		nodo.der = a.Insertar_(pedido, nodo.der)
		if a.getAltura(nodo.der)-a.getAltura(nodo.izq) == 2 {
			if pedido.IdCliente > nodo.der.pedido.IdCliente {
				nodo = a.rotarDer(nodo)
			} else {
				nodo = a.rotarDobleDer(nodo)
			}
		}
	}
	nodo.altura = a.alturaMax(a.getAltura(nodo.izq), a.getAltura(nodo.der)) + 1
	return nodo
}

func (a *ArbolAVL) alturaMax(izq, der int) int {
	if izq > der {
		return izq
	}
	return der
}

func (a *ArbolAVL) getAltura(nodo *Nodo) int {
	if nodo == nil {
		return -1
	}
	return nodo.altura
}

func (a *ArbolAVL) rotarIzq(nodo *Nodo) *Nodo {
	auxiliar := nodo.izq
	nodo.izq = auxiliar.der
	auxiliar.der = nodo
	nodo.altura = a.alturaMax(a.getAltura(nodo.izq), a.getAltura(nodo.der)) + 1
	auxiliar.altura = a.alturaMax(a.getAltura(auxiliar.izq), nodo.altura) + 1
	return auxiliar
}

func (a *ArbolAVL) rotarDer(nodo *Nodo) *Nodo {
	auxiliar := nodo.der
	nodo.der = auxiliar.izq
	auxiliar.izq = nodo
	nodo.altura = a.alturaMax(a.getAltura(nodo.izq), a.getAltura(nodo.der)) + 1
	auxiliar.altura = a.alturaMax(a.getAltura(auxiliar.der), nodo.altura) + 1
	return auxiliar
}

func (a *ArbolAVL) rotarDobleIzq(nodo *Nodo) *Nodo {
	nodo.izq = a.rotarDer(nodo.izq)
	return a.rotarIzq(nodo)
}

func (a *ArbolAVL) rotarDobleDer(nodo *Nodo) *Nodo {
	nodo.der = a.rotarIzq(nodo.der)
	return a.rotarDer(nodo)
}

func (a *ArbolAVL) Reporte() string {
	dot := "digraph Tree{\n\tnode [shape = record];"
	dot += "fontname=\"Arial\""
	dot += "bgcolor=\"#282A37\";\n"
	dot += "edge[color=\"white\"]"
	if a.raiz != nil {
		dot += a.dot1(a.raiz)
	}
	dot += "\n}"
	return dot
}

func (a *ArbolAVL) Inorder(cola *ColaPedidos) {
	a.inorder1(a.raiz, cola)
	fmt.Println()
}

func (a *ArbolAVL) inorder1(nodo *Nodo, cola *ColaPedidos) {
	if nodo != nil {
		a.inorder1(nodo.izq, cola)
		// fmt.Print(strconv.Itoa(nodo.pedido.IdCliente) + " ")
		cola.Insertar(nodo.pedido)
		a.inorder1(nodo.der, cola)
	}
}

func (a *ArbolAVL) dot1(nodo *Nodo) string {
	dot := ""
	if nodo.izq == nil && nodo.der == nil {
		dot = "\n\tnode_" + strconv.Itoa(nodo.pedido.IdCliente) + "[label=\"<C3>" + strconv.Itoa(nodo.pedido.IdCliente) + "\\n" + nodo.pedido.Imagen + "\" fontname=\"Arial\" color=\"white\" fontcolor=\"white\" fillcolor=\"#282A37\"];"
	} else {
		dot = "\n\tnode_" + strconv.Itoa(nodo.pedido.IdCliente) + "[label=\"<C0> | <C3>" + strconv.Itoa(nodo.pedido.IdCliente) + "\\n" + nodo.pedido.Imagen + " | <C1>\" fontname=\"Arial\" color=\"white\" fontcolor=\"white\" fillcolor=\"#282A37\"];"
	}
	if nodo.izq != nil {
		dot += a.dot1(nodo.izq)
		dot += "\n\tnode_" + strconv.Itoa(nodo.pedido.IdCliente) + ":C0 -> node_" + strconv.Itoa(nodo.izq.pedido.IdCliente) + ":C3;"
	}
	if nodo.der != nil {
		dot += a.dot1(nodo.der)
		dot += "\n\tnode_" + strconv.Itoa(nodo.pedido.IdCliente) + ":C1 -> node_" + strconv.Itoa(nodo.der.pedido.IdCliente) + ":C3;"
	}
	return dot
}
