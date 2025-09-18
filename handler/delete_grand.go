package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) DeleteGrandHandler(c echo.Context) error {
	token := c.Param("token")

	uid := c.Param("grand")
	id, err := uuid.Parse(uid)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed parse uid")
	}

	if err = h.core.DeleteData(token, id); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "ok")
}
