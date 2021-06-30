package admin

import (
	"context"
	"devlog/common"
	"devlog/ent"
	dbAdmin "devlog/ent/admin"
	dbCategory "devlog/ent/category"
	dbUnsavedPost "devlog/ent/unsavedpost"
	dbUnsavedPostImage "devlog/ent/unsavedpostimage"
	dbUnsavedPostThumbnail "devlog/ent/unsavedpostthumbnail"
	"devlog/env"
	"devlog/middleware"
	"devlog/regex"
	"devlog/util"
	"fmt"
	awsS3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
	"time"
)

func AttachUnsavedPost(group *echo.Group) {
	group.GET("", ListUnsavedPosts, middleware.WithSession, middleware.RequireSession)
	group.GET("/:post", GetUnsavedPost, middleware.WithSession, middleware.RequireSession)
	group.POST("", NewUnsavedPost, middleware.WithSession, middleware.RequireSession)
	group.PUT("/:post", UpdateUnsavedPost, middleware.WithSession, middleware.RequireSession)
	group.DELETE("/:uuid", DeleteUnsavedPost, middleware.WithSession, middleware.RequireSession)
	group.POST("/:post/thumbnail", NewUnsavedPostThumbnail)
	group.GET("/:post/thumbnail", GetUnsavedPostThumbnail, middleware.WithSession, middleware.RequireSession)
	group.PUT("/:post/thumbnail", SetUnsavedPostThumbnail)
	group.POST("/:post/images", NewUnsavedPostImage)
	group.GET("/:post/images/:image", GetUnsavedPostImage, middleware.WithSession, middleware.RequireSession)
	group.PUT("/:post/images/:image", SetUnsavedPostImage)
}

func ListUnsavedPosts(c echo.Context) error {
	ctx := c.(*common.Context)

	posts, err := ctx.Client().UnsavedPost.Query().
		Where(dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID))).
		Select(
			dbUnsavedPost.FieldUUID,
			dbUnsavedPost.FieldSlug,
			dbUnsavedPost.FieldAccessLevel,
			dbUnsavedPost.FieldTitle,
			dbUnsavedPost.FieldCreatedAt,
			dbUnsavedPost.FieldModifiedAt).
		WithCategory(func(query *ent.CategoryQuery) {
			query.Select(
				dbCategory.FieldName)
		}).
		WithThumbnail(func(query *ent.UnsavedPostThumbnailQuery) {
			query.Select(
				dbUnsavedPostThumbnail.FieldValidity,
				dbUnsavedPostThumbnail.FieldWidth,
				dbUnsavedPostThumbnail.FieldHeight,
				dbUnsavedPostThumbnail.FieldHash,
				dbUnsavedPostThumbnail.FieldURL,
				dbUnsavedPostThumbnail.FieldCreatedAt)
		}).
		Order(ent.Asc(dbUnsavedPost.FieldCreatedAt)).
		All(context.Background())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	type ThumbnailJSON struct {
		Validity  string    `json:"validity"`
		Width     *uint32   `json:"width"`
		Height    *uint32   `json:"height"`
		Hash      *string   `json:"hash"`
		URL       *string   `json:"url"`
		CreatedAt time.Time `json:"created-at"`
	}
	type PostJSON struct {
		UUID        string         `json:"uuid"`
		Slug        *string        `json:"slug"`
		AccessLevel *string        `json:"access-level"`
		Title       *string        `json:"title"`
		CreatedAt   time.Time      `json:"created-at"`
		ModifiedAt  time.Time      `json:"modified-at"`
		Category    *string        `json:"category"`
		Thumbnail   *ThumbnailJSON `json:"thumbnail"`
	}

	postJSONs := make([]PostJSON, len(posts))

	for index, post := range posts {
		var category *string

		if post.Edges.Category != nil {
			category = &post.Edges.Category.Name
		}

		var thumbnail *ThumbnailJSON

		if post.Edges.Thumbnail != nil {
			thumbnail = &ThumbnailJSON{
				Validity:  (string)(post.Edges.Thumbnail.Validity),
				Width:     post.Edges.Thumbnail.Width,
				Height:    post.Edges.Thumbnail.Height,
				Hash:      post.Edges.Thumbnail.Hash,
				URL:       post.Edges.Thumbnail.URL,
				CreatedAt: post.Edges.Thumbnail.CreatedAt,
			}
		}

		postJSONs[index] = PostJSON{
			UUID:        post.UUID,
			Slug:        post.Slug,
			AccessLevel: (*string)(post.AccessLevel),
			Title:       post.Title,
			CreatedAt:   post.CreatedAt,
			ModifiedAt:  post.ModifiedAt,
			Category:    category,
			Thumbnail:   thumbnail,
		}
	}

	return c.JSON(http.StatusOK, postJSONs)
}

