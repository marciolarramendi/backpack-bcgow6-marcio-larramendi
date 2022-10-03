package main

import "fmt"

func impuestoSalario(salario float64) float64 {
	var descuento float64
	if salario > 50000 && salario < 150000 {
		descuento = salario - salario*0.83
	} else if salario > 150000 {
		descuento = salario - salario*0.73
	} else {
		descuento = 0
	}

	return descuento
}

func main() {
	sueldo := 59000.0
	fmt.Println("Sueldo: ", sueldo)
	fmt.Println("El impuesto es: ", impuestoSalario(sueldo))
}
