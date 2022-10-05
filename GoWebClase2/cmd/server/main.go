package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marciolarramendi/backpack-bcgow6-marcio-larramendi/GoWebClase2/cmd/server/handler"
	"github.com/marciolarramendi/backpack-bcgow6-marcio-larramendi/GoWebClase2/internal/users"
)

func main() {
	repo := users.NewRepository()
	service := users.NewService(repo)

	u := handler.NewUser(service)

	r := gin.Default()

	ur := r.Group("/users")
	ur.POST("/", u.Store())
	ur.GET("/", u.GetAll())
	ur.PUT("/:id", u.Update())
	r.Run()
}
