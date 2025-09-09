package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/osamikoyo/ward/core"
)

type Handler struct {
	core *core.WardCore
}

func (h *Handler) RegisterRouters(e *echo.Echo) {

}
