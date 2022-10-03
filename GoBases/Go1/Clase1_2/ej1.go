package main

import "fmt"

func main() {
	palabra := "Murcielago"
	count := 0

	fmt.Println("La palabra es: ", palabra)
	fmt.Println("Tiene un largo de: ", len(palabra))

	for count < len(palabra) {
		fmt.Printf("%c\n", palabra[count])
		// Debajo dejo la forma comentada de imprimirlo mediante slice
		// fmt.Println("Letra: ", palabra[count:count+1])
		count++
	}

}
