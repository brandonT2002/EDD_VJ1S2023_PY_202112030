package pedidos

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type PedidoJSON struct {
	IdCliente string `json:"IdCliente"`
	Imagen    string `json:"Imagen"`
}

func LeerJSON(arbol *ArbolAVL, contenidoJSON string) error {
	var clientes []PedidoJSON
	err := json.Unmarshal([]byte(contenidoJSON), &clientes)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return err
	}

	for _, cliente := range clientes {
		// fmt.Println("Cliente: ", cliente.IdCliente, "Image: ", cliente.Imagen)
		id, _ := strconv.Atoi(cliente.IdCliente)
		arbol.Insertar(&Pedido{id, cliente.Imagen})
	}
	return nil
}
