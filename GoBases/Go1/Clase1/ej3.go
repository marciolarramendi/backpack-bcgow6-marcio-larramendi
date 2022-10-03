package main

import "fmt"

func main() {
	// var 1nombre string, hay que sacarle el numero del principio
	var nombre string
	// var apellido string, esta bien
	var apellido string
	// var int edad, esta mal el orden
	var edad int
	// 1apellido := 6, aca hay que sacarle el numero y ademas le esta asignando un valor de int cuando es un string
	apellido = "6"
	// var licencia_de_conducir = true, esta mal nombrada
	var licenciaDeConducir = true
	// var estatura de la persona int, esta mal nombrada
	var estaturaDeLaPersona int
	// cantidadDeHijos := 2, esta bien
	cantidadDeHijos := 2

	fmt.Println(nombre)
	fmt.Println(apellido)
	fmt.Println(edad)
	fmt.Println(licenciaDeConducir)
	fmt.Println(estaturaDeLaPersona)
	fmt.Println(cantidadDeHijos)
}
