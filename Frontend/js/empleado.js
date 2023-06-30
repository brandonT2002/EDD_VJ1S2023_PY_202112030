api = 'http://localhost:8080'

window.addEventListener('load', function () {
    if (!this.sessionStorage.getItem('sesionActiva')) {
        this.window.location.href = 'index.html'
    } else {
        var nombreUsuario = this.sessionStorage.getItem('nombreUsuario')
        var nombreElemento = this.document.getElementById('empleado');
        nombreElemento.textContent = 'Bienvenido ' + nombreUsuario

        var idUsuario = this.sessionStorage.getItem('idUsuario')
        var idElemento = this.document.getElementById('idEmp')
        idElemento.value = idUsuario;
        idElemento.readOnly = true;
    }
})

function cerrarSesion() {
    sessionStorage.removeItem('sesionActiva')
    window.location.href = "index.html";
}

function pedidosCola() {
    fetch(`${api}/pedidos`)
        .then(response => response.json())
        .then(data => {
            if (data.length != 0) {
                // console.log(data)
                table = '<tr><th>Cola</th><th>ID Cliente</th><th>Imagen</th></tr>'
                data.forEach((pedido, index) => {
                    table += `<tr>
                <td>${index + 1}</td>
                <td>${pedido.IdCliente}</td>
                <td>${pedido.Imagen}</td>
                </tr>`
                })
                document.getElementById('colaPedidos').innerHTML = table
                document.getElementById('idCliente').value = data[0].IdCliente
                document.getElementById('idCliente').readOnly = true;

                document.getElementById('imgCliente').value = data[0].Imagen
                document.getElementById('imgCliente').readOnly = true;
                
            }
        })
        .catch(error => {})
}

function vender() {
    var fecha = document.getElementById('fecha').value
    var idEmp = document.getElementById('idEmp').value;
    var idCliente = document.getElementById('idCliente').value;
    var imagen = document.getElementById('imgCliente').value
    var pago = document.getElementById('pago').value;

    var negativo = document.getElementById('negativo').checked;
    var escalaGrises = document.getElementById('escala-grises').checked;
    var espejoX = document.getElementById('espejo-x').checked;
    var espejoY = document.getElementById('espejo-y').checked;
    var dobleEspejo = document.getElementById('doble-espejo').checked;

    if (pago.replace(' ','') === ''){
        alert('Ingrese un costo de imagen')
    } else {
        var filtros = {
            N:negativo,
            G:escalaGrises,
            EX:espejoX,
            EY: espejoY,
            DE: dobleEspejo
        }

        var data = {
            Padre:idEmp,
            Cliente:idCliente,
            Imagen:imagen,
            Filtros:filtros
        }

        fetch(`${api}/ventas`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })
        .then(response => response.json())
        .then(data => {
            alert(data.msg)
            facturar(fecha,idEmp,idCliente,pago)
            limpiar()
        })
        .catch(error => {})
    }
}

function solicitudes() {
    fetch(`${api}/ventas`)
        .then(response => response.json())
        .then(data => {
            if (data.length != 0) {
                table1 = '<tr><th>No.</th><th>ID Cliente</th><th>Imagen</th><th>Filtros</th></tr>'
                data.forEach((pedido, index) => {
                    table1 += `<tr>
                <td>${index + 1}</td>
                <td>${pedido.Cliente}</td>
                <td>${pedido.Imagen}</td>
                <td>${pedido.Filtros}</td>
                </tr>`
                })
                document.getElementById('solicitudes').innerHTML = table1
            }
        })
        .catch(error => {})
}

function facturar(fecha,biller,customer,payment){
    var data = {
        Timestamp:fecha,
        Biller:biller,
        Customer:customer,
        Payment:payment
    }
    fetch(`${api}/facturas`,{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => response.json())
    .then(data => {
        guardarFactura(customer,data.idFactura)
    })
    .catch(error => {})
}

function guardarFactura(cliente, factura) {
    var data = {
        IdCliente:cliente,
        IdFactura:factura
    }
    fetch(`${api}/factura`,{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => response.json())
    .then(data => {
        // console.log(data.idFactura)
    })
    .catch(error => {})
}

function facturas() {
    fetch(`${api}/factura`)
        .then(response => response.json())
        .then(data => {
            //console.log(data)
            table3 = '<tr><th>No.</th><th>ID Cliente</th><th>ID Factura</th></tr>'
            data.forEach((pedido, index) => {
                if (pedido.Llave != -1){
                    console.log(data.Llave)
                    table3 += `<tr>
                    <td>${index}</td>
                    <td>${pedido.IdCliente}</td>
                    <td>${pedido.IdFactura}</td>
                    </tr>`
                }
            })
            document.getElementById('facturas').innerHTML = table3
        })
        .catch(error => {})
}

function limpiar(){
    window.location.href = 'Empleado.html#close'
    
    document.getElementById('fecha').value = '';
    document.getElementById('idCliente').value = '';
    document.getElementById('imgCliente').value = '';
    document.getElementById('pago').value = '';
    
    document.getElementById('negativo').checked = false;
    document.getElementById('escala-grises').checked = false;
    document.getElementById('espejo-x').checked = false;
    document.getElementById('espejo-y').checked = false;
    document.getElementById('doble-espejo').checked = false;

    pedidosCola()
    solicitudes()
    facturas()
    verGrafo()
}

// grafo 
function verGrafo(){
    fetch(`${api}/empleado`)
        .then(response => response.json())
        .then(data => {
            grafo(data)
        })
        .catch(error => {})
}

function grafo(dot) {
    d3.select('#grafo').graphviz().scale(0.5).height(675*1).width(document.getElementById('grafo').clientWidth).renderDot(`${dot}`)
}

function updateDateTime() {
    var fechaActual = new Date();

    var dia = fechaActual.getDate();
    var mes = fechaActual.getMonth() + 1;
    var anio = fechaActual.getFullYear();

    var horas = fechaActual.getHours();
    var minutos = fechaActual.getMinutes();
    var segundos = fechaActual.getSeconds();

    dia = padZero(dia);
    mes = padZero(mes);
    horas = padZero(horas);
    minutos = padZero(minutos);
    segundos = padZero(segundos);

    var formatoFechaHora = dia + '-' + mes + '-' + anio + '-::' + horas + ':' + minutos + ':' + segundos;
    document.getElementById('fecha').value = formatoFechaHora;
}

function padZero(numero) {
    return numero < 10 ? '0' + numero : numero;
}

// Update the date and time immediately
updateDateTime();

// Update the date and time every second (1000 milliseconds)
setInterval(updateDateTime, 1000);

function cambiarImagen(nuevaImagen) {
    var imgElement = document.getElementById('img1');
    imgElement.src = nuevaImagen;
}
