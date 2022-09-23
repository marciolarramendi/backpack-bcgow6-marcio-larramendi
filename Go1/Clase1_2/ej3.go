package main

import "fmt"

func main() {
	mes := 7
	var meses = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Setiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}

	fmt.Println(meses[mes])

	// Otra opcion podria ser escribir un switch case o ifs
}
