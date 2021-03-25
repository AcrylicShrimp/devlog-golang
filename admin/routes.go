package admin

import "github.com/labstack/echo"

func Attach(group *echo.Group) {
	AttachPost(group.Group("/posts"))
	AttachSession(group.Group("/sessions"))
}
