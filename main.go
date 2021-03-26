package main

import (
	"context"
	"devlog/admin"
	"devlog/common"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.HideBanner = true
	e.Validator = NewValidator()

	client, err := InitDB()
	if err != nil {
		panic(err)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(common.NewContext(c, client, context.Background()))
		}
	})

	v1 := e.Group("/v1")
	admin.Attach(v1.Group("/admin"))

	e.Logger.Fatal(e.Start("0.0.0.0:8000"))
}
