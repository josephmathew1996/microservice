package router

import (
	"auth-service/handlers"

	"github.com/labstack/echo/v4"
)

//Set sets all routes
func Set(e *echo.Echo) {
	e.POST("/auth", handlers.Authenticate)
}
