package admin

import (
	"devlog/common"
	"devlog/ent"
	dbAdmin "devlog/ent/admin"
	dbAdminSession "devlog/ent/adminsession"
	"github.com/labstack/echo"
	"net/http"
)

func WithSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		client := c.(*common.Context).Client()
		ctx := c.(*common.Context).Ctx()

		admin, err := client.AdminSession.Query().Where(dbAdminSession.TokenEQ(c.Request().Header.Get("api-token"))).WithUser(func(q *ent.AdminQuery) {
			q.Select(dbAdmin.FieldID)
		}).First(ctx)
		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return next(c)
			}
			c.Error(err)
		}

		c.(*common.Context).SetAdmin(admin.Edges.User)
		return next(c)
	}
}

func RequireSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.(*common.Context).Admin() == nil {
			c.NoContent(http.StatusUnauthorized)
			return nil
		}

		return next(c)
	}
}
