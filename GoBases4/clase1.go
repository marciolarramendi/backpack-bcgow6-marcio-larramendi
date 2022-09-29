package main

import (
	"errors"
	"fmt"
	"os"
)

func division(num1, num2 float64) (result float64, err error) {
	if num2 == 0 {
		err = errors.New("no se puede dividir por 0")
		return
	}

	result = num1 / num2
	return
}

func main() {
	result, err := division(10, 2)
	if err != nil {
		fmt.Printf("Ha ocurrido un error inesperado: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("El resultado es: %v\n", result)
	fmt.Printf("El error es: %v", err)

}

/*
// Ahora veremos unwrap,
// que basicamente sirve para mostrar el error contenido por el error principal
type ErrType2 struct{}

func (e ErrType2) Error() string {
	return "soy el error 2"
}

func main() {
	err2 := ErrType2{}
	err1 := fmt.Errorf("Soy el error #1, contengo al error #2: %w", err2)
	err := fmt.Errorf("Soy el error, contengo al error #1: %w", err1)

	fmt.Println(errors.Unwrap(err))

	fmt.Println(errors.Unwrap(err1))

	fmt.Println(
		errors.Unwrap(err2),
	)
}
*/

//Usando metodo Is para comparar
/*
var ErrNotFound = errors.New("not found")

func find() (result string, err error) {
	return "", ErrNotFound
}

func main() {
	_, err := find()
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			fmt.Println("No encontrado")
		} else {
			fmt.Println("Error interno")
		}
	}
}
*/

//Aca creamos error y usamos AS
/*
type MyCustomError struct {
	StatusCode int
	Message    string
}

func (err *MyCustomError) Error() string {
	return fmt.Sprintf(
		"%s (%d)",
		err.Message,
		err.StatusCode,
	)
}

func main() {
	err1 := &MyCustomError{
		StatusCode: 502,
		Message:    "Soy el error #1, soy distinto al #2",
	}
	err2 := &MyCustomError{
		StatusCode: 406,
		Message:    "Soy el error #2, soy distinto al #1",
	}

	bothErrorsAreEqual := errors.As(err1, &err2)

	fmt.Println(bothErrorsAreEqual)

}
*/

//ERror personalizado sin crear estructura
/*
func main() {
	statusCode := 200

	if statusCode > 399 {
		err := errors.New("la peticion ha fallado")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Status %d, funciona! :D", statusCode)
}
*/

//ACA creamos un tipo de error personalizado y lo retornamos
/*
type MyCustomError struct {
	StatusCode int
	Message    string
}

func (err *MyCustomError) Error() string {
	return fmt.Sprintf(
		"%s (%d)",
		err.Message,
		err.StatusCode,
	)
}

func ObtainError(status int) (code int, err error) {
	if status >= 400 {
		return 500, &MyCustomError{
			StatusCode: 500,
			Message:    "Algo salió mal :(",
		}
	}

	return 200, nil
}

func main() {
	status, err := ObtainError(436)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Status %d, funciona! :D", status)
}
*/

//Aca mostramos errores de prueba
/*
func main() {
	err1 := fmt.Errorf("first error :(")

	err2 := fmt.Errorf("Second error: %w", err1)

	fmt.Println(err1)
	fmt.Println(err2)
}
*/
