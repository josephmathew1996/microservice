package handlers

import (
	"net/http"
	"user-microservice/container"

	"github.com/labstack/echo/v4"
)

//GetUserProfile handler
func GetUserProfile(c echo.Context) error {
	return c.JSON(http.StatusOK, container.User)
}
