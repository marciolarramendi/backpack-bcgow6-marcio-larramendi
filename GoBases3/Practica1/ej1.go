/*

Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.


*/

package main

import (
	"fmt"
	"os"
)

type Producto struct {
	ID       int
	Precio   float64
	Cantidad int
}

func guardarArchivo(productos ...Producto) {
	var textoTotal string
	//headers := fmt.Sprint("ID,Precio,Cantidad\n")
	//textoTotal += headers
	for _, valor := range productos {
		test := fmt.Sprint(valor.ID, ",", valor.Precio, ",", valor.Cantidad, "\n")
		textoTotal += test
	}

	err := os.WriteFile("./testej1.csv", []byte(textoTotal), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("File creado correctamente")
}

func main() {
	p1 := Producto{1234, 100.5, 5}
	p2 := Producto{1235, 500.5, 3}
	p3 := Producto{1236, 1000.5, 10}

	guardarArchivo(p1, p2, p3)

}
