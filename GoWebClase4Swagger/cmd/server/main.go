package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marciolarramendi/backpack-bcgow6-marcio-larramendi/GoWebClase4Swagger/cmd/server/handler"
	"github.com/marciolarramendi/backpack-bcgow6-marcio-larramendi/GoWebClase4Swagger/docs"
	"github.com/marciolarramendi/backpack-bcgow6-marcio-larramendi/GoWebClase4Swagger/internal/users"
	"github.com/marciolarramendi/backpack-bcgow6-marcio-larramendi/GoWebClase4Swagger/pkg/store"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Bootcamp Go Wave 6 - API
// @version         1.0
// @description     This is a simple API development by Marcio.
// @termsOfService  https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name   API Support Digital House
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "./users.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)

	u := handler.NewUser(service)

	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ur := r.Group("/users")
	ur.POST("/", u.ValidateToken(), u.Store())
	ur.GET("/", u.ValidateToken(), u.GetAll())
	ur.PUT("/:id", u.ValidateToken(), u.Update())
	ur.PATCH("/:id", u.ValidateToken(), u.PartialUpdate())
	ur.DELETE("/:id", u.ValidateToken(), u.Delete())
	r.Run()
}
