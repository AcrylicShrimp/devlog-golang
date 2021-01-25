package main

import (
	"devlog/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler.Attach(e.Group("v1"))

	e.Logger.Fatal(e.Start("0.0.0.0:8000"))
}
