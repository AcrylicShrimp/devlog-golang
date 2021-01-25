package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func Attach(group *echo.Group) {
	group.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})
}
