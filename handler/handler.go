package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/osamikoyo/ward/core"
)

type Handler struct {
	core *core.WardCore
}

func (h *Handler) RegisterRouters(e *echo.Echo) {
	e.Use(middleware.Logger())

	user := e.Group("/user")
	user.POST("/register", h.RegisterUserHandler)

}
