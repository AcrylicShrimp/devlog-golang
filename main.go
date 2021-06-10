package main

import (
	"devlog/common"
	"devlog/env"
	"devlog/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	env.InitEnvVars()

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
			return next(common.NewContext(c, client))
		}
	})

	handler.Attach(e.Group("/v1"))

	e.Logger.Fatal(e.Start("0.0.0.0:8000"))
}
