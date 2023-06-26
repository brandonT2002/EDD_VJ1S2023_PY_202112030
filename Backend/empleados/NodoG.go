package empleados

type NodoMatrizA struct {
	Siguiente *NodoMatrizA
	Abajo     *NodoMatrizA
	valor     string
}

type EnvioMatriz struct {
	Padre   string
	Cliente string
	Imagen  string
	Filtros string
}
