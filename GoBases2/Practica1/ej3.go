package main

import "fmt"

func calcularSalario(minutos float64, categoria string) float64 {
	horas := minutos / 60
	var salario float64
	switch categoria {
	case "C":
		salario = 1000 * horas
	case "B":
		salario = (1500 * horas) * 1.20
	case "A":
		salario = (3000 * horas) * 1.50
	}
	return salario
}

func main() {
	fmt.Println(calcularSalario(2400, "A")) //Equivalente a 40 horas semanales
	fmt.Println(calcularSalario(2400, "B"))
	fmt.Println(calcularSalario(2400, "C"))
}
