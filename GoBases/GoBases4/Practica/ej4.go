package main

import (
	"errors"
	"fmt"
	"os"
)

var pocasHoras = errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")

func calcularSalario(horas, valor float64) (salarioCalculado float64, err error) {
	if horas < 80 || horas < 0 {
		err = pocasHoras
		return
	}

	salarioCalculado = horas * valor

	if salarioCalculado >= 150000 {
		salarioCalculado = salarioCalculado * 0.90
	}
	return
}

func calcularMedioAguinaldo(mejorSalario, meses float64) (medioAguinaldo float64, err error) {
	if mejorSalario < 0 {
		err = fmt.Errorf("el salario ingresado es: %.f", mejorSalario)
		return
	}

	if meses < 0 {
		err = fmt.Errorf("los meses ingresados son: %.f", meses)
		return
	}

	medioAguinaldo = mejorSalario / 12 * meses
	return
}

func main() {
	salario, err := calcularSalario(360, 2000)
	if err != nil {
		fmt.Printf("Ha ocurrido un error inesperado: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("El salario es: %v\n", salario)

	aguinaldo, err2 := calcularMedioAguinaldo(648000, 5)
	if err2 != nil {
		err3 := fmt.Errorf("Ha ocurrido un error inesperado: %v", err2)
		fmt.Println(err3)
		os.Exit(1)
	}
	fmt.Printf("El aguinaldo es: %v\n", aguinaldo)
}