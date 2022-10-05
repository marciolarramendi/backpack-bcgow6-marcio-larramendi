package users

import (
	"fmt"
)

type Usuario struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre" binding:"required"`
	Apellido      string `json:"apellido" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Edad          int    `json:"edad" binding:"required"`
	Altura        int    `json:"altura" binding:"required"`
	Activo        bool   `json:"activo" binding:"required"`
	FechaCreacion string `json:"fechacreacion" binding:"required"`
}

var usuarios []Usuario
var lastID int

// ***Importado**//
type Repository interface {
	GetAll() ([]Usuario, error)
	Store(id int, nombre, apellido, email string, edad, altura int, activo bool, fechacreacion string) (Usuario, error)
	LastID() (int, error)
	Update(id int, nombre, apellido, email string, edad, altura int, activo bool, fechacreacion string) (Usuario, error)
	PartialUpdate(id int, apellido string, edad int) (Usuario, error)
	Delete(id int) error
}

type repository struct{} //struct implementa los metodos de la interfaz

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Store(id int, nombre, apellido, email string, edad, altura int, activo bool, fechacreacion string) (Usuario, error) {
	u := Usuario{id, nombre, apellido, email, edad, altura, activo, fechacreacion}
	usuarios = append(usuarios, u)
	lastID = u.ID
	return u, nil
}

func (r *repository) GetAll() ([]Usuario, error) {
	return usuarios, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Update(id int, nombre, apellido, email string, edad, altura int, activo bool, fechacreacion string) (Usuario, error) {
	u := Usuario{Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, FechaCreacion: fechacreacion}
	updated := false
	for i := range usuarios {
		if usuarios[i].ID == id {
			u.ID = id
			usuarios[i] = u
			updated = true
		}
	}
	if !updated {
		return Usuario{}, fmt.Errorf("Usuario %d no encontrado", id)
	}
	return u, nil
}

func (r *repository) PartialUpdate(id int, apellido string, edad int) (Usuario, error) {
	updated := false
	var index int
	for i := range usuarios {
		if usuarios[i].ID == id {
			updated = true
			index = i
			usuarios[i].Apellido = apellido
			usuarios[i].Edad = edad
		}
	}
	if !updated {
		return Usuario{}, fmt.Errorf("Usuario %d no encontrado", id)
	}
	return usuarios[index], nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range usuarios {
		if usuarios[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Usuario %d no encontrado", id)
	}
	usuarios = append(usuarios[:index], usuarios[index+1:]...)
	return nil
}
