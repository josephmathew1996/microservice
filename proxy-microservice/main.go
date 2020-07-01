package main

import (
	"proxy-microservice/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	router.Set(e) //setting all the routes
	e.Start(":8083")
}
