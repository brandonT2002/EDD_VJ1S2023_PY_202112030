// window.addEventListener('load',function(){
//     if(!this.sessionStorage.getItem('sesionActiva')){
//         this.window.location.href = 'index.html'
//     }
// })

function cerrarSesion(){
    // sessionStorage.removeItem('sesionActiva')
    window.location.href = "index.html";
}

document.getElementById("card1").addEventListener("click", function () {
    var fileInput = document.createElement("input");
    fileInput.type = "file";

    fileInput.addEventListener("change", function (event) {
        var selectedFile = event.target.files[0];
        var fileReader = new FileReader();
        fileReader.onload = function(e){
            var fileContent = e.target.result;
            try{
                var pedidosJson = JSON.parse(fileContent)
                console.log(pedidosJson)
            }
            catch(error){
                console.error('Error al parsear JSON:',error)
            }
        };
        fileReader.readAsText(selectedFile)
    });
    fileInput.click();
});