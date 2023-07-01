# **PROYECTO ESTRUCTURAS DE DATOS - FASE 2**

Brandon Andy Jefferson Tejaxún Pichiyá - 202112030


## 🎯 **Objetivos**

### **General**

Aplicar los conocimientos del curso de Estructuras de Datos en el desarrollo de diferentes estructuras de datos y los diferentes algoritmos de manipulación de información en ellas.

### **Específicos**
* Utilizar el lenguaje Go para implementar estructuras de datos no lineales.
* Utilizar la herramienta Graphviz para graficar las estructuras de datos.
* Definir e implementar algoritmos de ordenamiento, búsqueda e inserción en las diferentes estructuras a implementar.

## 📌 **Manual de Usuario**

### 1. **Administrador**

1. Inicio de Sesión

    Se debe ingresar las credenciales válidas para el administrador para poder acceder al panel correspondiente.

    <p align="center">
        <img src="Image/1.png" width="575px">
    </p>

    Se muestran cada una de las funcionalidades con las que cuenta el administrador y los reportes a los que tiene acceso.

    <p align="center">
        <img src="Image/2.png" width="575px">
    </p>

<br>

2. Cargar Pedidos

    Se debe dar clic sobre la opción correspondiente y seleccionar el archivo en formato **.json** 

<br>

3. Cargar Empleados

    Sedebe dar clic sobre la opción correspondiente y seleccionar el archivo

<br>

4. Reportes

    Al cargar el archivo de pedidos, automáticamente se mostrarán en una estructura de **Árbol AVL**, al igual que cada una de las ventas serán registradas y se mostrará el reporte de, incluyendo datos relevantes como la hora, empleado que atendió el pedido, cliente, entre otros.

    <p align="center">
        <img src="Image/3.png" width="575px">
    </p>


### 2. **Empleado**

1. Inicio de Sesión

    Se debe ingresar las credenciales válidas para un empleado, para poder acceder al panel correspondiente.

    <p align="center">
        <img src="Image/4.png" width="575px">
    </p>

    Se muestran cada una de las funcionalidades con las que cuenta el administrador y los reportes a los que tiene acceso.

    <p align="center">
        <img src="Image/6.png" width="575px">
    </p>

<br>

2. Aplicar Filtros

    Esta funcionalidad permite al empleado generar la imagen que el cliente ha solicitado aplicando filtros según se requieral.

    Filtros:
    * Negativo
    * Escala de Grises
    * Espejo en X
    * Espejo en Y
    * Espejo Doble

    Al momentode antender un pedido, inicialmente se mostrará una previsualización simbólica de la imagen original y cada uno de los filtros según se vayan seleccionando.

    <p align="center">
        <img src="Image/7.png" width="500px">
        <img src="Image/8.png" width="500px">
    </p>

<br>

3. Reportes

    Cada venta realizada será asociado al Empleado que atendió el pedido, y se mostrará una tabla con los datos de cada venta, al igual que se mostrarán los datos de facturación en una tabla.

    <p align="center">
        <img src="Image/9.png" width="575px">
    </p>

    Así mismo cada venta realizada se mostrará en un grafo asociado a cada empleado.

    <p align="center">
        <img src="Image/10.png" width="575px">
    </p>

<br>


## 📌 **Manual Técinico**

### 1. **Estructuras**

1. Árbol AVL

   Se utilizó para el manejo y control de pedidos realizados, el método **Insertar** recibe un objeto de tipo Pedido, obteniedo el atributo *IdCliente* para realizar las validaciones y cálculos para balancear el arbol.

   ```go
   type ArbolAVL struct {
	raiz *Nodo
    }

    func (a *ArbolAVL) Insertar(pedido *Pedido) {
        a.raiz = a.Insertar_(pedido, a.raiz)
    }

    func (a *ArbolAVL) Insertar_(pedido *Pedido, nodo *Nodo) *Nodo {
        if nodo == nil {
            return &Nodo{
                pedido: pedido,
                izq:    nil,
                der:    nil,
            }
        }
        if pedido.IdCliente < nodo.pedido.IdCliente {
            nodo.izq = a.Insertar_(pedido, nodo.izq)
            if a.getAltura(nodo.izq)-a.getAltura(nodo.der) == 2 {
                if pedido.IdCliente < nodo.izq.pedido.IdCliente {
                    nodo = a.rotarIzq(nodo)
                } else {
                    nodo = a.rotarDobleIzq(nodo)
                }
            }
        } else if pedido.IdCliente > nodo.pedido.IdCliente {
            nodo.der = a.Insertar_(pedido, nodo.der)
            if a.getAltura(nodo.der)-a.getAltura(nodo.izq) == 2 {
                if pedido.IdCliente > nodo.der.pedido.IdCliente {
                    nodo = a.rotarDer(nodo)
                } else {
                    nodo = a.rotarDobleDer(nodo)
                }
            }
        }
        nodo.altura = a.alturaMax(a.getAltura(nodo.izq), a.getAltura(nodo.der)) + 1
        return nodo
    }

    func (a *ArbolAVL) alturaMax(izq, der int) int {
        if izq > der {
            return izq
        }
        return der
    }

    func (a *ArbolAVL) getAltura(nodo *Nodo) int {
        if nodo == nil {
            return -1
        }
        return nodo.altura
    }

    func (a *ArbolAVL) rotarIzq(nodo *Nodo) *Nodo {
        auxiliar := nodo.izq
        nodo.izq = auxiliar.der
        auxiliar.der = nodo
        nodo.altura = a.alturaMax(a.getAltura(nodo.izq), a.getAltura(nodo.der)) + 1
        auxiliar.altura = a.alturaMax(a.getAltura(auxiliar.izq), nodo.altura) + 1
        return auxiliar
    }

    func (a *ArbolAVL) rotarDer(nodo *Nodo) *Nodo {
        auxiliar := nodo.der
        nodo.der = auxiliar.izq
        auxiliar.izq = nodo
        nodo.altura = a.alturaMax(a.getAltura(nodo.izq), a.getAltura(nodo.der)) + 1
        auxiliar.altura = a.alturaMax(a.getAltura(auxiliar.der), nodo.altura) + 1
        return auxiliar
    }

    func (a *ArbolAVL) rotarDobleIzq(nodo *Nodo) *Nodo {
        nodo.izq = a.rotarDer(nodo.izq)
        return a.rotarIzq(nodo)
    }

    func (a *ArbolAVL) rotarDobleDer(nodo *Nodo) *Nodo {
        nodo.der = a.rotarIzq(nodo.der)
        return a.rotarDer(nodo)
    }
   ```

