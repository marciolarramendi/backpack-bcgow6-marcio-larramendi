package main

import (
	"fmt"
)

type MyError struct{}

func (e *MyError) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

func main() {
	var salary int = 111000
	if salary < 150000 {
		err := &MyError{}
		fmt.Println(err)
		return
	}

	fmt.Println("Debe pagar impuesto")

}
