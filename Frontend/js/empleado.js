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

function verPedidos() {
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
            } 
            document.getElementById('pagar').disabled = true
        })
        .catch(error => {
            console.log(error)
        })
}

function vender() {
    var fecha = document.getElementById('fecha').value;
    var idEmp = document.getElementById('idEmp').value;
    var idCliente = document.getElementById('idCliente').value;
    var pago = document.getElementById('pago').value;

    var negativo = document.getElementById('negativo').checked;
    var escalaGrises = document.getElementById('escala-grises').checked;
    var espejoX = document.getElementById('espejo-x').checked;
    var espejoY = document.getElementById('espejo-y').checked;
    var dobleEspejo = document.getElementById('doble-espejo').checked;

    if (pago.replace(' ','') === ''){
        alert('Ingrese un costo de imagen')
    } else {
        console.log('Fecha:', fecha);
        console.log('ID Empleado:', idEmp);
        console.log('ID Cliente:', idCliente);
        console.log('Pago:', pago);
        console.log('Negativo:', negativo);
        console.log('Escala de grises:', escalaGrises);
        console.log('Espejo X:', espejoX);
        console.log('Espejo Y:', espejoY);
        console.log('Doble espejo:', dobleEspejo);
    }
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

    var formatoFechaHora = dia + '-' + mes + '-' + anio + '   ' + horas + ':' + minutos + ':' + segundos;
    document.getElementById('fecha').value = formatoFechaHora;
}

function padZero(numero) {
    return numero < 10 ? '0' + numero : numero;
}

// Update the date and time immediately
updateDateTime();

// Update the date and time every second (1000 milliseconds)
setInterval(updateDateTime, 1000);