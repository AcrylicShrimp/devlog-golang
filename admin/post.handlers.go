package admin

import (
	"bytes"
	"devlog/common"
	"devlog/ent"
	dbCategory "devlog/ent/category"
	dbPost "devlog/ent/post"
	"devlog/markdown"
	"devlog/regex"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
	"net/http"
	"strings"
)

func AttachPost(group *echo.Group) {
	group.POST("", NewPostHandler)
}

func NewPostHandler(c echo.Context) error {
	type PostInfo struct {
		Category    string `json:"category" form:"category" query:"category"`
		Slug        string `json:"slug" form:"slug" query:"slug" validate:"required,max=255"`
		AccessLevel string `json:"access-level" form:"access-level" query:"access-level" validate:"required,oneof=public unlisted private"`
		Title       string `json:"title" form:"title" query:"title" validate:"required,max=255"`
		Content     string `json:"content" form:"content" query:"content" validate:"required"`
	}

	postInfo := new(PostInfo)
	if err := c.Bind(postInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(postInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	client := c.(*common.Context).Client()
	ctx := c.(*common.Context).Ctx()

	var categoryId *int
	if postInfo.Category != "" {
		category, err := client.Category.Query().Where(dbCategory.NameEQ(postInfo.Category)).Select(dbCategory.FieldID).First(ctx)
		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return echo.NewHTTPError(http.StatusBadRequest, err)
			}
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		categoryId = &category.ID
	}

	var accessLevel dbPost.AccessLevel
	switch postInfo.AccessLevel {
	case string(dbPost.AccessLevelPublic):
		accessLevel = dbPost.AccessLevelPublic
	case string(dbPost.AccessLevelUnlisted):
		accessLevel = dbPost.AccessLevelUnlisted
	case string(dbPost.AccessLevelPrivate):
		accessLevel = dbPost.AccessLevelPrivate
	}

	markdownHTML := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)
	markdownPlain := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithRendererOptions(renderer.WithNodeRenderers(util.Prioritized(markdown.NewPureMarkdownRenderer(), 1))),
	)

	var htmlContentBuf bytes.Buffer
	if err := markdownHTML.Convert([]byte(postInfo.Content), &htmlContentBuf); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	htmlContent := htmlContentBuf.String()

	var plainContentBuf bytes.Buffer
	if err := markdownPlain.Convert([]byte(postInfo.Content), &plainContentBuf); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	plainContent := plainContentBuf.String()
	plainContent = strings.TrimSpace(plainContent)
	plainContent = regex.Whitespaces.ReplaceAllString(plainContent, " ")

	form, err := c.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(""),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	uploader := s3manager.NewUploader(sess)

	if images, ok := form.File["images"]; ok {
		for _, image := range images {
			file, err := image.Open()
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError)
			}

			go func(uploader *s3manager.Uploader) {
				uploader.Upload(&s3manager.UploadInput{
					Bucket: aws.String(""),
					Key:    aws.String(""),
				})
			}(uploader)
			go func() {

			}()

			file.Close()
		}
	}
	//if formVideos, ok := form.File["videos"]; ok {
	//	videos = formVideos
	//}
	//if formAttachments, ok := form.File["attachments"]; ok {
	//	attachments = formAttachments
	//}

	post, err := client.Post.Create().SetSlug(postInfo.Slug).SetAccessLevel(accessLevel).SetTitle(postInfo.Title).SetContent(htmlContent).SetPreviewContent(string([]rune(plainContent)[:255])).SetNillableCategoryID(categoryId).Save(ctx)
	if err != nil {
		if _, ok := err.(*ent.ConstraintError); ok {
			return echo.NewHTTPError(http.StatusConflict)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	type Post struct {
		Slug string `json:"slug"`
	}

	//return c.JSON(http.StatusCreated, Post{Slug: post.Slug})
	return c.NoContent(http.StatusCreated)
}
