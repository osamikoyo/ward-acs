package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) ChangeDataGrand(c echo.Context) error {
	token, ok := c.Get("token").(string)
	if !ok {
		return c.String(http.StatusBadGateway, "invalid token")
	}

	req := struct {
		Data  string `json:"data"`
		Grand string `json:"grand"`
	}{}

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "failed bind body to req")
	}

	dataUID, err := uuid.Parse(req.Data)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed parse data uid")
	}

	grandUID, err := uuid.Parse(req.Grand)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed parse grand uid")
	}

	if err = h.core.ChangeDataGrand(token, dataUID, grandUID); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "ok")
}
