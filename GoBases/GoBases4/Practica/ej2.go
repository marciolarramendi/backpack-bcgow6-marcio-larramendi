package main

import (
	"errors"
	"fmt"
	"os"
)

var CustomError = errors.New("error: el salario ingresado no alcanza el minimo imponible")

func DebePagarImpuesto(salario int) error {
	if salario < 150000 {
		return CustomError
	}
	return nil
}

func main() {
	var salary int = 111000
	err := DebePagarImpuesto(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Debe pagar impuesto")

}