2. Grafo No Dirigido

    Se utilizó para tener una bitácora de los pedidos atendidos por cada empleado. El método **InsertarValores** recibe como parámetro un objeto de tipo *EnvioMatriz*, el cuál tiene los siguientes atributos:

    * Padre
    * Cliente
    * Imagen
    * Filtros

    NOTA: El atributo filtros es de tipo *EnvioFiltros*, el cuál tiene propiedades de cada filtro que es posible aplicar en el sistema y es de tipo booleano.

    ```go
    type Grafo struct {
	    Principal *NodoMatrizA
    }

    func (g *Grafo) InsertarC(padre, hijo string, filtro string) {
        nuevoNodo := &NodoMatrizA{valor: hijo}
        if g.Principal != nil && padre == g.Principal.valor {
            aux := g.Principal
            for aux.Siguiente != nil {
                aux = aux.Siguiente
            }
            aux.Siguiente = nuevoNodo
        } else {
            g.insertarF(padre)
            aux := g.Principal
            for aux != nil {
                if aux.valor == padre {
                    break
                }
                aux = aux.Abajo
            }
            if aux != nil {
                nuevoNodo.Siguiente = &NodoMatrizA{valor: filtro}
                aux.Siguiente = nuevoNodo
            }
        }
    }

    func (g *Grafo) insertarF(padre string) {
        nuevoNodo := &NodoMatrizA{valor: padre}
        if g.Principal == nil {
            g.Principal = nuevoNodo
        } else {
            aux := g.Principal
            for aux.Abajo != nil {
                if aux.valor == padre {
                    return
                }
                aux = aux.Abajo
            }
            aux.Abajo = nuevoNodo
        }
    }

    func (g *Grafo) InsertarValores(envio *EnvioMatriz) {
        fs := g.obtenerFiltros(envio.Filtros)
        if g.Principal == nil {
            g.insertarF(envio.Padre)
            g.InsertarC(envio.Padre, envio.Imagen, "")
            g.InsertarC(envio.Cliente, envio.Imagen, fs)
        } else {
            g.InsertarC(envio.Padre, envio.Cliente, "")
            g.InsertarC(envio.Cliente, envio.Imagen, fs)
        }
    }

    func (g *Grafo) obtenerFiltros(filtros *EnvioFiltros) string {
        fs := ""
        if filtros.N {
            fs += "Negativo"
        }
        if filtros.G {
            if fs != "" {
                fs += " - "
            }
            fs += "Grises"
        }
        if filtros.EX {
            if fs != "" {
                fs += " - "
            }
            fs += "Espejo X"
        }
        if filtros.EY {
            if fs != "" {
                fs += " - "
            }
            fs += "Espejo Y"
        }
        if filtros.DE {
            if fs != "" {
                fs += " - "
            }
            fs += "Doble Espejo"
        }
        if fs == "" {
            fs += "Sin Filtros"
        }
        return fs
    }
    ```

3. Tabla Hash

    Se utilizó para el almacenamiento de los datos de facturación.

    ```go
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
    ```

4. Blockchain

    Se utilizó para el control de pagos, la función **Insertar** recibe como parámetro un objeto de tipo Petición, el cual tiene atributos que son utlizados al momento de crear y almacenar datos en la estructura *map* que a su vez es insertado en un bloque de la Cadena de Bloques (Blockchain). La función retorna un valor de tipo string que contiene el hash del bloque que será asociado a la factura en la estructura de **Tabla Hash**.

    ```go
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
    ```