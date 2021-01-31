package admin

import "github.com/labstack/echo"

func Attach(group *echo.Group) {
	AttachSession(group.Group("/sessions"))
}
