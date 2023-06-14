package imagencapas

import (
	"fmt"
	"os"
)

type Nodo struct {
	capa      *MatrizDispersa
	siguiente *Nodo
}

type ListaCapas struct {
	primero  *Nodo
	ultimo   *Nodo
	longitud int
}

func (this *ListaCapas) Insertar(capa *MatrizDispersa) {
	if this.primero != nil {
		this.ultimo.siguiente = &Nodo{capa: capa}
		this.ultimo = this.ultimo.siguiente
		this.longitud++
		return
	}
	this.primero = &Nodo{capa: capa}
	this.ultimo = this.primero
	this.longitud++
}

func (this *ListaCapas) GenerarImg(anchoPx, ancho, altoPx, alto int, ruta, nombre string) {
	css := `body {
	background: #333333;
	height: 100vh;
	display: flex;
	justify-content: center;
	align-items: center;
}`

	css += fmt.Sprintf(`
.canvas {
	width: %dpx;
	height: %dpx;
}`, ancho*anchoPx, alto*altoPx)

	css += fmt.Sprintf(`
.pixel {
	width: %dpx;
	height: %dpx;
	float: left;
}`, anchoPx, altoPx)

	actual := this.primero
	for actual != nil {
		css += actual.capa.ObtenerCSS(ancho)
		actual = actual.siguiente
	}

	html := fmt.Sprintf(`<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet"  href="%s.css">
	</head>
	<body>
		<div class="canvas">`, nombre)

	for i := 0; i < alto; i++ {
		for j := 0; j < ancho; j++ {
			html += "\n\t\t\t<div class=\"pixel\"></div>"
		}
	}

	html += `
		</div>
	</body>
</html>`

	actual = this.primero
	for actual != nil {
		actual.capa.GenerarGrafo(nombre)
		actual = actual.siguiente
	}

	// Crear un nuevo archivo
	file, err := os.Create(ruta + "/" + nombre + ".html")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	// Escribir el nuevo contenido en el archivo
	_, err = file.WriteString(html)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

	// Crear un nuevo archivo
	file1, err := os.Create(ruta + "/" + nombre + ".css")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	// Escribir el nuevo contenido en el archivo
	_, err = file1.WriteString(css)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

}
