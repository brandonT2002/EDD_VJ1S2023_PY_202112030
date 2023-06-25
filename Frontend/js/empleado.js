window.addEventListener('load',function(){
    if(!this.sessionStorage.getItem('sesionActiva')){
        this.window.location.href = 'index.html'
    } else {
        var nombreUsuario = this.sessionStorage.getItem('nombreUsuario')
        var nombreElemento = this.document.getElementById('empleado');
        nombreElemento.textContent = 'Bienvenido ' + nombreUsuario
    }
})

function cerrarSesion() {
    sessionStorage.removeItem('sesionActiva')
    window.location.href = "index.html";
}