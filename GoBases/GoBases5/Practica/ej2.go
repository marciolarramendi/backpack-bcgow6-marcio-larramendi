package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var Legajo int = verificarUltimoLegajo()
var archivoVacio = errors.New("el archivo está vacío")

type Cliente struct {
	Legajo         int
	NombreCompleto string
	DNI            int
	Telefono       int
	Domicilio      string
}

func leerFile(nombreFile string) ([]byte, error) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	contenido, err := os.ReadFile(nombreFile)

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	return contenido, nil
}

func verificarUltimoLegajo() (int, error) {
	var legajo int
	contenido, err := leerFile("./customers.txt")
	if err != nil {
		panic("Ha ocurrido un problema inesperado, abortando el programa")
	}
	customers := strings.Split(string(contenido), "\n")
	if customers[0] != "" {
		for i := 0; i < len(customers)-1; i++ {
			empleado := strings.Split(customers[i], ",")
			ultimoValor, _ := strconv.ParseInt(empleado[4], 0, 64)
			legajo = int(ultimoValor)
		}
		legajo++
		return legajo, nil
	} else {
		legajo = 0
		return legajo, archivoVacio
	}
}

func generarLegajo() (int, error) {
	ultimoLegajo, err := verificarUltimoLegajo()
	ultimoLegajo++
	return ultimoLegajo, err
}

/*func guardarArchivo(productos ...Producto) {
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
}*/

func main() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
			fmt.Println("Soy yo el que imprime")
		}
	}()

	legajo, err := generarLegajo()
	if err != nil {
		panic("Ha ocurrido un error al generar el legajo")
	}
	fmt.Println("El legajo es: ", legajo)

	fmt.Println("Fin de la ejecución")

	/*

		contenido, err := os.ReadFile("./empleados.txt")
		if err != nil {
			defer func() {
				fmt.Println("ejecución finalizada")
			}()
			panic("el archivo indicado no fue encontrado o está dañado")
		}

		fmt.Printf("Nombre y Apellido \tDNI \tNumero de Telefono \tDomicilio \tLegajo\n")

		empleados := strings.Split(string(contenido), "\n")
		//fmt.Println(values)
		for i := 0; i < len(empleados)-1; i++ {
			empleado := strings.Split(empleados[i], ",")
			fmt.Printf("%s \t%s \t%s \t%s \t%s\n", empleado[0], empleado[1], empleado[2], empleado[3], empleado[4])
			//precio, _ := strconv.ParseFloat(empleado[1], 64)
		}
	*/

}
