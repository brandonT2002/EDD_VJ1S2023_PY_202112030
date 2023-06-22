package pedidos

type Nodo struct {
	pedido *Pedido
	altura int
	izq    *Nodo
	der    *Nodo
}
