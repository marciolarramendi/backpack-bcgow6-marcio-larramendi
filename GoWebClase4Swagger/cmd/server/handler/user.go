package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marciolarramendi/backpack-bcgow6-marcio-larramendi/GoWebClase4Swagger/internal/users"
	"github.com/marciolarramendi/backpack-bcgow6-marcio-larramendi/GoWebClase4Swagger/pkg/web"
)

type request struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
	Edad          int    `json:"edad"`
	Altura        int    `json:"altura"`
	Activo        bool   `json:"activo"`
	FechaCreacion string `json:"fechacreacion"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (c *User) ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			ctx.Abort()
		}
	}
}

// ListUsers godoc
// @Summary List users
// @Tags Users
// @Description get users
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response "List Users"
// @Failure 401 {object} web.Response "Unauthorized"
// @Failure 404 {object} web.Response "Not found user"
// @Failure 500 {object} web.Response "Internal Server Error"
// @Router /users [GET]
func (c *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

// StoreUsers godoc
// @Summary Store users
// @Tags Users
// @Description store users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param user body request true "User to store"
// @Success 200 {object} web.Response "Store User"
// @Failure 401 {object} web.Response "Unauthorized"
// @Failure 404 {object} web.Response "Not found user"
// @Failure 500 {object} web.Response "Internal Server Error"
// @Router /users [POST]
func (c *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		u, err := c.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

// UpdateUser godoc
// @Summary Update user
// @Tags Users
// @Description update user
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "ID user"
// @Param user body request true "User to update"
// @Success 200 {object} web.Response "Update User"
// @Failure 400 {object} web.Response "Bad Request"
// @Failure 401 {object} web.Response "Unauthorized"
// @Failure 404 {object} web.Response "Not found user"
// @Failure 500 {object} web.Response "Internal Server Error"
// @Router /users/{id} [PUT]
func (c *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID inválido"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Nombre == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del usuario es requerido"))
			return
		}
		if req.Apellido == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El apellido del usuario es requerido"))
			return
		}
		if req.Email == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El email del usuario es requerido"))
			return
		}
		if req.Edad == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "La edad del usuario es requerido"))
			return
		}
		if req.Altura == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "La altura del usuario es requerido"))
			return
		}
		if !req.Activo {
			ctx.JSON(400, web.NewResponse(400, nil, "El estado del usuario es requerido"))
			return
		}
		if req.FechaCreacion == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "La fecha de creacion del usuario es requerido"))
			return
		}
		u, err := c.service.Update(int(id), req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

// PartialUpdateUser godoc
// @Summary Partial update user
// @Tags Users
// @Description partial update user
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "User ID"
// @Param apellido body request true "Apellido"
// @Param edad body request true "Edad"
// @Success 200 {object} web.Response "Partial update User"
// @Failure 400 {object} web.Response "Bad Request"
// @Failure 401 {object} web.Response "Unauthorized"
// @Failure 404 {object} web.Response "User Not Found"
// @Failure 500 {object} web.Response "Internal Server Error"
// @Router /users{id} [PATCH]
func (c *User) PartialUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID inválido"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Apellido == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El apellido del usuario es requerido"))
			return
		}
		if req.Edad == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "La edad del usuario es requerida"))
			return
		}
		u, err := c.service.PartialUpdate(int(id), req.Apellido, req.Edad)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

// DeleteUser godoc
// @Summary Delete user
// @Tags Users
// @Description delete user
// @Param token header string true "token"
// @Param id path int true "User ID"
// @Success 200 {object} web.Response "Update User"
// @Failure 401 {object} web.Response "Unauthorized"
// @Failure 404 {object} web.Response "Not found user"
// @Failure 500 {object} web.Response "Internal Server Error"
// @Router /users/{id} [DELETE]
func (c *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID inválido"))
			return
		}
		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		message := fmt.Sprintf("El usuario %d ha sido eliminado", id)
		ctx.JSON(200, web.NewResponse(200, message, ""))
	}
}
