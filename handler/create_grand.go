package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateGrandHandler(c echo.Context) error {
	token, ok := c.Get("token").(string)
	if !ok {
		return c.String(http.StatusBadGateway, "invalid token")
	}

	grand := struct {
		Level int    `json:"level"`
		Name  string `json:"name"`
	}{}

	if err := c.Bind(&grand); err != nil {
		return c.String(http.StatusBadRequest, "failed bind body to grand")
	}

	if err := h.core.CreateGrand(token, grand.Name, grand.Level); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusCreated, "ok")
}
