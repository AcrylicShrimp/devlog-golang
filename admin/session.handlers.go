package admin

import (
	"devlog/common"
	"devlog/ent"
	dbAdmin "devlog/ent/admin"
	dbAdminSession "devlog/ent/adminsession"
	"devlog/util"
	"encoding/hex"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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

	client := c.(*common.Context).Client()
	ctx := c.(*common.Context).Ctx()

	admin, err := client.Admin.Query().Where(dbAdmin.UsernameEQ(authInfo.Username)).Select(dbAdmin.FieldID, dbAdmin.FieldPassword).First(ctx)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	password, err := hex.DecodeString(admin.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if bcrypt.CompareHashAndPassword(password, []byte(authInfo.Password)) != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	token, err := util.GenerateToken256()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if _, err := client.AdminSession.Create().SetUser(admin).SetToken(token).Save(ctx); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	type Token struct {
		Token string `json:"token"`
	}

	return c.JSON(http.StatusCreated, Token{Token: token})
}

func DeleteSessionHandler(c echo.Context) error {
	type TokenInfo struct {
		Token string `param:"token" validate:"required"`
	}

	tokenInfo := new(TokenInfo)
	if err := c.Bind(tokenInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(tokenInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	client := c.(*common.Context).Client()
	ctx := c.(*common.Context).Ctx()

	if _, err := client.AdminSession.Delete().Where(dbAdminSession.TokenEQ(tokenInfo.Token)).Exec(ctx); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}
