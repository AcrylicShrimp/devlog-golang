package admin

import (
	"crypto/rand"
	"devlog/common"
	"devlog/ent"
	"devlog/ent/admin"
	"encoding/hex"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func AttachSession(group *echo.Group) {
	group.POST("", NewSessionHandler)
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

	user, err := client.Admin.Query().Where(admin.UsernameEQ(authInfo.Username)).Select(admin.FieldID, admin.FieldPassword).First(ctx)

	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authInfo.Password)) != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	bytes := make([]byte, 128)

	if _, err := rand.Read(bytes); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	token := hex.EncodeToString(bytes)

	if _, err := client.AdminSession.Create().SetUser(user).SetToken(token).Save(ctx); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	type Token struct {
		Token string `json:"token"`
	}

	return c.JSON(http.StatusCreated, Token{Token: token})
}
