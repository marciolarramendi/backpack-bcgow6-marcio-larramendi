/*
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios.
Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del programa como en las funciones.
Se necesitan las estructuras:
	- Usuario: Nombre, Apellido, Correo, Productos (array de productos).
	- Producto: Nombre, precio, cantidad.
Se requieren las funciones:
	- Nuevo producto: recibe nombre y precio, y retorna un producto.
	- Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
	- Borrar productos: recibe un usuario, borra los productos del usuario.
*/

package main

import "fmt"

type usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Producto
}

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func nuevoProducto(nombre string, precio float64) Producto {
	return Producto{
		Nombre: nombre,
		Precio: precio,
	}
}

func main() {
	//prod1 := nuevoProducto("Xiaomi Note 11 pro", 14970.3)
	//p1 := Producto{"Notebook",1500,1}
	//p2 := Producto{"Celular",500,1}
	u1 := usuario{"Pedro", "Lopez", "pedrolopez@gmail.com", Producto{"Notebook", 1500, 1}}
	fmt.Println(u1)
}
