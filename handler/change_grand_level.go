package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) ChangeGrandLevelHandler(c echo.Context) error {
	token := c.Param("token")

	req := struct {
		GrandUID string `json:"grand_uid"`
		Level    int    `json:"level"`
	}{}

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "failed bind body to request")
	}

	grandUID, err := uuid.Parse(req.GrandUID)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed parse grand uid")
	}

	if err = h.core.ChangeGrandLevel(token, grandUID, req.Level); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "ok")
}
