package router

import (
	"user-microservice/handlers"

	"github.com/labstack/echo/v4"
)

//Set sets all routes
func Set(e *echo.Echo) {
	e.GET("/user/profile", handlers.GetUserProfile)
	e.GET("/microservice/name", handlers.GetMicroServiceName)
}
