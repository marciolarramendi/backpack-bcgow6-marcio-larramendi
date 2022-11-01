package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marciolarramendi/backpack-bcgow6-marcio-larramendi/GoTesting/C2-TM/cmd/server/handler"
	"github.com/marciolarramendi/backpack-bcgow6-marcio-larramendi/GoTesting/C2-TM/internal/users"
	"github.com/marciolarramendi/backpack-bcgow6-marcio-larramendi/GoTesting/C2-TM/pkg/store"
)

func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "./users.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)

	u := handler.NewUser(service)

	r := gin.Default()

	ur := r.Group("/users")
	ur.POST("/", u.ValidateToken(), u.Store())
	ur.GET("/", u.ValidateToken(), u.GetAll())
	ur.PUT("/:id", u.ValidateToken(), u.Update())
	ur.PATCH("/:id", u.ValidateToken(), u.PartialUpdate())
	ur.DELETE("/:id", u.ValidateToken(), u.Delete())
	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
