package admin

import (
	"devlog/common"
	"devlog/ent"
	dbCategory "devlog/ent/category"
	dbPost "devlog/ent/post"
	"github.com/labstack/echo"
	"net/http"
)

func AttachCategory(group *echo.Group) {
	group.GET("", ListCategories)
	group.POST("", NewCategoryHandler)
	group.DELETE("/:name", DeleteCategoryHandler)
}

func ListCategories(c echo.Context) error {
	client := c.(*common.Context).Client()
	ctx := c.(*common.Context).Ctx()

	categories, err := client.Category.Query().Select(dbCategory.FieldName, dbCategory.FieldDescription, dbCategory.FieldCreatedAt, dbCategory.FieldModifiedAt).Order(ent.Asc(dbCategory.FieldName)).All(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, categories)
}

func NewCategoryHandler(c echo.Context) error {
	type CategoryInfo struct {
		Name        string  `json:"name" form:"name" query:"name" validate:"required,max=32"`
		Description *string `json:"description" form:"description" query:"description" validate:"omitempty,min=1,max=255"`
	}

	categoryInfo := new(CategoryInfo)
	if err := c.Bind(categoryInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(categoryInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	client := c.(*common.Context).Client()
	ctx := c.(*common.Context).Ctx()

	_, err := client.Category.Create().SetName(categoryInfo.Name).SetNillableDescription(categoryInfo.Description).Save(ctx)
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

	client := c.(*common.Context).Client()
	ctx := c.(*common.Context).Ctx()

	category, err := client.Category.Query().Where(dbCategory.NameEQ(categoryInfo.Name)).Select(dbCategory.FieldID).First(ctx)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if _, err := client.Post.Update().Where(dbPost.HasCategoryWith(dbCategory.IDEQ(category.ID))).SetNillableCategoryID(nil).Save(ctx); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}
