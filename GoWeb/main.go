package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var usuarios = getUsers()

type usuario struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
	Edad          int    `json:"edad"`
	Altura        int    `json:"altura"`
	Activo        bool   `json:"activo"`
	FechaCreacion string `json:"fechacreacion"`
}

func getUsers() []usuario {
	data, err := os.ReadFile("usuarios.json")
	if err != nil {
		panic(err)
	}
	var u []usuario
	if err := json.Unmarshal(data, &u); err != nil {
		panic(err)
	}
	return u
}

func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, usuarios)
}

func GetAllWithFilters(c *gin.Context) {
	var usuariosEcontrados []usuario
	id, _ := strconv.Atoi(c.Query("id"))
	nombre := c.Query("nombre")
	apellido := c.Query("apellido")
	email := c.Query("email")
	edad, _ := strconv.Atoi(c.Query("edad"))
	altura, _ := strconv.Atoi(c.Query("altura"))
	activo := c.Query("activo")
	fechacreacion := c.Query("fechacreacion")

	for _, usuario := range usuarios {
		if usuario.ID == id || usuario.Nombre == nombre || usuario.Apellido == apellido || usuario.Email == email || usuario.Edad == edad || usuario.Altura == altura || strconv.FormatBool(usuario.Activo) == activo || usuario.FechaCreacion == fechacreacion {
			usuariosEcontrados = append(usuariosEcontrados, usuario)
		}
	}
	if len(usuariosEcontrados) > 0 {
		c.String(200, "%v \n", usuariosEcontrados)
	} else {
		c.String(404, "No se encontraron usuarios\n")
	}
}

func GetByID(c *gin.Context) {
	var usuarioEcontrado []usuario
	id, _ := strconv.Atoi(c.Param("id"))
	for _, usuario := range usuarios {
		if usuario.ID == id {
			usuarioEcontrado = append(usuarioEcontrado, usuario)
			break
		}
	}
	if len(usuarioEcontrado) > 0 {
		c.String(200, "%v \n", usuarioEcontrado)
	} else {
		c.String(404, "No se encontró el usuario\n")
	}
}

func main() {
	router := gin.Default()

	router.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Marcio!",
		})
	})

	router.GET("/users", GetAll)
	router.GET("/users/", GetAllWithFilters)
	router.GET("/users/:id", GetByID)

	router.Run()
}
