package handler

import (
	"context"
	"devlog/common"
	"devlog/ent"
	dbCategory "devlog/ent/category"
	dbPost "devlog/ent/post"
	dbPostImage "devlog/ent/postimage"
	dbPostThumbnail "devlog/ent/postthumbnail"
	"devlog/middleware"
	"devlog/regex"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func AttachPost(group *echo.Group) {
	group.GET("", ListPosts, middleware.WithSession)
	group.GET("/:slug", GetPost, middleware.WithSession)
}

func ListPosts(c echo.Context) error {
	ctx := c.(*common.Context)

	posts, err := ctx.Client().Post.Query().
		Select(
			dbPost.FieldSlug,
			dbPost.FieldAccessLevel,
			dbPost.FieldTitle,
			dbPost.FieldPreviewContent,
			dbPost.FieldCreatedAt,
			dbPost.FieldModifiedAt).
		WithCategory(func(query *ent.CategoryQuery) {
			query.Select(
				dbCategory.FieldName)
		}).
		WithThumbnail(func(query *ent.PostThumbnailQuery) {
			query.Select(
				dbPostThumbnail.FieldWidth,
				dbPostThumbnail.FieldHeight,
				dbPostThumbnail.FieldHash,
				dbPostThumbnail.FieldURL)
		}).
		Order(ent.Asc(dbPost.FieldCreatedAt)).
		All(context.Background())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	type PostJSON struct {
		Slug           string             `json:"slug"`
		AccessLevel    string             `json:"access-level"`
		Title          string             `json:"title"`
		PreviewContent string             `json:"preview-content"`
		CreatedAt      time.Time          `json:"created-at"`
		ModifiedAt     time.Time          `json:"modified-at"`
		Category       *string            `json:"category"`
		Thumbnail      *ent.PostThumbnail `json:"thumbnail"`
	}

	postsJSON := make([]PostJSON, len(posts))
	for index, post := range posts {
		var category *string
		if post.Edges.Category != nil {
			category = &post.Edges.Category.Name
		}

		postsJSON[index] = PostJSON{
			Slug:           post.Slug,
			AccessLevel:    (string)(post.AccessLevel),
			Title:          post.Title,
			PreviewContent: post.PreviewContent,
			CreatedAt:      post.CreatedAt,
			ModifiedAt:     post.ModifiedAt,
			Category:       category,
			Thumbnail:      post.Edges.Thumbnail,
		}
	}

	return c.JSON(http.StatusOK, postsJSON)
}

func GetPost(c echo.Context) error {
	type SlugInfo struct {
		Slug string `param:"slug" validate:"required,min=1,max=255"`
	}

	slugInfo := new(SlugInfo)
	if err := c.Bind(slugInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(slugInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if !regex.Slug.MatchString(slugInfo.Slug) {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	ctx := c.(*common.Context)

	post, err := ctx.Client().Post.Query().
		Select(
			dbPost.FieldSlug,
			dbPost.FieldAccessLevel,
			dbPost.FieldTitle,
			dbPost.FieldHTMLContent,
			dbPost.FieldPreviewContent,
			dbPost.FieldCreatedAt,
			dbPost.FieldModifiedAt).
		WithCategory(func(query *ent.CategoryQuery) {
			query.Select(
				dbCategory.FieldName)
		}).
		WithThumbnail(func(query *ent.PostThumbnailQuery) {
			query.Select(
				dbPostThumbnail.FieldWidth,
				dbPostThumbnail.FieldHeight,
				dbPostThumbnail.FieldHash,
				dbPostThumbnail.FieldURL)
		}).
		WithImages(func(query *ent.PostImageQuery) {
			query.Select(
				dbPostImage.FieldWidth,
				dbPostImage.FieldHeight,
				dbPostImage.FieldHash,
				dbPostImage.FieldTitle,
				dbPostImage.FieldURL)
		}).
		Order(ent.Asc(dbPost.FieldCreatedAt)).
		First(context.Background())
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	type PostJSON struct {
		Slug           string             `json:"slug"`
		AccessLevel    string             `json:"access-level"`
		Title          string             `json:"title"`
		HTMLContent    string             `json:"html-content"`
		PreviewContent string             `json:"preview-content"`
		CreatedAt      time.Time          `json:"created-at"`
		ModifiedAt     time.Time          `json:"modified-at"`
		Category       *string            `json:"category"`
		Thumbnail      *ent.PostThumbnail `json:"thumbnail"`
		Images         ent.PostImages     `json:"images"`
	}

	var category *string
	if post.Edges.Category != nil {
		category = &post.Edges.Category.Name
	}

	return c.JSON(http.StatusOK, PostJSON{
		Slug:           post.Slug,
		AccessLevel:    (string)(post.AccessLevel),
		Title:          post.Title,
		HTMLContent:    post.HTMLContent,
		PreviewContent: post.PreviewContent,
		CreatedAt:      post.CreatedAt,
		ModifiedAt:     post.ModifiedAt,
		Category:       category,
		Thumbnail:      post.Edges.Thumbnail,
		Images:         post.Edges.Images,
	})
}
