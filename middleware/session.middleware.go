package middleware

import (
	"context"
	"devlog/common"
	"devlog/ent"
	dbAdmin "devlog/ent/admin"
	dbAdminSession "devlog/ent/adminsession"
	"devlog/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func WithSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
			session, err := ctx.Client().AdminSession.Query().
				Where(dbAdminSession.And(
					dbAdminSession.TokenEQ(c.Request().Header.Get("api-token")),
					dbAdminSession.ExpiresAtGT(time.Now()))).
				WithUser(func(q *ent.AdminQuery) {
					q.Select(dbAdmin.FieldID)
				}).
				Select(dbAdminSession.FieldID).
				First(context.Background())

			if err != nil {
				if _, ok := err.(*ent.NotFoundError); ok {
					return nil, nil
				}
				return nil, err
			}

			_, err = session.Update().SetExpiresAt(time.Now().Add(time.Hour * 24 * 5)).Save(context.Background())
			if err != nil {
				return nil, err
			}

			return session.Edges.User, nil
		})
		if err != nil {
			c.Error(err)
		}

		if res != nil {
			c.(*common.Context).SetAdmin(res.(*ent.Admin))
		}

		return next(c)
	}
}

func RequireSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.(*common.Context).Admin() == nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		return next(c)
	}
}