func GetUnsavedPost(c echo.Context) error {
	type PostInfo struct {
		PostUUID string `param:"post" validate:"required,hexadecimal,len=64"`
	}

	postInfo := new(PostInfo)

	if err := c.Bind(postInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(postInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	ctx := c.(*common.Context)

	post, err := ctx.Client().UnsavedPost.Query().
		Where(
			dbUnsavedPost.And(
				dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID)),
				dbUnsavedPost.UUIDEQ(postInfo.PostUUID))).
		Select(
			dbUnsavedPost.FieldUUID,
			dbUnsavedPost.FieldSlug,
			dbUnsavedPost.FieldAccessLevel,
			dbUnsavedPost.FieldTitle,
			dbUnsavedPost.FieldContent,
			dbUnsavedPost.FieldCreatedAt,
			dbUnsavedPost.FieldModifiedAt).
		WithCategory(func(query *ent.CategoryQuery) {
			query.Select(
				dbCategory.FieldName)
		}).
		WithThumbnail(func(query *ent.UnsavedPostThumbnailQuery) {
			query.Select(
				dbUnsavedPostThumbnail.FieldValidity,
				dbUnsavedPostThumbnail.FieldWidth,
				dbUnsavedPostThumbnail.FieldHeight,
				dbUnsavedPostThumbnail.FieldHash,
				dbUnsavedPostThumbnail.FieldURL,
				dbUnsavedPostThumbnail.FieldCreatedAt)
		}).
		WithImages(func(query *ent.UnsavedPostImageQuery) {
			query.Select(
				dbUnsavedPostImage.FieldUUID,
				dbUnsavedPostImage.FieldValidity,
				dbUnsavedPostImage.FieldWidth,
				dbUnsavedPostImage.FieldHeight,
				dbUnsavedPostImage.FieldHash,
				dbUnsavedPostImage.FieldTitle,
				dbUnsavedPostImage.FieldURL,
				dbUnsavedPostImage.FieldCreatedAt).
				Order(
					ent.Asc(dbUnsavedPostImage.FieldCreatedAt),
					ent.Asc(dbUnsavedPostImage.FieldUUID))
		}).
		First(context.Background())

	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	type ThumbnailJSON struct {
		Validity  string    `json:"validity"`
		Width     *uint32   `json:"width"`
		Height    *uint32   `json:"height"`
		Hash      *string   `json:"hash"`
		URL       *string   `json:"url"`
		CreatedAt time.Time `json:"created-at"`
	}
	type ImageJSON struct {
		UUID      string    `json:"uuid"`
		Validity  string    `json:"validity"`
		Width     *uint32   `json:"width"`
		Height    *uint32   `json:"height"`
		Hash      *string   `json:"hash"`
		Title     *string   `json:"title"`
		URL       *string   `json:"url"`
		CreatedAt time.Time `json:"created-at"`
	}
	type PostJSON struct {
		UUID        string         `json:"uuid"`
		Slug        *string        `json:"slug"`
		AccessLevel *string        `json:"access-level"`
		Title       *string        `json:"title"`
		Content     *string        `json:"content"`
		CreatedAt   time.Time      `json:"created-at"`
		ModifiedAt  time.Time      `json:"modified-at"`
		Category    *string        `json:"category"`
		Thumbnail   *ThumbnailJSON `json:"thumbnail"`
		Images      []ImageJSON    `json:"images"`
	}

	var category *string

	if post.Edges.Category != nil {
		category = &post.Edges.Category.Name
	}

	var thumbnail *ThumbnailJSON

	if post.Edges.Thumbnail != nil {
		thumbnail = &ThumbnailJSON{
			Validity:  (string)(post.Edges.Thumbnail.Validity),
			Width:     post.Edges.Thumbnail.Width,
			Height:    post.Edges.Thumbnail.Height,
			Hash:      post.Edges.Thumbnail.Hash,
			URL:       post.Edges.Thumbnail.URL,
			CreatedAt: post.Edges.Thumbnail.CreatedAt,
		}
	}

	images := make([]ImageJSON, len(post.Edges.Images))

	for index, image := range post.Edges.Images {
		images[index] = ImageJSON{
			UUID:      image.UUID,
			Validity:  (string)(image.Validity),
			Width:     image.Width,
			Height:    image.Height,
			Hash:      image.Hash,
			Title:     image.Title,
			URL:       image.URL,
			CreatedAt: image.CreatedAt,
		}
	}

	return c.JSON(http.StatusOK, PostJSON{
		UUID:        post.UUID,
		Slug:        post.Slug,
		AccessLevel: (*string)(post.AccessLevel),
		Title:       post.Title,
		Content:     post.Content,
		CreatedAt:   post.CreatedAt,
		ModifiedAt:  post.ModifiedAt,
		Category:    category,
		Thumbnail:   thumbnail,
		Images:      images,
	})
}

