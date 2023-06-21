package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"paquete/pedidos"
)

type Usuario struct {
	Usuario    string `json:"usuario"`
	Contrasena string `json:"contrasena"`
}

var empleados = []Usuario{
	{Usuario: "Brandon", Contrasena: "123"},
	{Usuario: "Alice", Contrasena: "456"},
	{Usuario: "Bob", Contrasena: "789"},
}

func main() {
	arbol := &pedidos.ArbolAVL{}
	arbol.Insertar(&pedidos.Pedido{IdCliente: 50, Imagen: "mario"})
	arbol.Insertar(&pedidos.Pedido{IdCliente: 17, Imagen: "luigi"})
	arbol.Insertar(&pedidos.Pedido{IdCliente: 84, Imagen: "wario"})
	arbol.Insertar(&pedidos.Pedido{IdCliente: 32, Imagen: "bowser"})
	arbol.Insertar(&pedidos.Pedido{IdCliente: 26, Imagen: "bowser"})
	arbol.Insertar(&pedidos.Pedido{IdCliente: 25, Imagen: "bowser"})
	arbol.Insertar(&pedidos.Pedido{IdCliente: 47, Imagen: "bowser"})
	arbol.Insertar(&pedidos.Pedido{IdCliente: 95, Imagen: "bowser"})
	arbol.Insertar(&pedidos.Pedido{IdCliente: 12, Imagen: "bowser"})
	arbol.Insertar(&pedidos.Pedido{IdCliente: 49, Imagen: "bowser"})
	arbol.Insertar(&pedidos.Pedido{IdCliente: 100, Imagen: "bowser"})
	arbol.Insertar(&pedidos.Pedido{IdCliente: 200, Imagen: "bowser"})
	arbol.Insertar(&pedidos.Pedido{IdCliente: 85, Imagen: "bowser"})
	arbol.Insertar(&pedidos.Pedido{IdCliente: 210, Imagen: "bowser"})
	arbol.Insertar(&pedidos.Pedido{IdCliente: 90, Imagen: "bowser"})
	fmt.Println(arbol.Dot())
	/*
		r := mux.NewRouter()
		r.HandleFunc("/", Raiz).Methods("GET")
		r.HandleFunc("/login", Login).Methods("POST", "OPTIONS")

		headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
		originsOk := handlers.AllowedOrigins([]string{"*"})
		methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

		fmt.Println("Servidor iniciado en http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headersOk, originsOk, methodsOk)(r)))
	*/
}

func Raiz(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "API en go")
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var usuario Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Nombre:", usuario.Usuario)
	fmt.Println("Contraseña:", usuario.Contrasena)

	// Verificar si el usuario existe en la lista de usuarios registrados
	existe := false
	for _, u := range empleados {
		if u.Usuario == usuario.Usuario && u.Contrasena == usuario.Contrasena {
			existe = true
			break
		}
	}

	if existe {
		response := struct {
			Message string `json:"message"`
		}{
			Message: "Datos de inicio de sesión válidos",
		}
		json.NewEncoder(w).Encode(response)
	} else {
		response := struct {
			Error string `json:"error"`
		}{
			Error: "Credenciales inválidas",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
	}
}
