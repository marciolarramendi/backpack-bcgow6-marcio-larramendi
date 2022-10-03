package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	count := 0

	// Imprimiendo map inicial
	fmt.Println("Mapa inicial")
	fmt.Println(employees)

	//Imprimiendo edad de Benjamin
	fmt.Println("La edad de Benjamin es: ", employees["Benjamin"])

	// Recorriendo el map para ver cuantos son mayores de 21
	for employee := range employees {
		if employees[employee] > 21 {
			count++
		}
	}
	fmt.Println("Hay", count, "empleados mayores de 21 años")

	// Agregando a Federico de 25 años
	employees["Federico"] = 25

	// Imprimiendo eployees entero
	fmt.Println("Se ha añadido a Federico a la lista")
	fmt.Println(employees)

	// Borrando a Pedro
	delete(employees, "Pedro")

	// Imprimiendo eployees entero
	fmt.Println("Se ha eliminado a Pedro a la lista")
	fmt.Println(employees)
}