func NewUnsavedPost(c echo.Context) error {
	uuid, err := util.GenerateToken64()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	ctx := c.(*common.Context)

	_, err = ctx.Client().UnsavedPost.Create().
		SetUUID(uuid).
		SetAuthor(ctx.Admin()).
		Save(context.Background())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	type PostJSON struct {
		UUID string `json:"uuid"`
	}

	return c.JSON(http.StatusCreated, PostJSON{UUID: uuid})
}

func UpdateUnsavedPost(c echo.Context) error {
	type PostInfo struct {
		PostUUID    string                     `param:"post" validate:"required,hexadecimal,len=64"`
		Slug        *string                    `json:"slug" form:"slug" validate:"omitempty,min=1,max=255"`
		AccessLevel *dbUnsavedPost.AccessLevel `json:"access-level" form:"access-level" query:"access-level" validate:"omitempty,oneof=public unlisted private"`
		Title       *string                    `json:"title" form:"title" query:"title" validate:"omitempty,min=1,max=255"`
		Content     *string                    `json:"content" form:"content" query:"content" validate:"omitempty,min=1"`
		Category    *string                    `json:"category" form:"category" query:"category" validate:"omitempty,min=1"`
	}

	postInfo := new(PostInfo)

	if err := c.Bind(postInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(postInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if postInfo.Slug != nil && !regex.Slug.MatchString(*postInfo.Slug) {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		post, err := tx.UnsavedPost.Query().
			Where(dbUnsavedPost.And(
				dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID)),
				dbUnsavedPost.UUIDEQ(postInfo.PostUUID))).
			Select(dbUnsavedPost.FieldID).
			First(context.Background())

		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, echo.NewHTTPError(http.StatusNotFound)
			}
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		var categoryID *int

		if postInfo.Category != nil {
			category, err := tx.Category.Query().
				Where(dbCategory.NameEQ(*postInfo.Category)).
				Select(dbCategory.FieldID).
				First(context.Background())

			if err != nil {
				if _, ok := err.(*ent.NotFoundError); ok {
					return nil, echo.NewHTTPError(http.StatusBadRequest)
				}
				return nil, echo.NewHTTPError(http.StatusInternalServerError)
			}

			categoryID = &category.ID
		}

		if _, err := post.Update().
			ClearSlug().
			ClearAccessLevel().
			ClearTitle().
			ClearContent().
			ClearCategory().
			SetNillableSlug(postInfo.Slug).
			SetNillableAccessLevel(postInfo.AccessLevel).
			SetNillableTitle(postInfo.Title).
			SetNillableContent(postInfo.Content).
			SetNillableCategoryID(categoryID).
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
	type PostInfo struct {
		PostUUID string `param:"post" validate:"required,hexadecimal,len=64"`
	}

	postInfo := new(PostInfo)

	if err := c.Bind(postInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(postInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		post, err := tx.UnsavedPost.Query().
			Where(dbUnsavedPost.And(
				dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID)),
				dbUnsavedPost.UUIDEQ(postInfo.PostUUID))).
			Select(
				dbUnsavedPost.FieldID,
				dbUnsavedPost.FieldUUID).
			WithThumbnail(func(query *ent.UnsavedPostThumbnailQuery) {
				query.Select(
					dbUnsavedPostThumbnail.FieldValidity)
			}).
			WithImages(func(query *ent.UnsavedPostImageQuery) {
				query.Select(
					dbUnsavedPostImage.FieldUUID,
					dbUnsavedPostImage.FieldValidity)
			}).
			First(context.Background())

		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, echo.NewHTTPError(http.StatusNotFound)
			}
			return nil, err
		}

		var deletionWaitGroup sync.WaitGroup

		if post.Edges.Thumbnail != nil && post.Edges.Thumbnail.Validity == dbUnsavedPostThumbnail.ValidityValid {
			deletionWaitGroup.Add(1)

			go func() {
				defer deletionWaitGroup.Done()

				key := fmt.Sprintf("v1/%v/thumbnail", post.UUID)
				_, err := ctx.AWSS3Client().DeleteObject(context.Background(), &awsS3.DeleteObjectInput{
					Bucket: &env.AWSS3Bucket,
					Key:    &key,
				})

				if err != nil {
					ctx.Logger().Warnf("unable to remove the thumbnail of the post %v from the AWS S3 due to: %v", post.UUID, err)
				}
			}()
		}

		for _, image := range post.Edges.Images {
			if image.Validity == dbUnsavedPostImage.ValidityValid {
				deletionWaitGroup.Add(1)

				go func(image *ent.UnsavedPostImage) {
					defer deletionWaitGroup.Done()

					key := fmt.Sprintf("v1/%v/images/%v", post.UUID, image.UUID)
					_, err := ctx.AWSS3Client().DeleteObject(context.Background(), &awsS3.DeleteObjectInput{
						Bucket: &env.AWSS3Bucket,
						Key:    &key,
					})

					if err != nil {
						ctx.Logger().Warnf("unable to remove the image %v of the post %v from the AWS S3 due to: %v", image.UUID, post.UUID, err)
					}
				}(image)
			}
		}

		deletionWaitGroup.Wait()

		return nil, tx.UnsavedPost.DeleteOne(post).Exec(context.Background())
	})

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func NewUnsavedPostThumbnail(c echo.Context) error {
	type PostInfo struct {
		PostUUID string `param:"post" validate:"required,hexadecimal,len=64"`
	}

	postInfo := new(PostInfo)

	if err := c.Bind(postInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(postInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		post, err := tx.UnsavedPost.Query().
			Where(dbUnsavedPost.UUIDEQ(postInfo.PostUUID)).
			Select(dbUnsavedPost.FieldID).
			First(context.Background())

		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, echo.NewHTTPError(http.StatusNotFound)
			}
			return nil, err
		}

		return tx.UnsavedPostThumbnail.Create().
			SetUnsavedPost(post).
			Save(context.Background())
	})

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func GetUnsavedPostThumbnail(c echo.Context) error {
	type ThumbnailInfo struct {
		PostUUID string `param:"post" validate:"required,hexadecimal,len=64"`
	}

	thumbnailInfo := new(ThumbnailInfo)

	if err := c.Bind(thumbnailInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(thumbnailInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	client := c.(*common.Context).Client()

	thumbnail, err := client.UnsavedPostThumbnail.Query().
		Where(dbUnsavedPostThumbnail.HasUnsavedPostWith(dbUnsavedPost.UUIDEQ(thumbnailInfo.PostUUID))).
		Select(
			dbUnsavedPostThumbnail.FieldValidity,
			dbUnsavedPostThumbnail.FieldWidth,
			dbUnsavedPostThumbnail.FieldHeight,
			dbUnsavedPostThumbnail.FieldHash,
			dbUnsavedPostThumbnail.FieldURL,
			dbUnsavedPostThumbnail.FieldCreatedAt).
		First(context.Background())

	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	type ThumbnailJSON struct {
		Validity string  `json:"validity"`
		Width    *uint32 `json:"width"`
		Height   *uint32 `json:"height"`
		Hash     *string `json:"hash"`
		URL      *string `json:"url"`
	}

	return c.JSON(http.StatusCreated, ThumbnailJSON{
		Validity: string(thumbnail.Validity),
		Width:    thumbnail.Width,
		Height:   thumbnail.Height,
		Hash:     thumbnail.Hash,
		URL:      thumbnail.URL,
	})
}

func SetUnsavedPostThumbnail(c echo.Context) error {
	type ThumbnailInfo struct {
		PostUUID string  `param:"post" validate:"required,hexadecimal,len=64"`
		Validity string  `json:"validity" validate:"required,oneof=valid invalid"`
		Width    *uint32 `json:"width" validate:"required"`
		Height   *uint32 `json:"height" validate:"required"`
		Hash     *string `json:"hash" validate:"required,min=1"`
		URL      *string `json:"url" validate:"required,min=1"`
	}

	thumbnailInfo := new(ThumbnailInfo)

	if err := c.Bind(thumbnailInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(thumbnailInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		thumbnail, err := tx.UnsavedPostThumbnail.Query().
			Where(dbUnsavedPostThumbnail.HasUnsavedPostWith(dbUnsavedPost.UUIDEQ(thumbnailInfo.PostUUID))).
			Select(dbUnsavedPostThumbnail.FieldID, dbUnsavedPostThumbnail.FieldValidity).
			First(context.Background())

		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, echo.NewHTTPError(http.StatusNotFound)
			}
			return nil, err
		}

		if thumbnail.Validity != dbUnsavedPostThumbnail.ValidityPending {
			return nil, echo.NewHTTPError(http.StatusBadRequest)
		}

		query := thumbnail.Update()

		if thumbnailInfo.Validity == "valid" {
			if thumbnailInfo.Width == nil ||
				thumbnailInfo.Height == nil ||
				thumbnailInfo.Hash == nil ||
				thumbnailInfo.URL == nil {
				return nil, echo.NewHTTPError(http.StatusBadRequest)
			}

			query.SetValidity(dbUnsavedPostThumbnail.ValidityValid).
				SetWidth(*thumbnailInfo.Width).
				SetHeight(*thumbnailInfo.Height).
				SetHash(*thumbnailInfo.Hash).
				SetURL(*thumbnailInfo.URL)
		} else {
			query.SetValidity(dbUnsavedPostThumbnail.ValidityInvalid)
		}

		_, err = query.Save(context.Background())

		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func NewUnsavedPostImage(c echo.Context) error {
	type PostInfo struct {
		PostUUID string `param:"post" validate:"required,hexadecimal,len=64"`
	}

	postInfo := new(PostInfo)

	if err := c.Bind(postInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(postInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	image, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		unsavedPost, err := tx.UnsavedPost.Query().
			Where(dbUnsavedPost.UUIDEQ(postInfo.PostUUID)).
			Select(dbUnsavedPost.FieldID).
			First(context.Background())

		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, echo.NewHTTPError(http.StatusNotFound)
			}
			return nil, err
		}

		uuid, err := util.GenerateToken64()

		if err != nil {
			return nil, err
		}

		return tx.UnsavedPostImage.Create().
			SetUUID(uuid).
			SetUnsavedPost(unsavedPost).
			Save(context.Background())
	})

	if err != nil {
		return err
	}

	type ImageUUIDInfo struct {
		UUID string `json:"uuid"`
	}

	return c.JSON(http.StatusCreated, ImageUUIDInfo{UUID: image.(*ent.UnsavedPostImage).UUID})
}

func GetUnsavedPostImage(c echo.Context) error {
	type ImageInfo struct {
		PostUUID  string `param:"post" validate:"required,hexadecimal,len=64"`
		ImageUUID string `param:"image" validate:"required,hexadecimal,len=64"`
	}

	imageInfo := new(ImageInfo)

	if err := c.Bind(imageInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(imageInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	client := c.(*common.Context).Client()

	image, err := client.UnsavedPostImage.Query().
		Where(dbUnsavedPostImage.And(
			dbUnsavedPostImage.HasUnsavedPostWith(dbUnsavedPost.UUIDEQ(imageInfo.PostUUID)),
			dbUnsavedPostImage.UUIDEQ(imageInfo.ImageUUID))).
		Select(
			dbUnsavedPostImage.FieldValidity,
			dbUnsavedPostImage.FieldWidth,
			dbUnsavedPostImage.FieldHeight,
			dbUnsavedPostImage.FieldHash,
			dbUnsavedPostImage.FieldURL,
			dbUnsavedPostImage.FieldCreatedAt).
		First(context.Background())

	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	type ImageJSON struct {
		Validity string  `json:"validity"`
		Width    *uint32 `json:"width"`
		Height   *uint32 `json:"height"`
		Hash     *string `json:"hash"`
		URL      *string `json:"url"`
	}

	return c.JSON(http.StatusCreated, ImageJSON{
		Validity: string(image.Validity),
		Width:    image.Width,
		Height:   image.Height,
		Hash:     image.Hash,
		URL:      image.URL,
	})
}

func SetUnsavedPostImage(c echo.Context) error {
	type ImageInfo struct {
		PostUUID  string  `param:"post" validate:"required,hexadecimal,len=64"`
		ImageUUID string  `param:"image" validate:"required,hexadecimal,len=64"`
		Validity  string  `json:"validity" validate:"required,oneof=valid invalid"`
		Width     *uint32 `json:"width" validate:"required,min=1"`
		Height    *uint32 `json:"height" validate:"required,min=1"`
		Hash      *string `json:"hash" validate:"required,min=1"`
		Title     *string `json:"title" validate:"required,min=1"`
		URL       *string `json:"url" validate:"required,min=1"`
	}

	imageInfo := new(ImageInfo)

	if err := c.Bind(imageInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(imageInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		image, err := tx.UnsavedPostImage.Query().
			Where(dbUnsavedPostImage.And(
				dbUnsavedPostImage.HasUnsavedPostWith(dbUnsavedPost.UUIDEQ(imageInfo.PostUUID)),
				dbUnsavedPostImage.UUIDEQ(imageInfo.ImageUUID))).
			Select(dbUnsavedPostImage.FieldID, dbUnsavedPostImage.FieldValidity).
			First(context.Background())

		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, echo.NewHTTPError(http.StatusNotFound)
			}
			return nil, err
		}

		if image.Validity != dbUnsavedPostImage.ValidityPending {
			return nil, echo.NewHTTPError(http.StatusBadRequest)
		}

		query := image.Update()

		if imageInfo.Validity == "valid" {
			if imageInfo.Width == nil ||
				imageInfo.Height == nil ||
				imageInfo.Hash == nil ||
				imageInfo.Title == nil ||
				imageInfo.URL == nil {
				return nil, echo.NewHTTPError(http.StatusBadRequest)
			}

			query.SetValidity(dbUnsavedPostImage.ValidityValid).
				SetWidth(*imageInfo.Width).
				SetHeight(*imageInfo.Height).
				SetHash(*imageInfo.Hash).
				SetTitle(*imageInfo.Title).
				SetURL(*imageInfo.URL)
		} else {
			query.SetValidity(dbUnsavedPostImage.ValidityInvalid)
		}

		_, err = query.Save(context.Background())

		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
