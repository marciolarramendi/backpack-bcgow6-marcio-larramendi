package main

import "fmt"

func main() {
	edad := 27
	sueldo := 100001.0
	empleado := true
	antiguedad := 2

	if edad >= 22 {
		if empleado {
			if antiguedad >= 1 {
				if sueldo > 100000.0 {
					fmt.Println("Hay prestamo sin interes disponible")
				} else {
					fmt.Println("Hay prestamo con interes disponible")
				}
			} else {
				fmt.Println("No tiene la antiguedad, no hay prestamo disponible")
			}
		} else {
			fmt.Println("No está empleado, no hay prestamo disponible")
		}
	} else {
		fmt.Println("Es menor de 22 años, no hay prestamo disponible")
	}
}
