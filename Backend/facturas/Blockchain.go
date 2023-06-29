package facturas

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type Blockchain struct {
	primero  *NodoB
	ultimo   *NodoB
	longitud int
}

func (b *Blockchain) Insertar(fecha, biller, customer, payment string) {
	cadenaFuncion := strconv.Itoa(b.longitud) + fecha + biller + customer + payment
	hash := SHA256(cadenaFuncion)
	if b.primero != nil {
		datosBloque := map[string]string{
			"index":        strconv.Itoa(b.longitud),
			"timestamp":    fecha,
			"biller":       biller,
			"customer":     customer,
			"payment":      payment,
			"previoushash": b.ultimo.bloque["hash"],
			"hash":         hash,
		}
		b.ultimo.siguiente = &NodoB{bloque: datosBloque}
		b.ultimo.siguiente.anterior = b.ultimo
		b.ultimo = b.ultimo.siguiente
		b.longitud++
		return
	}
	datosBloque := map[string]string{
		"index":        strconv.Itoa(b.longitud),
		"timestamp":    fecha,
		"biller":       biller,
		"customer":     customer,
		"payment":      payment,
		"previoushash": "0000",
		"hash":         hash,
	}
	nuevoBloque := &NodoB{bloque: datosBloque}
	b.primero = nuevoBloque
	b.ultimo = b.primero
	b.longitud++
}

func (b *Blockchain) Facturas() []Respuesta {
	actual := b.primero
	var arreglo []Respuesta
	for actual != nil {
		arreglo = append(arreglo, Respuesta{Id: actual.bloque["customer"], Factura: actual.bloque["hash"]})
		actual = actual.siguiente
	}
	return arreglo
}

func (b *Blockchain) Reporte() string {
	dot := "digraph G {\n"
	dot += "fontname=\"Arial\""
	dot += "label=\"Lista Doble - Imagenes\"\n"
	dot += "labelloc = t\n"
	dot += "rankdir=LR;\n"
	dot += "node[shape=\"box\" fontname=\"Arial\"];\n"

	dot += "null0 [label=\"null\"]\n"
	dot += "null1 [label=\"null\"]\n"

	actual := b.primero
	long := 0
	for actual != nil {
		dot += "nodo_" + strconv.Itoa(long) +
			" [label=\"TimeStamp:" + actual.bloque["timestamp"] +
			"\\nBiller: " + actual.bloque["biller"] +
			"\\nCustomer" + actual.bloque["customer"] +
			"\\nPreviousHash: " + actual.bloque["previoushash"] +
			"\\nPreviousHash: " + actual.bloque["hash"] + "\"];\n"
		long++
		actual = actual.siguiente
	}

	actual = b.primero
	long = 0
	for actual != nil {
		if long < b.longitud {
			dot += "nodo_" + strconv.Itoa(long) + " -> "
			long++
		} else {
			dot += "nodo_" + strconv.Itoa(long)
		}
		if actual.siguiente == nil {
			dot += ";\n"
		}
		actual = actual.siguiente
	}

	dot += "}"
	return dot
}

func SHA256(cadena string) string {
	hexaString := ""
	h := sha256.New()
	h.Write([]byte(cadena))
	hash := h.Sum(nil)
	hexaString = hex.EncodeToString(hash)
	return hexaString
}
