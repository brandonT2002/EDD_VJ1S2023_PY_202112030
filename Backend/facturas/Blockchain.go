package facturas

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type Blockchain struct {
	primero  *NodoB
	ultimo   *NodoB
	Longitud int
}

func (b *Blockchain) Insertar(bloque *Peticion) string {
	cadenaFuncion := strconv.Itoa(b.Longitud) + bloque.Timestamp + bloque.Biller + bloque.Customer + bloque.Payment
	hash := SHA256(cadenaFuncion)
	if b.primero != nil {
		datosBloque := map[string]string{
			"index":        strconv.Itoa(b.Longitud),
			"timestamp":    bloque.Timestamp,
			"biller":       bloque.Biller,
			"customer":     bloque.Customer,
			"payment":      bloque.Payment,
			"previoushash": b.ultimo.bloque["hash"],
			"hash":         hash,
		}
		b.ultimo.siguiente = &NodoB{bloque: datosBloque}
		b.ultimo.siguiente.anterior = b.ultimo
		b.ultimo = b.ultimo.siguiente
		b.Longitud++
		return hash
	}
	datosBloque := map[string]string{
		"index":        strconv.Itoa(b.Longitud),
		"timestamp":    bloque.Timestamp,
		"biller":       bloque.Biller,
		"customer":     bloque.Customer,
		"payment":      bloque.Payment,
		"previoushash": "0000",
		"hash":         hash,
	}
	nuevoBloque := &NodoB{bloque: datosBloque}
	b.primero = nuevoBloque
	b.ultimo = b.primero
	b.Longitud++
	return hash
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
	dot += "bgcolor=\"#282A37\";\n"
	dot += "edge[color=\"white\"]"
	dot += "rankdir=LR;\n"
	dot += "node[shape=\"box\" fontname=\"Arial\" color=\"white\" fontcolor=\"white\" fillcolor=\"#282A37\"];\n"

	actual := b.primero
	long := 0
	for actual != nil {
		dot += "nodo_" + strconv.Itoa(long) +
			" [label=\"TimeStamp:" + actual.bloque["timestamp"] +
			"\\nBiller: " + actual.bloque["biller"] +
			"\\nCustomer: " + actual.bloque["customer"] +
			"\\nPreviousHash: " + actual.bloque["previoushash"] +
			"\\nHash: " + actual.bloque["hash"] + "\"];\n"
		long++
		actual = actual.siguiente
	}

	actual = b.primero
	long = 0
	for actual != nil {
		if actual.siguiente == nil {
			dot += "nodo_" + strconv.Itoa(long)
			dot += ";\n"
		} else {
			dot += "nodo_" + strconv.Itoa(long) + " -> "
		}
		long++
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
