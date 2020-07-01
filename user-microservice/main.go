package main

import (
	"user-microservice/container"
	"user-microservice/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	container.LoadData()
	router.Set(e) //setting all the routes
	e.Start(":8082")
}
