/*Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura.
Para optimizar y ahorrar memoria requieren que la estructura de usuarios ocupe el mismo lugar en memoria para el main del programa
y para las funciones.
La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
Y deben implementarse las funciones:
	- Cambiar nombre: me permite cambiar el nombre y apellido.
	- Cambiar edad: me permite cambiar la edad.
	- Cambiar correo: me permite cambiar el correo.
	- Cambiar contraseña: me permite cambiar la contraseña.
*/

package main

import "fmt"

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

func (u *Usuario) CambiarNombre(nombre, apellido string) {
	u.Nombre = nombre
	u.Apellido = apellido
}

func (u *Usuario) CambiarEdad(edad int) {
	u.Edad = edad
}

func (u *Usuario) CambiarCorreo(correo string) {
	u.Correo = correo
}

func (u *Usuario) CambiarContraseña(contraseña string) {
	u.Contraseña = contraseña
}

func main() {
	u1 := Usuario{"Pedro", "Lopez", 25, "pedrolopez@gmail.com", "pass1234"}
	fmt.Println("Nombre: ", u1.Nombre)
	fmt.Println("Apellido: ", u1.Apellido)
	fmt.Println("Edad: ", u1.Edad)
	fmt.Println("Correo: ", u1.Correo)
	fmt.Println("Contraseña: ", u1.Contraseña)

	fmt.Println("Cambiando nombre y apellido")
	u1.CambiarNombre("Martin", "Perez")
	//fmt.Printf("El usuario ahora se llama: %s %s", u1.Nombre, u1.Apellido)

	fmt.Println("Cambiando edad")
	u1.CambiarEdad(23)
	//fmt.Printf("El usuario ahora se llama: %s %s", u1.Nombre, u1.Apellido)

	fmt.Println("Cambiando correo")
	u1.CambiarCorreo("martinperez@gmail.com")

	fmt.Println("Cambiando contraseña")
	u1.CambiarContraseña("newpass1234")

	fmt.Println("Nuevos datos:")
	fmt.Println("Nombre: ", u1.Nombre)
	fmt.Println("Apellido: ", u1.Apellido)
	fmt.Println("Edad: ", u1.Edad)
	fmt.Println("Correo: ", u1.Correo)
	fmt.Println("Contraseña: ", u1.Contraseña)
}
