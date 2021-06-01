package admin

import (
	"context"
	"devlog/common"
	"devlog/ent"
	dbAdmin "devlog/ent/admin"
	dbUnsavedPost "devlog/ent/unsavedpost"
	"devlog/regex"
	"devlog/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AttachUnsavedPost(group *echo.Group) {
	group.GET("", ListUnsavedPosts, WithSession, RequireSession)
	group.GET("/:uuid", GetUnsavedPost, WithSession, RequireSession)
	group.POST("", NewUnsavedPost, WithSession, RequireSession)
	group.PUT("/:uuid", UpdateUnsavedPost, WithSession, RequireSession)
	group.DELETE("/:uuid", DeleteUnsavedPost, WithSession, RequireSession)
}

func ListUnsavedPosts(c echo.Context) error {
	ctx := c.(*common.Context)

	unsavedPosts, err := ctx.Client().UnsavedPost.Query().
		Where(dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID))).
		Select(
			dbUnsavedPost.FieldUUID,
			dbUnsavedPost.FieldSlug,
			dbUnsavedPost.FieldAccessLevel,
			dbUnsavedPost.FieldTitle,
			dbUnsavedPost.FieldCreatedAt,
			dbUnsavedPost.FieldModifiedAt).
		Order(ent.Asc(dbUnsavedPost.FieldCreatedAt)).
		All(context.Background())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, unsavedPosts)
}

func GetUnsavedPost(c echo.Context) error {
	type UUIDInfo struct {
		UUID string `param:"uuid" validate:"required,hexadecimal,len=64"`
	}

	uuidInfo := new(UUIDInfo)
	if err := c.Bind(uuidInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(uuidInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	ctx := c.(*common.Context)

	unsavedPost, err := ctx.Client().UnsavedPost.Query().
		Where(dbUnsavedPost.And(dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID)), dbUnsavedPost.UUIDEQ(uuidInfo.UUID))).
		Select(
			dbUnsavedPost.FieldUUID,
			dbUnsavedPost.FieldSlug,
			dbUnsavedPost.FieldSlug,
			dbUnsavedPost.FieldAccessLevel,
			dbUnsavedPost.FieldTitle,
			dbUnsavedPost.FieldAccessLevel,
			dbUnsavedPost.FieldContent,
			dbUnsavedPost.FieldCreatedAt,
			dbUnsavedPost.FieldModifiedAt).
		First(context.Background())
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, unsavedPost)
}

func NewUnsavedPost(c echo.Context) error {
	token, err := util.GenerateToken64()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	ctx := c.(*common.Context)

	_, err = ctx.Client().UnsavedPost.Create().SetUUID(token).SetAuthor(ctx.Admin()).
		Save(context.Background())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	type UnsavedPostInfo struct {
		UUID string `json:"uuid"`
	}

	return c.JSON(http.StatusCreated, UnsavedPostInfo{UUID: token})
}

func UpdateUnsavedPost(c echo.Context) error {
	type UnsavedPostInfo struct {
		UUID        string                     `param:"uuid" validate:"required,hexadecimal,len=64"`
		Slug        *string                    `json:"slug" form:"slug" validate:"omitempty,min=1,max=255"`
		AccessLevel *dbUnsavedPost.AccessLevel `json:"access-level" form:"access-level" query:"access-level" validate:"omitempty,oneof=public unlisted private"`
		Title       *string                    `json:"title" form:"title" query:"title" validate:"omitempty,min=1,max=255"`
		Content     *string                    `json:"content" form:"content" query:"content" validate:"omitempty,min=1"`
	}

	unsavedPostInfo := new(UnsavedPostInfo)
	if err := c.Bind(unsavedPostInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(unsavedPostInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if unsavedPostInfo.Slug != nil && !regex.Slug.MatchString(*unsavedPostInfo.Slug) {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		unsavedPost, err := tx.UnsavedPost.Query().
			Where(dbUnsavedPost.And(
				dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID)),
				dbUnsavedPost.UUIDEQ(unsavedPostInfo.UUID))).
			Select(dbUnsavedPost.FieldID).
			First(context.Background())
		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, echo.NewHTTPError(http.StatusNotFound)
			}

			return nil, err
		}

		if _, err := unsavedPost.Update().
			ClearSlug().
			ClearAccessLevel().
			ClearTitle().
			ClearContent().
			SetNillableSlug(unsavedPostInfo.Slug).
			SetNillableAccessLevel(unsavedPostInfo.AccessLevel).
			SetNillableTitle(unsavedPostInfo.Title).
			SetNillableContent(unsavedPostInfo.Content).
			Save(context.Background()); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		return nil, nil
	})
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func DeleteUnsavedPost(c echo.Context) error {
	type UUIDInfo struct {
		UUID string `param:"uuid" validate:"required,hexadecimal,len=64"`
	}

	uuidInfo := new(UUIDInfo)
	if err := c.Bind(uuidInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(uuidInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		unsavedPost, err := tx.UnsavedPost.Query().
			Where(dbUnsavedPost.And(dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID)), dbUnsavedPost.UUIDEQ(uuidInfo.UUID))).
			Select(dbUnsavedPost.FieldID).
			First(context.Background())
		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, echo.NewHTTPError(http.StatusNotFound)
			}

			return nil, err
		}

		return nil, tx.UnsavedPost.DeleteOne(unsavedPost).Exec(context.Background())
	})
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
