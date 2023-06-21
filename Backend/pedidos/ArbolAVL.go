package pedidos

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
