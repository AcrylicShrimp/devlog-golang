package admin

import "github.com/labstack/echo/v4"

func Attach(group *echo.Group) {
	AttachCategory(group.Group("/categories"))
	AttachPost(group.Group("/posts"))
	AttachSession(group.Group("/sessions"))
	AttachUnsavedPost(group.Group("/unsaved-posts"))
}
