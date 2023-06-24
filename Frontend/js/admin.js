api = 'http://localhost:8080'

/*
window.addEventListener('load',function(){
    if(!this.sessionStorage.getItem('sesionActiva')){
        this.window.location.href = 'index.html'
    }
})

function cerrarSesion() {
    sessionStorage.removeItem('sesionActiva')
    window.location.href = "index.html";
}
*/

document.getElementById("card1").addEventListener("click", function () {
    var fileInput = document.createElement("input");
    fileInput.type = "file";

    fileInput.addEventListener("change", function (event) {
        var selectedFile = event.target.files[0];
        var fileReader = new FileReader();
        fileReader.onload = function (e) {
            var fileContent = e.target.result;

            var pedido = JSON.parse(fileContent);

            fetch(`${api}/pedidos`,{
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(pedido)
            })
            .then(response => response.json())
            .then(data => {
                if (data.msg === 'ok'){
                    alert('Archivo cargado exitosamente')
                }else{
                    alert('Ocurrio un error')
                }
            })
            .catch(error => {
                console.log(error)
                alert('Ocurrio un error en el servidor')
            })

        };
        fileReader.readAsText(selectedFile)
    });
    fileInput.click();
});

document.getElementById('card2').addEventListener('click', function () {
    var fileInput = document.createElement("input");
    fileInput.type = "file";

    fileInput.addEventListener("change", function (event) {
        var selectedFile = event.target.files[0];
        var fileReader = new FileReader();
        fileReader.onload = function (e) {
            var fileContent = e.target.result;
            var empleado = {
                Credenciales: fileContent
            };

            fetch(`${api}/empleado`,{
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(empleado)
            })
            .then(response => response.json())
            .then(data => {
                if (data.msg === 'ok'){
                    alert('Archivo cargado exitosamente')
                }else{
                    alert('Ocurrio un error')
                }
            })
            .catch(error => {
                alert('Ocurrio un error en el servidor')
            })

        };
        fileReader.readAsText(selectedFile)
    });
    fileInput.click();
});