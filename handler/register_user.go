package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterUserHandler(c echo.Context) error {
	token := c.Param("token")

	usr := struct {
		Token    string `json:"token"`
		GrandUID string `json:"grand_uid"`
	}{}

	if err := c.Bind(&usr); err != nil {
		return c.String(http.StatusBadRequest, "failed bind user")
	}

	uid, err := uuid.Parse(usr.GrandUID)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed parse uid")
	}

	tokenUsr, err := h.core.RegisterUser(token, uid)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusCreated, tokenUsr)
}
