package admin

import (
	"context"
	"devlog/common"
	"devlog/ent"
	dbRobotAccess "devlog/ent/robotaccess"
	"devlog/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AttachRobotAccess(group *echo.Group) {
	group.GET("", ListRobotAccessesHandler)
}

func ListRobotAccessesHandler(c echo.Context) error {
	ctx := c.(*common.Context)

	accesses, err := ctx.Client().RobotAccess.Query().
		Select(
			dbRobotAccess.FieldID,
			dbRobotAccess.FieldMemo,
			dbRobotAccess.FieldCreatedAt,
			dbRobotAccess.FieldExpiresAt,
			dbRobotAccess.FieldLastAccessAt).
		Order(ent.Desc(dbRobotAccess.FieldLastAccessAt)).
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
