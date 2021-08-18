package admin

import (
	"context"
	"devlog/common"
	"devlog/ent"
	dbCategory "devlog/ent/category"
	dbPost "devlog/ent/post"
	"devlog/middleware"
	"devlog/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AttachCategory(group *echo.Group) {
	group.GET("", ListCategories, middleware.WithSession, middleware.RequireSession)
	group.POST("", NewCategoryHandler, middleware.WithSession, middleware.RequireSession)
	group.DELETE("/:name", DeleteCategoryHandler, middleware.WithSession, middleware.RequireSession)
}

// ListCategories godoc
// @router /admin/categories [get]
// @summary List categories
// @description Lists all categories.
// @produce json
// @success 200 {array} model.AdminCategory
// @failure 401 {object} model.HTTPError
// @failure 500 {object} model.HTTPError
func ListCategories(c echo.Context) error {
	ctx := c.(*common.Context)

	categories, err := ctx.Client().Category.Query().
		Select(
			dbCategory.FieldName,
			dbCategory.FieldDescription,
			dbCategory.FieldCreatedAt,
			dbCategory.FieldModifiedAt).
		Order(ent.Asc(dbCategory.FieldName)).
		All(context.Background())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, categories)
}

func NewCategoryHandler(c echo.Context) error {
	type CategoryInfo struct {
		Name        string  `json:"name" form:"name" query:"name" validate:"required,max=32"`
		Description *string `json:"description" form:"description" query:"description" validate:"required,min=1,max=255"`
	}

	categoryInfo := new(CategoryInfo)
	if err := c.Bind(categoryInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(categoryInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	ctx := c.(*common.Context)

	_, err := ctx.Client().Category.Create().
		SetName(categoryInfo.Name).
		SetNillableDescription(categoryInfo.Description).
		Save(context.Background())
	if err != nil {
		if _, ok := err.(*ent.ConstraintError); ok {
			return echo.NewHTTPError(http.StatusConflict)
		}

		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	type Category struct {
		Name string `json:"name"`
	}

	return c.JSON(http.StatusCreated, Category{Name: categoryInfo.Name})
}

func DeleteCategoryHandler(c echo.Context) error {
	type CategoryInfo struct {
		Name string `param:"name" validate:"required"`
	}

	categoryInfo := new(CategoryInfo)
	if err := c.Bind(categoryInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(categoryInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		category, err := ctx.Client().Category.Query().
			Where(dbCategory.NameEQ(categoryInfo.Name)).
			Select(dbCategory.FieldID).
			First(context.Background())
		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, echo.NewHTTPError(http.StatusNotFound)
			}

			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		if _, err := tx.Post.Update().
			Where(dbPost.HasCategoryWith(dbCategory.IDEQ(category.ID))).
			ClearCategory().
			Save(context.Background()); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		return nil, nil
	})
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
