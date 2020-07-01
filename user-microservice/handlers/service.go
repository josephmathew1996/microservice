package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//GetMicroServiceName gets micro service name
func GetMicroServiceName(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"micro-service-name": "user-microservice"})
}
