/* Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y cuál es el valor máximo.
*/

package main

import "fmt"

type Matrix struct {
	Alto       int
	Ancho      int
	Cuadratica bool
	Maximo     float64
	Valores    []float64
}

func (m *Matrix) setValores(valores ...float64) {
	max := valores[0]
	m.Valores = valores
	for _, valor := range valores {
		if valor > max {
			m.Maximo = valor
		}
	}
}

func (m *Matrix) setDimensiones(alto, ancho int) {
	m.Alto = alto
	m.Ancho = ancho
	if alto == ancho {
		m.Cuadratica = true
	} else {
		m.Cuadratica = false
	}
}

func (m Matrix) Print() {
	if (m.Alto * m.Ancho) > len(m.Valores) {
		cantidadDeCeros := m.Alto*m.Ancho - len(m.Valores)
		i := 0
		for i < cantidadDeCeros {
			m.Valores = append(m.Valores, 0)
			i++
		}
	}
	pos := 0
	for i := 0; i < m.Alto; i++ {
		j := 0
		for j < m.Ancho {
			fmt.Print(m.Valores[pos+j])
			j++
		}
		pos += j
		fmt.Print("\n")
	}
}

func main() {
	var m Matrix
	m.setValores(1, 2, 3, 4, 5, 6, 7, 8, 9)
	m.setDimensiones(3, 3)
	m.Print()
	fmt.Println()

	// Si tengo menos valores que el total de la dimension la autocompleto con 0
	var m2 Matrix
	m2.setValores(1, 2, 3, 4, 5, 6, 7, 8, 9)
	m2.setDimensiones(3, 5)
	m2.Print()
}
