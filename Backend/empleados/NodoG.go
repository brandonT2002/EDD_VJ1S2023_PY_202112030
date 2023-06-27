package empleados

type NodoMatrizA struct {
	Siguiente *NodoMatrizA
	Abajo     *NodoMatrizA
	valor     string
}

type EnvioFiltros struct {
	N  bool
	G  bool
	EX bool
	EY bool
	DE bool
}

type EnvioMatriz struct {
	Padre   string
	Cliente string
	Imagen  string
	Filtros *EnvioFiltros
}

type Solicitud struct {
	Padre   string
	Cliente string
	Imagen  string
	Filtros string
}
