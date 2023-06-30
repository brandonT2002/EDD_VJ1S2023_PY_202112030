package pedidos

type NodoC struct {
	Pedido    *Pedido
	siguiente *NodoC
}
