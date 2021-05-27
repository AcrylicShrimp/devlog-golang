package admin

import (
	"context"
	"devlog/common"
	"devlog/ent"
	dbAdmin "devlog/ent/admin"
	dbAdminSession "devlog/ent/adminsession"
	"devlog/util"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func AttachSession(group *echo.Group) {
	group.POST("", NewSessionHandler)
	group.DELETE("/:token", DeleteSessionHandler)
}

func NewSessionHandler(c echo.Context) error {
	type AuthInfo struct {
		Username string `json:"username" form:"username" query:"username" validate:"required"`
		Password string `json:"password" form:"password" query:"password" validate:"required"`
	}

	authInfo := new(AuthInfo)
	if err := c.Bind(authInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(authInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	res, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		admin, err := tx.Admin.Query().
			Where(dbAdmin.UsernameEQ(authInfo.Username)).
			Select(dbAdmin.FieldID, dbAdmin.FieldPassword).
			First(context.Background())
		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized)
			}
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		if bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(authInfo.Password)) != nil {
			return nil, echo.NewHTTPError(http.StatusUnauthorized)
		}

		token, err := util.GenerateToken256()
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		if _, err := tx.AdminSession.Create().
			SetUser(admin).
			SetToken(token).
			SetExpiresAt(time.Now().Add(time.Hour * 24 * 5)).
			Save(context.Background()); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		type Token struct {
			Token string `json:"token"`
		}

		return Token{Token: token}, nil
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)

}

func DeleteSessionHandler(c echo.Context) error {
	type TokenInfo struct {
		Token string `param:"token" validate:"required,hexadecimal,len=256"`
	}

	tokenInfo := new(TokenInfo)
	if err := c.Bind(tokenInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(tokenInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	ctx := c.(*common.Context)

	if _, err := ctx.Client().AdminSession.Delete().
		Where(dbAdminSession.TokenEQ(tokenInfo.Token)).
		Exec(context.Background()); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}
