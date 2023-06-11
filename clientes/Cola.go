package clientes

type Cola struct {
	primero  *Nodo
	ultimo   *Nodo
	longitud int
}

func (c *Cola) Insertar(cliente *Cliente) {
	if c.primero != nil {
		c.ultimo.siguiente = &Nodo{cliente: cliente}
		c.ultimo = c.ultimo.siguiente
		c.longitud++
		return
	}
	c.primero = &Nodo{cliente: cliente}
	c.ultimo = c.primero
	c.longitud++
}

/*
func (l *Cola) eliminar() interface{} {
	if l.primero != nil {
		// temporal := l.primero.cliente
		l.primero = l.primero.siguiente
		l.longitud--
	}
	return nil
}
*/
