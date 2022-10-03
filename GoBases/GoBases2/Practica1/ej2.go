package main

import (
	"errors"
	"fmt"
)

func promedio(notas ...int) (int, error) {
	var suma int
	for _, value := range notas {
		if value < 0 {
			return 0, errors.New("Ha ingresado un numero negativo")
		} else {
			suma += value
		}
	}
	return suma / len(notas), nil
}

func main() {
	res, err := promedio(-9, 8, 7, 10, 11)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
