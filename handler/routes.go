package handler

import (
	"devlog/handler/admin"
	"github.com/labstack/echo/v4"
)

func Attach(group *echo.Group) {
	AttachPost(group.Group("/posts"))
	admin.Attach(group.Group("/admin"))
}
