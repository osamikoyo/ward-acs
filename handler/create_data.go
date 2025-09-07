package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateDataHandler(c echo.Context) error {
	token := c.Param("token")

	data := struct {
		Payload  string `json:"payload"`
		GrandUID string `json:"grand_uid"`
		DoEnc    bool   `json:"do_enc"`
	}{}

	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, "failed bind body to data")
	}

	uid, err := uuid.Parse(data.GrandUID)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed parse uuid")
	}

	if err := h.core.CreateData(token, data.Payload, data.DoEnc, uid); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusCreated, "ok")
}
