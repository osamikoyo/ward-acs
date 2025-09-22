package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func TokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Token")
		if len(token) == 0 {
			return c.String(http.StatusBadGateway, "token is empty")
		}

		c.Set("token", token)

		return next(c)
	}
}
