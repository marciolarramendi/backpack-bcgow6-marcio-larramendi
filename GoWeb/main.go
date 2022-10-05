package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	usuarios = getUsers()
	token    = "asfcb1245drrt433qyt57"
)

type usuario struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre" binding:"required"`
	Apellido      string `json:"apellido" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Edad          int    `json:"edad" binding:"required"`
	Altura        int    `json:"altura" binding:"required"`
	Activo        bool   `json:"activo" binding:"required"`
	FechaCreacion string `json:"fechacreacion" binding:"required"`
}

func newID() int {
	var maxID int = 0
	for _, value := range usuarios {
		if value.ID > maxID {
			maxID = value.ID
		}
	}
	//maxID++
	return maxID
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

func validateToken(c *gin.Context) {
	tokenRequest := c.GetHeader("token")
	if tokenRequest != token || tokenRequest == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no tiene permisos para realizar la petición solicitada"})
		c.Abort()
	}
	c.Next()
}

func AddUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req usuario
		if err := c.ShouldBindJSON(&req); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				result := ""
				if len(ve) > 1 {
					result += fmt.Sprintf("Los campos ")
					for i, field := range ve {
						if i != len(ve)-1 && i == 0 {
							result += fmt.Sprintf("%s", field.Field())
						} else if i != len(ve)-1 && i > 0 {
							result += fmt.Sprintf(", %s", field.Field())
						} else {
							result += fmt.Sprintf(" y %s", field.Field())
						}
					}
					result += fmt.Sprintf(" son requeridos")
				} else {
					result += fmt.Sprintf("El campo %s es requerido", ve[0].Field())
				}
				c.JSON(400, result)
			}
			return
		}
		req.ID = newID()
		usuarios = append(usuarios, req)
		c.JSON(http.StatusOK, req)
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
	router.POST("/users/add", validateToken, AddUser())

	router.Run()
}
