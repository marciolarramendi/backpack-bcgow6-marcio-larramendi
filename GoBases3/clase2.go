package main

import "fmt"

func Incrementar(puntero *int) {
	*puntero = 20
}

func main() {
	//var puntero *int
	//p2 := new(int)

	var numero int = 10
	p3 := &numero

	fmt.Printf("El valor de numero es %d\n", numero)

	Incrementar(p3)

	fmt.Printf("El valor de numero es %d\n", numero)
	//fmt.Printf("La direccion es: %p\n", p3)
	//fmt.Printf("La direccion es: %p\n", &numero)

}
