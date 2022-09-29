package main

import (
	"fmt"
	"os"
	"strings"
)

/*func main() {
	fileContentAsBytes, err := os.ReadFile("./testej1.csv")
	if err != nil {
		panic(err)
	}

	//Asi se ve el contenido del file cambiandolo a string
	fmt.Println(string(fileContentAsBytes))
}*/

func main() {

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

}
