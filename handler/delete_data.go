package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) DeleteDataHandler(c echo.Context) error {
	token, ok := c.Get("token").(string)
	if !ok {
		return c.String(http.StatusBadGateway, "invalid token")
	}

	uid := c.Param("uid")

	uuid, err := uuid.Parse(uid)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed parse uid")
	}

	if err = h.core.DeleteData(token, uuid); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "ok")
}
