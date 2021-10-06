package admin

import (
	"context"
	"devlog/common"
	"devlog/ent"
	dbAdmin "devlog/ent/admin"
	"devlog/env"
	"devlog/model"
	"devlog/util"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Attach(group *echo.Group) {
	group.Use(
		middleware.JWTWithConfig(
			middleware.JWTConfig{
				SigningKey:    []byte(env.JWTSigningKey),
				SigningMethod: "HS256",
				ContextKey:    "admin",
				Claims:        &model.SessionClaim{},
			},
		),
	)
	group.Use(
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				token := c.Get("admin").(*jwt.Token)
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
					return echo.NewHTTPError(http.StatusUnauthorized)
				}

				c.(*common.Context).SetAdmin(admin.(*ent.Admin))

				return next(c)
			}
		},
	)
	AttachCategory(group.Group("/categories"))
	AttachPost(group.Group("/posts"))
	AttachSession(group.Group("/sessions"))
	AttachUnsavedPost(group.Group("/unsaved-posts"))
}
