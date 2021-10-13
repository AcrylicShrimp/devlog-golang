package admin

import (
	"devlog/middleware"
	"github.com/labstack/echo/v4"
)

func Attach(group *echo.Group) {
	AttachCategory(group.Group("/categories", middleware.RequireSession))
	AttachPost(group.Group("/posts", middleware.RequireSession))
	AttachRobotAccess(group.Group("/robot-accesses", middleware.RequireSession))
	AttachSession(group.Group("/sessions"))
	AttachUnsavedPost(group.Group("/unsaved-posts", middleware.RequireSession))
}
