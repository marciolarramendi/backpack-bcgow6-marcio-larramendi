package users

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	mockedData []Usuario
}

func (s StubStore) Read(data interface{}) error {
	a, ok := data.(*[]Usuario)
	if !ok {
		return errors.New("Falló")
	}
	*a = s.mockedData
	return nil
}

func (s StubStore) Write(data interface{}) error {
	return nil
}

type MockStore struct {
	readWasCalled bool
	mockedData    []Usuario
}

func (m *MockStore) Read(data interface{}) error {
	m.readWasCalled = true
	users := data.(*[]Usuario)
	*users = m.mockedData
	return nil
}
func (m MockStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	expected := []Usuario{
		{
			ID:            1,
			Nombre:        "Marcelo",
			Apellido:      "Lopez",
			Email:         "marcelopez@gmail.com",
			Edad:          27,
			Altura:        165,
			Activo:        true,
			FechaCreacion: "05/10/2022",
		},
		{
			ID:            2,
			Nombre:        "Martin",
			Apellido:      "Lopez",
			Email:         "martinlopez@gmail.com",
			Edad:          27,
			Altura:        166,
			Activo:        true,
			FechaCreacion: "06/10/2022",
		},
	}
	db := &StubStore{mockedData: expected}
	repo := NewRepository(db)

	result, err := repo.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestUpdateLastNameAndAge(t *testing.T) {
	ExpectedApellido := "After to update"
	ExpectedEdad := 45
	usuarioRegistrado := []Usuario{
		{
			ID:            1,
			Nombre:        "Marcelo",
			Apellido:      "Before to update",
			Email:         "marcelopez@gmail.com",
			Edad:          27,
			Altura:        165,
			Activo:        true,
			FechaCreacion: "05/10/2022",
		},
	}
	db := MockStore{readWasCalled: false, mockedData: usuarioRegistrado}
	repo := NewRepository(&db)
	usuarioModificado, err := repo.PartialUpdate(1, ExpectedApellido, ExpectedEdad)

	assert.Nil(t, err)
	assert.True(t, db.readWasCalled)
	assert.Equal(t, usuarioModificado.Apellido, ExpectedApellido)
	assert.Equal(t, usuarioModificado.Edad, ExpectedEdad)
}
