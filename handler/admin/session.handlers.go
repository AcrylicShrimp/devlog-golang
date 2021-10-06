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
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func AttachSession(group *echo.Group) {
	group.POST("", NewSessionHandler)
}

func NewSessionHandler(c echo.Context) error {
	param := new(model.SessionParam)

	if err := c.Bind(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	res, err := util.WithTx(
		c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
			admin, err := tx.Admin.Query().
				Where(dbAdmin.UsernameEQ(param.Username)).
				Select(dbAdmin.FieldID, dbAdmin.FieldPassword, dbAdmin.FieldKey).
				First(context.Background())

			if err != nil {
				if _, ok := err.(*ent.NotFoundError); ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized)
				}
				return nil, echo.NewHTTPError(http.StatusInternalServerError)
			}

			if bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(param.Password)) != nil {
				return nil, echo.NewHTTPError(http.StatusUnauthorized)
			}

			id, err := util.GenerateToken64()

			if err != nil {
				return nil, echo.NewHTTPError(http.StatusInternalServerError)
			}

			now := time.Now()
			claims := model.SessionClaim{
				Key: admin.Key,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: now.AddDate(1, 0, 0).Unix(),
					Id:        id,
					IssuedAt:  now.Unix(),
					Issuer:    param.Username,
					NotBefore: now.Unix(),
					Subject:   "admin",
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			signed, err := token.SignedString([]byte(env.JWTSigningKey))

			if err != nil {
				return nil, echo.NewHTTPError(http.StatusInternalServerError)
			}

			type Token struct {
				Token string `json:"token"`
			}

			return Token{Token: signed}, nil
		},
	)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}
