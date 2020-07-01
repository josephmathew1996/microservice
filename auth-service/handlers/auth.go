package handlers

import (
	"auth-service/container"
	"net/http"

	"github.com/labstack/echo/v4"
)

//Authenticate handler
func Authenticate(c echo.Context) error {
	username := c.Request().Header.Get("Username")
	if username == container.User["username"] {
		return c.JSON(http.StatusOK, nil)
	}
	return echo.NewHTTPError(401)
}
