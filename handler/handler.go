package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/osamikoyo/ward/core"

	tokenmw "github.com/osamikoyo/ward/middleware"
)

type Handler struct {
	core *core.WardCore
}

func NewHandler(core *core.WardCore) *Handler {
	return &Handler{
		core: core,
	}
}

func (h *Handler) RegisterRouters(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(tokenmw.TokenMiddleware)

	user := e.Group("/user")
	user.POST("/register", h.RegisterUserHandler)
	user.DELETE("/delete/:uid", h.DeleteUserHandler)

	grand := e.Group("/grand")
	grand.POST("/create", h.CreateGrandHandler)
	grand.DELETE("/delete/:uid", h.DeleteGrandHandler)
	grand.PATCH("/change/level/:uid", h.ChangeGrandLevelHandler)
	grand.GET("/list", h.ListGrandHandler)

	data := e.Group("/data")
	data.POST("/create", h.CreateDataHandler)
	data.PATCH("/change/grand/:uid", h.ChangeDataGrand)
	data.DELETE("/delete/:uid", h.DeleteDataHandler)
	data.GET("/list", h.ListDataHandler)
}
