package main

import "fmt"

var (
	pequeño = "pequeño"
	mediano = "mediano"
	grande  = "grande"
)

type tienda struct {
	Productos []producto
}

type producto struct {
	Tipo   string
	Nombre string
	Precio float64
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(producto)
}

func (p producto) CalcularCosto() float64 {
	switch p.Tipo {
	case mediano:
		return p.Precio + p.Precio*0.03
	case grande:
		return p.Precio + p.Precio*0.06 + 2500
	default:
		return p.Precio
	}
}

func (t tienda) Total() float64 {
	var total float64
	for _, prod := range t.Productos {
		total += prod.CalcularCosto()
	}
	return total
}

func (t *tienda) Agregar(p producto) {
	t.Productos = append(t.Productos, p)
}

func nuevoProducto(tipo string, nombre string, precio float64) Producto {
	return producto{
		Tipo:   tipo,
		Nombre: nombre,
		Precio: precio,
	}
}

func nuevaTienda() Ecommerce {
	return &tienda{}
}

func main() {
	// Precios sacados de la tienda oficial Xiaomi Uruguay con el dolar a 41.7
	prod1 := nuevoProducto(pequeño, "Xiaomi Note 11 pro", 14970.3)
	prod2 := nuevoProducto(mediano, "Freidora sin aceite Xiaomi", 5379.3)
	prod3 := nuevoProducto(grande, "Mi LED TV 4S 4K 65''", 49998.3)
	tienda := nuevaTienda()

	tienda.Agregar(prod1.(producto))
	tienda.Agregar(prod2.(producto))
	tienda.Agregar(prod3.(producto))
	fmt.Println(tienda)

	fmt.Println("Total: $", tienda.Total())
}
