package main

import (
	"devlog/common"
	_ "devlog/docs"
	"devlog/env"
	"devlog/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

// @title devlog
// @version 0.0.1
// @description A markdown-based lightweight blog service.

// @contact.name Author
// @contact.url https://github.com/AcrylicShrimp
// @contact.email led789zxpp@naver.com

// @license.name MIT

// @BasePath /v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Api-Token

func main() {
	if os.Getenv("DEVEL") == "true" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
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
