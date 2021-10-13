package admin

import (
	"context"
	"devlog/common"
	"devlog/ent"
	dbAdmin "devlog/ent/admin"
	dbRobotAccess "devlog/ent/robotaccess"
	"devlog/model"
	"devlog/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func AttachRobotAccess(group *echo.Group) {
	group.GET("", ListRobotAccesses)
	group.POST("", NewRobotAccess)
	group.DELETE("/:token", DeleteRobotAccess)
}

// ListRobotAccesses godoc
// @router /admin/robot-accesses [get]
// @summary List robot accesses
// @description Lists all robot accesses.
// @description The robot accesses are sorted by the field 'last-access-at' in descending order.
// @tags admin robot access management
// @produce json
// @success 200 {array} model.RobotAccess
// @failure 401 {object} model.HTTPError401
// @failure 500 {object} model.HTTPError500
func ListRobotAccesses(c echo.Context) error {
	ctx := c.(*common.Context)

	accesses, err := ctx.Client().RobotAccess.Query().
		Select(
			dbRobotAccess.FieldID,
			dbRobotAccess.FieldMemo,
			dbRobotAccess.FieldCreatedAt,
			dbRobotAccess.FieldExpiresAt,
			dbRobotAccess.FieldLastAccessAt).
		WithUser(func(q *ent.AdminQuery) {
			q.Select(dbAdmin.FieldUsername)
		}).
		Order(
			ent.Desc(dbRobotAccess.FieldLastAccessAt),
			ent.Asc(dbRobotAccess.FieldCreatedAt)).
		All(context.Background())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	accessJSONs := make([]model.RobotAccess, len(accesses))

	for index, access := range accesses {
		accessJSONs[index] = model.RobotAccessFromModel(access)
	}

	return c.JSON(http.StatusOK, accessJSONs)
}

// NewRobotAccess godoc
// @router /admin/robot-accesses [post]
// @summary Create robot access
// @description Creates a robot access.
// @tags admin robot access management
// @produce json
// @success 201 {object} model.RobotAccessTokenOnly
// @failure 400 {object} model.HTTPError400
// @failure 401 {object} model.HTTPError401
// @failure 500 {object} model.HTTPError500
func NewRobotAccess(c echo.Context) error {
	param := new(model.RobotAccessParam)

	if err := c.Bind(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	ctx := c.(*common.Context)

	token, err := util.GenerateToken256()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	var expiresAt *time.Time

	if param.TTL != nil {
		value := time.Now().Add(time.Duration(*param.TTL) * time.Second)
		expiresAt = &value
	}

	if _, err := ctx.Client().RobotAccess.Create().
		SetToken(token).
		SetNillableMemo(param.Memo).
		SetNillableExpiresAt(expiresAt).
		Save(context.Background()); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, model.RobotAccessTokenOnlyFromToken(token))
}

// DeleteRobotAccess godoc
// @router /admin/robot-accesses/{token} [delete]
// @summary Remove robot access
// @description Removes the given robot access.
// @tags admin robot access management
// @param token path string true "A robot access token to be removed"
// @produce json
// @success 204 "NoContent: when the robot access has been removed successfully"
// @failure 400 {object} model.HTTPError400
// @failure 401 {object} model.HTTPError401
// @failure 404 {object} model.HTTPError404
// @failure 500 {object} model.HTTPError500
func DeleteRobotAccess(c echo.Context) error {
	param := new(model.RobotAccessTokenParam)

	if err := c.Bind(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		robotAccess, err := ctx.Client().RobotAccess.Query().
			Where(dbRobotAccess.TokenEQ(param.Token)).
			Select(dbRobotAccess.FieldID).
			First(context.Background())

		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, echo.NewHTTPError(http.StatusNotFound)
			}
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		if err := tx.RobotAccess.DeleteOneID(robotAccess.ID).Exec(context.Background()); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		return nil, nil
	})

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
