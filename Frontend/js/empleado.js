window.addEventListener('load',function(){
    if(!this.sessionStorage.getItem('sesionActiva')){
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

// Obtener la fecha actual automáticamente
var fechaActual = new Date();
var formatoFecha = obtenerFormatoFecha(fechaActual);
var fecha = document.getElementById('fecha')
fecha.value = formatoFecha;
fecha.readOnly = true;

// Función para obtener la hora después de hacer clic en el botón
function obtenerHora() {
    var fechaHoraActual = new Date();
    var formatoHora = obtenerFormatoHora(fechaHoraActual);
    alert(formatoHora)
}

// Función para obtener el formato de fecha DD-MM-YYYY
function obtenerFormatoFecha(fecha) {
    var dia = padZero(fecha.getDate());
    var mes = padZero(fecha.getMonth() + 1); // Los meses van de 0 a 11, por lo que se suma 1
    var anio = fecha.getFullYear();
    return dia + '-' + mes + '-' + anio;
}

// Función para obtener el formato de hora HH:MM:SS
function obtenerFormatoHora(fecha) {
    var horas = padZero(fecha.getHours());
    var minutos = padZero(fecha.getMinutes());
    var segundos = padZero(fecha.getSeconds());
    return horas + ':' + minutos + ':' + segundos;
}

// Función auxiliar para añadir un cero delante de los números menores de 10
function padZero(numero) {
    return numero < 10 ? '0' + numero : numero;
}
