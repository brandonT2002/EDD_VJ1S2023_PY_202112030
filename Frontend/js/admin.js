document.getElementById("card1").addEventListener("click", function () {
    // Obtener el elemento de entrada de archivo
    var fileInput = document.createElement("input");
    fileInput.type = "file";

    // Escuchar cambios en el archivo seleccionado
    fileInput.addEventListener("change", function (event) {
        var selectedFile = event.target.files[0];
        // Hacer algo con el archivo seleccionado (por ejemplo, cargarlo)
        console.log("Archivo seleccionado:", selectedFile);
    });

    // Hacer clic en el elemento de entrada de archivo
    fileInput.click();
});

document.getElementById("card2").addEventListener("click", function () {
    var fileInput = document.createElement("input");
    fileInput.type = "file";

    fileInput.addEventListener("change", function (event) {
        var selectedFile = event.target.files[0];
        console.log("Archivo seleccionado:", selectedFile);
    });
    fileInput.click();
});