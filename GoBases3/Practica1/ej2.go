package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
	contenido, err := os.ReadFile("./testej1.csv")
	if err != nil {
		log.Fatal("err")
	}

	fmt.Printf("ID \t\tPrecio \tCantidad\n")

	productos := strings.Split(string(contenido), "\n")
	//fmt.Println(values)
	var suma float64
	for i := 0; i < len(productos)-1; i++ {
		producto := strings.Split(productos[i], ",")
		fmt.Printf("%s \t\t%s \t%s\n", producto[0], producto[1], producto[2])
		precio, _ := strconv.ParseFloat(producto[1], 64)
		suma += precio
	}
	fmt.Printf("\t\t%.2f\t \n", suma)

}
