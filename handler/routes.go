package handler

import (
	"devlog/env"
	"devlog/handler/admin"
	devlogMiddleware "devlog/middleware"
	"devlog/model"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Attach(group *echo.Group) {
	group.Use(
		echoMiddleware.JWTWithConfig(
			echoMiddleware.JWTConfig{
				Skipper: func(c echo.Context) bool {
					return c.Request().Header.Get("Authorization") == ""
				},
				SigningKey:    []byte(env.JWTSigningKey),
				SigningMethod: "HS256",
				ContextKey:    "admin",
				Claims:        &model.SessionClaim{},
			},
		),
		devlogMiddleware.WithSession)

	AttachPost(group.Group("/posts"))
	admin.Attach(group.Group("/admin"))
}
