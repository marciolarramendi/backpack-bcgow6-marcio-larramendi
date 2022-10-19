package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamientoAsc(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	slice := []int{5, 3, 4, 7, 8, 9}
	resultadoEsperado := []int{3, 4, 5, 7, 8, 9}

	// Se ejecuta el test
	resultado := OrdenarSliceAsc(slice)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

}
