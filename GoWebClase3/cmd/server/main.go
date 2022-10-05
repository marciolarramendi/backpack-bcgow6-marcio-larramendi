package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marciolarramendi/backpack-bcgow6-marcio-larramendi/GoWebClase3/cmd/server/handler"
	"github.com/marciolarramendi/backpack-bcgow6-marcio-larramendi/GoWebClase3/internal/users"
	"github.com/marciolarramendi/backpack-bcgow6-marcio-larramendi/GoWebClase3/pkg/store"
)

func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "./users.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)

	u := handler.NewUser(service)

	r := gin.Default()

	ur := r.Group("/users")
	ur.POST("/", u.Store())
	ur.GET("/", u.GetAll())
	ur.PUT("/:id", u.Update())
	ur.PATCH("/:id", u.PartialUpdate())
	ur.DELETE("/:id", u.Delete())
	r.Run()
}
