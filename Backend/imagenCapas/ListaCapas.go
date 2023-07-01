package imagencapas

import (
	"fmt"
	"os"
	"paquete/empleados"
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

func (l *ListaCapas) Insertar(capa *MatrizDispersa) {
	if l.primero != nil {
		l.ultimo.siguiente = &Nodo{capa: capa}
		l.ultimo = l.ultimo.siguiente
		l.longitud++
		return
	}
	l.primero = &Nodo{capa: capa}
	l.ultimo = l.primero
	l.longitud++
}

func (l *ListaCapas) GenerarImg(anchoPx, ancho, altoPx, alto int, ruta, nombre string, filtros *empleados.EnvioFiltros) {
	nombreFiltro := ""
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
	`, anchoPx, altoPx)

	espejo := 3

	if filtros != nil {
		if filtros.N {
			nombreFiltro = "Negativo"
			css += "filter: invert();"
		} else if filtros.G {
			nombreFiltro = "Grises"
			css += "filter: grayscale();"
		} else if filtros.EX {
			nombreFiltro = "EspejoX"
			espejo = 0
		} else if filtros.EY {
			nombreFiltro = "EspejoY"
			espejo = 1
		} else if filtros.DE {
			nombreFiltro = "EspejoD"
			espejo = 2
		}
	}
	css += "}"
	actual := l.primero
	for actual != nil {
		css += actual.capa.ObtenerCSS(ancho, alto, espejo)
		actual = actual.siguiente
	}

	estilos := ""
	if filtros == nil {
		estilos = fmt.Sprintf(`<link rel="stylesheet"  href="%s.css">`, nombre)
	} else {
		estilos = fmt.Sprintf(`<link rel="stylesheet"  href="%s_Filtro_%s.css">`, nombre, nombreFiltro)
	}

	html := fmt.Sprintf(`<!DOCTYPE html>
<html>
	<head>
		%s
	</head>
	<body>
		<div class="canvas">`, estilos)

	for i := 0; i < alto; i++ {
		for j := 0; j < ancho; j++ {
			html += "\n\t\t\t<div class=\"pixel\"></div>"
		}
	}

	html += `
		</div>
	</body>
</html>`

	if filtros == nil {
		file, err := os.Create(ruta + "/" + nombre + ".html")
		if err != nil {
			fmt.Println("Error al crear el archivo:", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(html)
		if err != nil {
			fmt.Println("Error al escribir en el archivo:", err)
			return
		}

		file1, err := os.Create(ruta + "/" + nombre + ".css")
		if err != nil {
			fmt.Println("Error al crear el archivo:", err)
			return
		}
		defer file1.Close()

		_, err = file1.WriteString(css)
		if err != nil {
			fmt.Println("Error al escribir en el archivo:", err)
			return
		}
	} else {
		file, err := os.Create(ruta + "/" + nombre + "_Filtro_" + nombreFiltro + ".html")
		if err != nil {
			fmt.Println("Error al crear el archivo:", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(html)
		if err != nil {
			fmt.Println("Error al escribir en el archivo:", err)
			return
		}

		file1, err := os.Create(ruta + "/" + nombre + "_Filtro_" + nombreFiltro + ".css")
		if err != nil {
			fmt.Println("Error al crear el archivo:", err)
			return
		}
		defer file1.Close()

		_, err = file1.WriteString(css)
		if err != nil {
			fmt.Println("Error al escribir en el archivo:", err)
			return
		}
	}
}
