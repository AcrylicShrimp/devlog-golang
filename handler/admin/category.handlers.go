package admin

import (
	"context"
	"devlog/common"
	"devlog/ent"
	dbCategory "devlog/ent/category"
	dbPost "devlog/ent/post"
	"devlog/model"
	"devlog/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AttachCategory(group *echo.Group) {
	group.GET("", ListCategories)
	group.POST("", NewCategory)
	group.DELETE("/:name", DeleteCategory)
}

// ListCategories godoc
// @router /admin/categories [get]
// @summary List categories
// @description Lists all categories.
// @description The categories are sorted by the field 'name' in ascending order.
// @tags admin category management
// @produce json
// @success 200 {array} model.Category
// @failure 401 {object} model.HTTPError401
// @failure 500 {object} model.HTTPError500
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

	categoryJSONs := make([]model.Category, len(categories))

	for index, category := range categories {
		categoryJSONs[index] = model.CategoryFromModel(category)
	}

	return c.JSON(http.StatusOK, categoryJSONs)
}

// NewCategory godoc
// @router /admin/categories [post]
// @summary Create category
// @description Creates a new category.
// @description The field 'name' must be unique across all categories.
// @tags admin category management
// @accept json
// @param category body model.NewCategoryParam true "The category to be created"
// @produce json
// @success 201 "NoContent: when the category has been removed successfully"
// @failure 400 {object} model.HTTPError400
// @failure 401 {object} model.HTTPError401
// @failure 409 {object} model.HTTPError409 "Conflict: when the name is not unique(already taken)"
// @failure 500 {object} model.HTTPError500
func NewCategory(c echo.Context) error {
	param := new(model.CategoryParam)

	if err := c.Bind(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	ctx := c.(*common.Context)

	if _, err := ctx.Client().Category.Create().
		SetName(param.Name).
		SetNillableDescription(param.Description).
		Save(context.Background()); err != nil {
		if _, ok := err.(*ent.ConstraintError); ok {
			return echo.NewHTTPError(http.StatusConflict)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusCreated)
}

// DeleteCategory godoc
// @router /admin/categories/{name} [delete]
// @summary Remove category
// @description Removes the given category.
// @tags admin category management
// @param name path string true "A category name to be removed"
// @produce json
// @success 204 "NoContent: when the category has been removed successfully"
// @failure 400 {object} model.HTTPError400
// @failure 401 {object} model.HTTPError401
// @failure 404 {object} model.HTTPError404
// @failure 500 {object} model.HTTPError500
func DeleteCategory(c echo.Context) error {
	param := new(model.CategoryNameParam)

	if err := c.Bind(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		category, err := ctx.Client().Category.Query().
			Where(dbCategory.NameEQ(param.Name)).
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

		if err := tx.Category.DeleteOneID(category.ID).Exec(context.Background()); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		return nil, nil
	})

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
