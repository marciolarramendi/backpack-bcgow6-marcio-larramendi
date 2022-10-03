package main

import "fmt"

func main() {
	var temperatura float64 = 16.1
	var humedad float64 = 63.1
	var presion string = "1021mbar"

	fmt.Println("La temperatura en Uruguay es: ", temperatura, "°")
	fmt.Println("La humedad en Uruguay es: ", humedad, "%")
	fmt.Println("La presion en Uruguay es: ", presion)
}
