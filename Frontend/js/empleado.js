window.addEventListener('load',function(){
    if(!this.sessionStorage.getItem('sesionActiva')){
        this.window.location.href = 'index.html'
    }
})

function cerrarSesion() {
    sessionStorage.removeItem('sesionActiva')
    window.location.href = "index.html";
}