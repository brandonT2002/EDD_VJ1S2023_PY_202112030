function login() {
    var usuario = document.getElementById('usuario').value;
    var contrasena = document.getElementById('contrasena').value;

    var loginData = {
        usuario: usuario,
        contrasena: contrasena
    };

    fetch('http://localhost:8080/login', {
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