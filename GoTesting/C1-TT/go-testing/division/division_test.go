package division

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 4
	num2 := 2
	resultadoEsperado := 2
	//denominadorZero := "El denominador no puede ser 0"
	// Se ejecuta el test
	resultado, err := Dividir(num1, num2)

	// Se validan los resultados aprovechando testify
	assert.Equal(t, resultado, resultadoEsperado, "deben ser iguales")
	assert.Nil(t, err)
}
