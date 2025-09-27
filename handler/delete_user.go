package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) DeleteUserHandler(c echo.Context) error {
	token, ok := c.Get("token").(string)
	if !ok {
		return c.String(http.StatusBadGateway, "token not found")
	}

	useruid := c.Param("uid")
	uid, err := uuid.Parse(useruid)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed parse uid")
	}

	if err = h.core.DeleteData(token, uid); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "ok")
}
