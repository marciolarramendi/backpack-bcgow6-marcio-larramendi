package main

import "fmt"

func main() {
	// var apellido string = "Gomez", esta ok
	var apellido string = "Gomez"
	// var edad int = "35", le esta poniendo un string a un int
	var edad int = 35
	// boolean := "false"; los booleanos no llevan comilla y ademas la linea tiene ; al final que no es necesario
	boolean := false
	// var sueldo string = 45857.90, le pasa un float a una variable definida como string
	var sueldo float64 = 45857.90
	// var nombre string = "Julián", esta ok
	var nombre string = "Julián"

	fmt.Println(apellido)
	fmt.Println(edad)
	fmt.Println(boolean)
	fmt.Println(sueldo)
	fmt.Println(nombre)

}
