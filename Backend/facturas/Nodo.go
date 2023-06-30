package facturas

type NodoB struct {
	bloque    map[string]string
	siguiente *NodoB
	anterior  *NodoB
}

type Peticion struct {
	Timestamp string
	Biller    string
	Customer  string
	Payment   string
}

type Respuesta struct {
	Id      string
	Factura string
}
