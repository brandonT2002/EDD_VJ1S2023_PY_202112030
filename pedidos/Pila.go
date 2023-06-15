package pedidos

import "fmt"

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
