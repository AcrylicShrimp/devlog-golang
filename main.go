package main

import (
	"devlog/common"
	_ "devlog/docs"
	"devlog/env"
	"devlog/handler"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/swaggo/echo-swagger"
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
	development := os.Getenv("DEVEL") == "true"

	if development {

		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}

	env.InitEnvVars()

	e := echo.New()
	e.Debug = development
	e.HideBanner = true
	e.Validator = NewValidator()

	if development {
		e.Logger.SetLevel(log.DEBUG)
	}

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

	if development {
		apiDocs := e.Group("/api-docs")
		apiDocs.GET("/*", echoSwagger.WrapHandler)
		e.Logger.Infof("Swagger API docs is live on http://localhost:%v/api-docs/index.html\n", env.Port)
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%v", env.Port)))
}
