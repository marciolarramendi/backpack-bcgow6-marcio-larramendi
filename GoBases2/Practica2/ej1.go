/*
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle
*/

package main

import "fmt"

type estudiante struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (e estudiante) detalle() {
	fmt.Printf("Nombre: %v\nApellido: %v\nDNI: %v\nFecha: %v\n", e.Nombre, e.Apellido, e.DNI, e.Fecha)
}

func main() {
	e1 := estudiante{"Marcio", "Larramendi", 53688299, "19/09/2022"}
	e1.detalle()
}
