package admin

import (
	"bytes"
	"context"
	"devlog/common"
	"devlog/ent"
	dbCategory "devlog/ent/category"
	dbPost "devlog/ent/post"
	dbUnsavedPost "devlog/ent/unsavedpost"
	"devlog/env"
	"devlog/markdown"
	"devlog/regex"
	"devlog/util"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/h2non/bimg"
	"github.com/labstack/echo"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	goldmarkUtil "github.com/yuin/goldmark/util"
	"mime/multipart"
	"net/http"
	"strings"
)

func AttachPost(group *echo.Group) {
	group.POST("/unsaved", NewUnsavedPost, WithSession, RequireSession)
	group.DELETE("/unsaved/:uuid", DeleteUnsavedPost, WithSession, RequireSession)

	group.POST("", NewPostHandler, WithSession, RequireSession)
}

func ListUnsavedPosts(c echo.Context) error {

}

func GetUnsavedPost(c echo.Context) error {

}

func NewUnsavedPost(c echo.Context) error {
	token, err := util.GenerateToken64()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	res, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		_, err = tx.UnsavedPost.Create().SetUUID(token).SetAuthor(ctx.Admin()).Save(context.Background())
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		type UnsavedPostInfo struct {
			UUID string `json:"uuid"`
		}

		return UnsavedPostInfo{UUID: token}, nil
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}

func DeleteUnsavedPost(c echo.Context) error {
	type UnsavedPostInfo struct {
		UUID string `param:"uuid" validate:"required,hexadecimal,length=64"`
	}

	unsavedPostInfo := new(UnsavedPostInfo)
	if err := c.Bind(unsavedPostInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(unsavedPostInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		unsavedPost, err := tx.UnsavedPost.Query().Where(dbUnsavedPost.UUIDEQ(unsavedPostInfo.UUID)).Select(dbUnsavedPost.FieldID).First(context.Background())
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

	if !regex.Slug.MatchString(postInfo.Slug) {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	client := c.(*common.Context).Client()
	ctx := c.(*common.Context).Ctx()
	admin := c.(*common.Context).Admin()

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

	form, err := c.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	var imageFiles []*multipart.FileHeader

	if images, ok := form.File["images"]; ok {
		imageFiles = images
	}
	//if formVideos, ok := form.File["videos"]; ok {
	//	videos = formVideos
	//}
	//if formAttachments, ok := form.File["attachments"]; ok {
	//	attachments = formAttachments
	//}

	type UploadResult struct {
		uuid   *string
		width  *uint32
		height *uint32
		output *s3manager.UploadOutput
		err    error
	}

	imageUploadResults := make(chan UploadResult, len(imageFiles))

	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(env.AWSS3Region),
		},
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	uploader := s3manager.NewUploader(sess)

	for _, image := range imageFiles {
		file, err := image.Open()
		if err != nil {
			imageUploadResults <- UploadResult{nil, nil, nil, nil, err}
			continue
		}

		image, err := bimg.NewImage()

		go func() {
			uuid, err := util.GenerateToken64()
			if err != nil {
				imageUploadResults <- UploadResult{&uuid, nil, err}
				return
			}

			result, err := uploader.Upload(
				&s3manager.UploadInput{
					Bucket: aws.String(env.AWSS3Bucket),
					Key:    aws.String(fmt.Sprintf("posts/%v/%v", postInfo.Slug, uuid)),
					Body:   file,
				},
			)
			if err != nil {
				imageUploadResults <- UploadResult{&uuid, nil, err}
				return
			}

			imageUploadResults <- UploadResult{&uuid, result, nil}
		}()
	}

	type Uploaded struct {
		uuid   *string
		output *s3manager.UploadOutput
	}

	uploadedImages := make([]Uploaded, 0, len(imageFiles))

	for range imageFiles {
		result := <-imageUploadResults

		if result.err != nil {
			continue
		}

		uploadedImages = append(uploadedImages, Uploaded{result.uuid, result.output})
	}

	eliminateS3Objects := func() {
		svc := s3.New(sess)
		iter := s3manager.NewDeleteListIterator(
			svc, &s3.ListObjectsInput{
				Bucket: aws.String(env.AWSS3Bucket),
				Prefix: aws.String(fmt.Sprintf("posts/%v", postInfo.Slug)),
			},
		)
		s3manager.NewBatchDeleteWithClient(svc).Delete(aws.BackgroundContext(), iter)
	}

	defer eliminateS3Objects()

	if len(uploadedImages) != len(imageFiles) {
		return echo.NewHTTPError(http.StatusInternalServerError)
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

	post, err := client.Post.Create().
		SetSlug(postInfo.Slug).
		SetAccessLevel(accessLevel).
		SetTitle(postInfo.Title).
		SetContent(htmlContent).
		SetPreviewContent(string([]rune(plainContent)[:255])).
		SetAuthorID(admin.ID).
		SetNillableCategoryID(categoryId).
		Save(ctx)
	if err != nil {
		if _, ok := err.(*ent.ConstraintError); ok {
			return echo.NewHTTPError(http.StatusConflict)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	imageBulk := make([]*ent.PostImageCreate, len(uploadedImages))
	for index, image := range uploadedImages {
		imageBulk[index] = client.PostImage.Create().
			SetUUID(*image.uuid).
			SetWidth(0).
			SetHeight(0).
			SetPost(post)
	}

	_, err = client.PostImage.CreateBulk(imageBulk...).Save(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	eliminateS3Objects = func() {}

	return c.NoContent(http.StatusCreated)
}
