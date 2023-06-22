api = 'http://localhost:8080/'

function login() {
    let usuario = document.getElementById('usuario').value;
    let contrasena = document.getElementById('contrasena').value;

    if (usuario.replace(' ','') == '' || contrasena.replace(' ','') == ''){
        alert('Todos los campos son obligatorios')
    } else{
        var loginData = {
            usuario: usuario,
            contrasena: contrasena
        };
    
        fetch(api+'login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(loginData)
        })
        .then(response => response.json())
        .then(data => {
            if (data.message) {
                alert(data.message);
            } else if (data.error) {
                alert(data.error);
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });
    }
}