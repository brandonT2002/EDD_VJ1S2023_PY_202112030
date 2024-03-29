api = 'http://localhost:8080'

function iniciarSesion() {
    let usuario = document.getElementById('usuario').value;
    let contrasena = document.getElementById('contrasena').value;

    if (usuario.replace(' ','') === '' || contrasena.replace(' ','') === ''){
        alert('Todos los campos son obligatorios')
    } else {
        var loginData = {
            Usuario: usuario,
            Contrasena: contrasena
        };

        fetch(`${api}/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(loginData)
        })
        .then(response => response.json())
        .then(data => {
            if (data.msg === 'admin') {
                sessionStorage.setItem('sesionActiva', 'true')
                sessionStorage.setItem('nombreUsuario',data.usuario)
                window.location.href = 'Admin.html'
            } else if (data.msg === 'emp') {
                sessionStorage.setItem('sesionActiva','true')
                sessionStorage.setItem('nombreUsuario',data.usuario.Nombre)
                sessionStorage.setItem('idUsuario',data.usuario.Id)
                window.location.href = 'Empleado.html'
            } else {
                alert('Credenciales inválidas')
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });
    }
}