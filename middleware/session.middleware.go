package middleware

import (
	"context"
	"devlog/common"
	"devlog/ent"
	dbAdmin "devlog/ent/admin"
	dbRobotAccess "devlog/ent/robotaccess"
	"devlog/model"
	"devlog/util"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func WithSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if token, ok := c.Get("admin").(*jwt.Token); ok {
			claims := token.Claims.(*model.SessionClaim)
			admin, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
				admin, err := ctx.Client().Admin.Query().
					Where(dbAdmin.And(
						dbAdmin.UsernameEQ(claims.Issuer),
						dbAdmin.KeyEQ(claims.Key))).
					Select(dbAdmin.FieldID).
					First(context.Background())

				if err != nil {
					if _, ok := err.(*ent.NotFoundError); ok {
						return nil, nil
					}
					return nil, err
				}

				if err != nil {
					return nil, err
				}

				return admin, nil
			})

			if err != nil {
				c.Error(err)
			}

			if admin != nil {
				c.(*common.Context).SetAdmin(admin.(*ent.Admin))
			}
		} else if token := c.Request().Header.Get("robot-access-token"); token != "" {
			admin, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
				robotAccess, err := ctx.Client().RobotAccess.Query().
					Where(dbRobotAccess.TokenEQ(token)).
					WithUser(func(q *ent.AdminQuery) {
						q.Select(dbAdmin.FieldID)
					}).
					Select(dbRobotAccess.FieldID, dbRobotAccess.FieldExpiresAt).
					First(context.Background())

				if err != nil {
					if _, ok := err.(*ent.NotFoundError); ok {
						return nil, nil
					}
					return nil, echo.NewHTTPError(http.StatusInternalServerError)
				}

				now := time.Now()

				if robotAccess.ExpiresAt.Before(now) {
					return nil, nil
				}

				_, err = robotAccess.Update().
					SetLastAccessAt(now).
					Save(context.Background())

				if err != nil {
					return nil, echo.NewHTTPError(http.StatusInternalServerError)
				}

				return robotAccess.Edges.User, nil
			})

			if err != nil {
				c.Error(err)
			}

			if admin != nil {
				c.(*common.Context).SetAdmin(admin.(*ent.Admin))
			}
		}

		return next(c)
	}
}

func RequireSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.(*common.Context).Admin() == nil {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return next(c)
	}
}
