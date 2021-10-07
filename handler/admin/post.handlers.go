package admin

import (
	"bytes"
	"context"
	"devlog/common"
	"devlog/ent"
	dbAdmin "devlog/ent/admin"
	dbCategory "devlog/ent/category"
	dbPost "devlog/ent/post"
	dbUnsavedPost "devlog/ent/unsavedpost"
	dbUnsavedPostImage "devlog/ent/unsavedpostimage"
	dbUnsavedPostThumbnail "devlog/ent/unsavedpostthumbnail"
	"devlog/markdown"
	"devlog/regex"
	"devlog/util"
	"github.com/labstack/echo/v4"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	goldmarkUtil "github.com/yuin/goldmark/util"
	"net/http"
	"strings"
)

func AttachPost(group *echo.Group) {
	group.POST("", NewPost)
}

func NewPost(c echo.Context) error {
	type PostInfo struct {
		PostUUID string `json:"post-uuid" validate:"required,hexadecimal,len=64"`
	}

	uuidInfo := new(PostInfo)

	if err := c.Bind(uuidInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(uuidInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		unsavedPost, err := tx.UnsavedPost.Query().
			Where(dbUnsavedPost.And(
				dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID)),
				dbUnsavedPost.UUIDEQ(uuidInfo.PostUUID))).
			Select(
				dbUnsavedPost.FieldUUID,
				dbUnsavedPost.FieldSlug,
				dbUnsavedPost.FieldAccessLevel,
				dbUnsavedPost.FieldTitle,
				dbUnsavedPost.FieldContent).
			WithCategory(func(query *ent.CategoryQuery) {
				query.Select(
					dbCategory.FieldID)
			}).
			WithThumbnail(func(query *ent.UnsavedPostThumbnailQuery) {
				query.Select(
					dbUnsavedPostThumbnail.FieldID,
					dbUnsavedPostThumbnail.FieldValidity,
					dbUnsavedPostThumbnail.FieldWidth,
					dbUnsavedPostThumbnail.FieldHeight,
					dbUnsavedPostThumbnail.FieldHash,
					dbUnsavedPostThumbnail.FieldURL)
			}).
			WithImages(func(query *ent.UnsavedPostImageQuery) {
				query.Select(
					dbUnsavedPostImage.FieldID,
					dbUnsavedPostImage.FieldValidity,
					dbUnsavedPostImage.FieldUUID,
					dbUnsavedPostImage.FieldWidth,
					dbUnsavedPostImage.FieldHeight,
					dbUnsavedPostImage.FieldHash,
					dbUnsavedPostImage.FieldTitle,
					dbUnsavedPostImage.FieldURL).
					Order(
						ent.Asc(dbUnsavedPostImage.FieldCreatedAt),
						ent.Asc(dbUnsavedPostImage.FieldUUID))
			}).
			First(context.Background())

		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, echo.NewHTTPError(http.StatusBadRequest)
			}
			return nil, err
		}

		if unsavedPost.Slug == nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest)
		}
		if unsavedPost.AccessLevel == nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest)
		}
		if unsavedPost.Title == nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest)
		}
		if unsavedPost.Content == nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest)
		}

		if unsavedPost.Edges.Thumbnail != nil &&
			unsavedPost.Edges.Thumbnail.Validity != dbUnsavedPostThumbnail.ValidityValid {
			return nil, echo.NewHTTPError(http.StatusBadRequest)
		}

		for _, image := range unsavedPost.Edges.Images {
			if image.Validity != dbUnsavedPostImage.ValidityValid {
				return nil, echo.NewHTTPError(http.StatusBadRequest)
			}
		}

		markdownHTML := goldmark.New(
			goldmark.WithExtensions(extension.GFM),
			goldmark.WithParserOptions(
				parser.WithAutoHeadingID(),
			),
		)
		markdownPlain := goldmark.New(
			goldmark.WithExtensions(extension.GFM),
			goldmark.WithRendererOptions(
				renderer.WithNodeRenderers(
					goldmarkUtil.Prioritized(
						markdown.NewPureMarkdownRenderer(), 1,
					),
				),
			),
		)

		var htmlContentBuf bytes.Buffer

		if err := markdownHTML.Convert([]byte(*unsavedPost.Content), &htmlContentBuf); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		htmlContent := htmlContentBuf.String()

		var plainContentBuf bytes.Buffer

		if err := markdownPlain.Convert([]byte(*unsavedPost.Content), &plainContentBuf); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		plainContent := plainContentBuf.String()
		plainContent = strings.TrimSpace(plainContent)
		plainContent = regex.Whitespaces.ReplaceAllString(plainContent, " ")

		if 255 < len(plainContent) {
			plainContent = string([]rune(plainContent)[:255])
		}

		var categoryID *int

		if unsavedPost.Edges.Category != nil {
			categoryID = &unsavedPost.Edges.Category.ID
		}

		post, err := tx.Post.Create().
			SetUUID(unsavedPost.UUID).
			SetSlug(*unsavedPost.Slug).
			SetAccessLevel(dbPost.AccessLevel(*unsavedPost.AccessLevel)).
			SetTitle(*unsavedPost.Title).
			SetContent(*unsavedPost.Content).
			SetHTMLContent(htmlContent).
			SetPreviewContent(plainContent).
			SetAuthor(ctx.Admin()).
			SetNillableCategoryID(categoryID).
			Save(context.Background())

		if err != nil {
			if _, ok := err.(*ent.ConstraintError); ok {
				return nil, echo.NewHTTPError(http.StatusConflict)
			}

			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		if thumbnail := unsavedPost.Edges.Thumbnail; thumbnail != nil {
			_, err := tx.PostThumbnail.Create().
				SetWidth(*thumbnail.Width).
				SetHeight(*thumbnail.Height).
				SetHash(*thumbnail.Hash).
				SetURL(*thumbnail.URL).
				SetPost(post).
				Save(context.Background())

			if err != nil {
				return nil, echo.NewHTTPError(http.StatusInternalServerError)
			}
		}

		for _, image := range unsavedPost.Edges.Images {
			_, err := tx.PostImage.Create().
				SetUUID(image.UUID).
				SetWidth(*image.Width).
				SetHeight(*image.Height).
				SetHash(*image.Hash).
				SetTitle(*image.Title).
				SetURL(*image.URL).
				SetPost(post).
				Save(context.Background())

			if err != nil {
				return nil, echo.NewHTTPError(http.StatusInternalServerError)
			}
		}

		return nil, tx.UnsavedPost.DeleteOne(unsavedPost).Exec(context.Background())
	})

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}
