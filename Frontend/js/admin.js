api = 'http://localhost:8080'

window.addEventListener('load',function(){
    if(!this.sessionStorage.getItem('sesionActiva')){
        this.window.location.href = 'index.html'
    } else {
        var nombreUsuario = this.sessionStorage.getItem('nombreUsuario')
        var nombreElemento = this.document.getElementById('admin');
        nombreElemento.textContent = 'Bienvenido ' + nombreUsuario
    }
})

function cerrarSesion() {
    sessionStorage.removeItem('sesionActiva')
    window.location.href = "index.html";
}

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
                    Arbol()
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

function Arbol() {
    fetch(`${api}/Arbol`)
        .then(response => response.json())
        .then(data => {
            Arbol1(data)
        })
        .catch(error => {})
    }
    
function Arbol1(dot) {
    d3.select('#arbol').graphviz().scale(1).height(550*1).width(document.getElementById('arbol').clientWidth).renderDot(`${dot}`)
}

function Block() {
    fetch(`${api}/Blockchain`)
        .then(response => response.json())
        .then(data => {
            console.log(data)
            Block1(data)
        })
        .catch(error => {})
}

function Block1(dot) {
    d3.select('#block').graphviz().scale(1).height(250*1).width(document.getElementById('block').clientWidth).renderDot(`${dot}`)
}