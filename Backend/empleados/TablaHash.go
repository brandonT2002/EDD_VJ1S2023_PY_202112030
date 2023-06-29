package empleados

import (
	"math"
	"strconv"
)

type TablaHash struct {
	Tabla       [30]NodoHash
	Capacidad   int
	Utilizacion int
}

func (t *TablaHash) calculoIndice(id_cliente int, multiplicador int) int {
	indice := (30*id_cliente + 202112030*multiplicador) % t.Capacidad
	return indice
}

func (t *TablaHash) capacidadTabla() {
	auxCap := float64(t.Capacidad) * 0.6
	if t.Utilizacion > int(auxCap) {
		t.Capacidad = t.nuevaCapacidad()
		t.Utilizacion = 0
		t.reInsertar()
	}
}

func (t *TablaHash) reInsertar() {
	auxTabla := t.Tabla
	t.NuevaTabla()
	for i := 0; i < 30; i++ {
		if auxTabla[i].Llave != -1 {
			t.Insertar(auxTabla[i].IdCliente, auxTabla[i].IdFactura)
		}
	}
}

func (t *TablaHash) NuevaTabla() {
	for i := 0; i < 30; i++ {
		t.Tabla[i].Llave = -1
		t.Tabla[i].IdCliente = ""
		t.Tabla[i].IdFactura = ""
	}
}

func (t *TablaHash) nuevaCapacidad() int {
	numero := t.Capacidad + 1
	for !t.esPrimo(numero) {
		numero++
	}
	return numero
}

func (t *TablaHash) esPrimo(numero int) bool {
	if numero <= 1 {
		return false
	}
	if numero == 2 {
		return true
	}
	if numero%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(numero))); i += 2 {
		if (numero % i) == 0 {
			return false
		}
	}
	return true
}

func (t *TablaHash) Insertar(id_cliente string, id_factura string) {
	numVar, _ := strconv.Atoi(id_cliente)
	indice := t.calculoIndice(numVar, 1)
	nuevoNodo := &NodoHash{indice, id_cliente, id_factura}
	if indice < t.Capacidad {
		if t.Tabla[indice].Llave == -1 {
			t.Tabla[indice] = *nuevoNodo
			t.Utilizacion++
			t.capacidadTabla()
		} else {
			indice = t.calculoIndice(numVar, 2)
			if t.Tabla[indice].Llave == -1 {
				nuevoNodo.Llave = indice
				t.Tabla[indice] = *nuevoNodo
				t.Utilizacion++
				t.capacidadTabla()
				return
			}
			for i := indice; i < t.Capacidad; i++ {
				if t.Tabla[i].Llave == -1 {
					nuevoNodo.Llave = i
					t.Tabla[i] = *nuevoNodo
					t.Utilizacion++
					t.capacidadTabla()
					return
				}
			}
		}
	}
}
