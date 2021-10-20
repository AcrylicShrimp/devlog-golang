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
	"devlog/model"
	"devlog/regex"
	"devlog/util"
	"fmt"
	awsS3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

func AttachUnsavedPost(group *echo.Group) {
	group.GET("", ListUnsavedPosts)
	group.GET("/:uuid", GetUnsavedPost)
	group.POST("", NewUnsavedPost)
	group.PUT("/:uuid", UpdateUnsavedPost)
	group.DELETE("/:uuid", DeleteUnsavedPost)
	group.GET("/:uuid/thumbnail", GetUnsavedPostThumbnail)
	group.POST("/:uuid/thumbnail", NewUnsavedPostThumbnail)
	group.PUT("/:uuid/thumbnail", UpdateUnsavedPostThumbnail)
	group.DELETE("/:uuid/thumbnail", DeleteUnsavedPostThumbnail)
	group.GET("/:uuid/images", ListUnsavedPostImage)
	group.GET("/:uuid/images/:image", GetUnsavedPostImage)
	group.POST("/:uuid/images", NewUnsavedPostImage)
	group.PUT("/:uuid/images/:image", UpdateUnsavedPostImage)
	group.DELETE("/:uuid/images/:image", DeleteUnsavedPostImage)
}

// ListUnsavedPosts godoc
// @router /admin/unsaved-posts [get]
// @summary List unsaved posts
// @description Lists all unsaved posts without its images.
// @description The unsaved posts are sorted by the field 'created-at' in ascending order.
// @tags admin post management
// @produce json
// @success 200 {array} model.UnsavedPostWithoutImage
// @failure 400 {object} model.HTTPError400
// @failure 401 {object} model.HTTPError401
// @failure 500 {object} model.HTTPError500
func ListUnsavedPosts(c echo.Context) error {
	ctx := c.(*common.Context)

	posts, err := ctx.Client().UnsavedPost.Query().
		Where(
			dbUnsavedPost.HasAuthorWith(
				dbAdmin.IDEQ(ctx.Admin().ID))).
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

	postJSONs := make([]model.UnsavedPostWithoutImage, len(posts))

	for index, post := range posts {
		postJSONs[index] = model.UnsavedPostWithoutImageFromModel(post)
	}

	return c.JSON(http.StatusOK, postJSONs)
}

// GetUnsavedPost godoc
// @router /admin/unsaved-posts/{uuid} [get]
// @summary Get unsaved post
// @description Gets an unsaved post by its UUID.
// @description The unsaved post will contain images if any.
// @tags admin post management
// @param uuid path string true "An UUID of the unsaved post to be fetched"
// @produce json
// @success 200 {object} model.UnsavedPost
// @failure 400 {object} model.HTTPError400
// @failure 401 {object} model.HTTPError401
// @failure 404 {object} model.HTTPError404
// @failure 500 {object} model.HTTPError500
func GetUnsavedPost(c echo.Context) error {
	param := new(model.UnsavedPostUUIDParam)

	if err := c.Bind(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	ctx := c.(*common.Context)

	post, err := ctx.Client().UnsavedPost.Query().
		Where(
			dbUnsavedPost.And(
				dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID)),
				dbUnsavedPost.UUIDEQ(param.UUID))).
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

	return c.JSON(http.StatusOK, model.UnsavedPostFromModel(post))
}

// NewUnsavedPost godoc
// @router /admin/unsaved-posts [post]
// @summary Create unsaved post
// @description Creates a new unsaved post.
// @description UUIDs are guaranteed to be unique across all unsaved posts.
// @tags admin post management
// @produce json
// @success 201 {object} model.UnsavedPostUUIDOnly
// @failure 400 {object} model.HTTPError400
// @failure 401 {object} model.HTTPError401
// @failure 500 {object} model.HTTPError500
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

	return c.JSON(http.StatusCreated, model.UnsavedPostUUIDOnlyFromUUID(uuid))
}

// UpdateUnsavedPost godoc
// @router /admin/unsaved-posts/{uuid} [put]
// @summary Update unsaved post
// @description Updates an unsaved post by its UUID.
// @tags admin post management
// @accept json
// @param uuid path string true "An UUID of the unsaved post to be updated"
// @param unsaved-post body model.UnsavedPostParam true "The unsaved post to be updated"
// @produce json
// @success 200 "NoContent: when the unsaved post has been updated successfully"
// @failure 400 {object} model.HTTPError400
// @failure 401 {object} model.HTTPError401
// @failure 404 {object} model.HTTPError404
// @failure 500 {object} model.HTTPError500
func UpdateUnsavedPost(c echo.Context) error {
	param := new(model.UnsavedPostParam)

	if err := c.Bind(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if param.Slug != nil && !regex.Slug.MatchString(*param.Slug) {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		post, err := tx.UnsavedPost.Query().
			Where(
				dbUnsavedPost.And(
					dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID)),
					dbUnsavedPost.UUIDEQ(param.UUID))).
			Select(dbUnsavedPost.FieldID).
			First(context.Background())

		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, echo.NewHTTPError(http.StatusNotFound)
			}
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		var categoryID *int

		if param.Category != nil {
			category, err := tx.Category.Query().
				Where(dbCategory.NameEQ(*param.Category)).
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
			SetNillableSlug(param.Slug).
			SetNillableAccessLevel(param.AccessLevel).
			SetNillableTitle(param.Title).
			SetNillableContent(param.Content).
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

// DeleteUnsavedPost godoc
// @router /admin/unsaved-posts/{uuid} [delete]
// @summary Delete unsaved post
// @description Deletes an unsaved post by its UUID.
// @tags admin post management
// @param uuid path string true "An UUID of the unsaved post to be deleted"
// @produce json
// @success 200 "NoContent: when the unsaved post has been deleted successfully"
// @failure 400 {object} model.HTTPError400
// @failure 401 {object} model.HTTPError401
// @failure 404 {object} model.HTTPError404
// @failure 500 {object} model.HTTPError500
func DeleteUnsavedPost(c echo.Context) error {
	param := new(model.UnsavedPostUUIDParam)

	if err := c.Bind(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		post, err := tx.UnsavedPost.Query().
			Where(
				dbUnsavedPost.And(
					dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID)),
					dbUnsavedPost.UUIDEQ(param.UUID))).
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

				if _, err := ctx.AWSS3Client().DeleteObject(context.Background(), &awsS3.DeleteObjectInput{
					Bucket: &env.AWSS3Bucket,
					Key:    &key,
				}); err != nil {
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

					if _, err := ctx.AWSS3Client().DeleteObject(context.Background(), &awsS3.DeleteObjectInput{
						Bucket: &env.AWSS3Bucket,
						Key:    &key,
					}); err != nil {
						ctx.Logger().Warnf("unable to remove the image %v of the post %v from the AWS S3 due to: %v", image.UUID, post.UUID, err)
					}
				}(image)
			}
		}

		deletionWaitGroup.Wait()

		if err := tx.UnsavedPost.DeleteOneID(post.ID).Exec(context.Background()); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		return nil, nil
	})

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

// GetUnsavedPostThumbnail godoc
// @router /admin/unsaved-posts/{uuid}/thumbnail [get]
// @summary Get unsaved post thumbnail
// @description Gets the thumbnail of the unsaved post.
// @tags admin post management
// @param uuid path string true "An UUID of the unsaved post to be fetched"
// @produce json
// @success 200 {object} model.UnsavedPostThumbnail
// @failure 400 {object} model.HTTPError400
// @failure 401 {object} model.HTTPError401
// @failure 404 {object} model.HTTPError404
// @failure 500 {object} model.HTTPError500
func GetUnsavedPostThumbnail(c echo.Context) error {
	param := new(model.UnsavedPostUUIDParam)

	if err := c.Bind(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	ctx := c.(*common.Context)
	client := ctx.Client()

	thumbnail, err := client.UnsavedPostThumbnail.Query().
		Where(
			dbUnsavedPostThumbnail.HasUnsavedPostWith(
				dbUnsavedPost.And(
					dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID)),
					dbUnsavedPost.UUIDEQ(param.UUID)))).
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

	return c.JSON(http.StatusOK, model.UnsavedPostThumbnailFromModel(thumbnail))
}

// NewUnsavedPostThumbnail godoc
// @router /admin/unsaved-posts/{uuid}/thumbnail [post]
// @summary Create unsaved post thumbnail
// @description Creates a new thumbnail in pending state for the unsaved post.
// @description A subsequent request must be sent to fill its state and fields.
// @tags admin post management
// @param uuid path string true "An UUID of the unsaved post to be created"
// @produce json
// @success 201 "NoContent: when the thumbnail of the unsaved post has been created successfully"
// @failure 400 {object} model.HTTPError400
// @failure 401 {object} model.HTTPError401
// @failure 404 {object} model.HTTPError404
// @failure 409 {object} model.HTTPError409
// @failure 500 {object} model.HTTPError500
func NewUnsavedPostThumbnail(c echo.Context) error {
	param := new(model.UnsavedPostUUIDParam)

	if err := c.Bind(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		post, err := tx.UnsavedPost.Query().
			Where(
				dbUnsavedPost.And(
					dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID)),
					dbUnsavedPost.UUIDEQ(param.UUID))).
			Select(dbUnsavedPost.FieldID).
			First(context.Background())

		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, echo.NewHTTPError(http.StatusNotFound)
			}
			return nil, err
		}

		if _, err := tx.UnsavedPostThumbnail.Create().
			SetUnsavedPost(post).
			Save(context.Background()); err != nil {
			if _, ok := err.(*ent.ConstraintError); ok {
				return nil, echo.NewHTTPError(http.StatusConflict)
			}
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		return nil, nil
	})

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

// UpdateUnsavedPostThumbnail godoc
// @router /admin/unsaved-posts/{uuid}/thumbnail [put]
// @summary Update unsaved post thumbnail
// @description Updates the thumbnail for the given unsaved post.
// @description The thumbnail can be either valid or invalid. The fields are required if it is valid. Ignored otherwise.
// @tags admin post management
// @accept json
// @param uuid path string true "An UUID of the unsaved post to be updated"
// @param fields body model.UnsavedPostThumbnailParam true "The thumbnail of the unsaved post to be updated"
// @produce json
// @success 200 "NoContent: when the thumbnail of the unsaved post has been updated successfully"
// @failure 400 {object} model.HTTPError400
// @failure 401 {object} model.HTTPError401
// @failure 404 {object} model.HTTPError404
// @failure 500 {object} model.HTTPError500
func UpdateUnsavedPostThumbnail(c echo.Context) error {
	param := new(model.UnsavedPostThumbnailParam)

	if err := c.Bind(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		thumbnail, err := tx.UnsavedPostThumbnail.Query().
			Where(
				dbUnsavedPostThumbnail.HasUnsavedPostWith(
					dbUnsavedPost.And(
						dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID)),
						dbUnsavedPost.UUIDEQ(param.UUID)))).
			Select(dbUnsavedPostThumbnail.FieldID, dbUnsavedPostThumbnail.FieldValidity).
			First(context.Background())

		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, echo.NewHTTPError(http.StatusNotFound)
			}
			return nil, err
		}

		query := thumbnail.Update()

		if param.Validity == "valid" {
			if param.Width == nil ||
				param.Height == nil ||
				param.Hash == nil ||
				param.URL == nil {
				return nil, echo.NewHTTPError(http.StatusBadRequest)
			}

			query.SetValidity(dbUnsavedPostThumbnail.ValidityValid).
				SetWidth(*param.Width).
				SetHeight(*param.Height).
				SetHash(*param.Hash).
				SetURL(*param.URL)
		} else {
			query.SetValidity(dbUnsavedPostThumbnail.ValidityInvalid).
				ClearWidth().
				ClearHeight().
				ClearHash().
				ClearURL()
		}

		if _, err = query.Save(context.Background()); err != nil {
			return nil, err
		}

		return nil, nil
	})

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

// DeleteUnsavedPostThumbnail godoc
// @router /admin/unsaved-posts/{uuid}/thumbnail [delete]
// @summary Delete unsaved post thumbnail
// @description Deletes the thumbnail for the given unsaved post.
// @tags admin post management
// @param uuid path string true "An UUID of the unsaved post to be deleted"
// @produce json
// @success 200 "NoContent: when the thumbnail of the unsaved post has been deleted successfully"
// @failure 400 {object} model.HTTPError400
// @failure 401 {object} model.HTTPError401
// @failure 404 {object} model.HTTPError404
// @failure 500 {object} model.HTTPError500
func DeleteUnsavedPostThumbnail(c echo.Context) error {
	param := new(model.UnsavedPostUUIDParam)

	if err := c.Bind(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		thumbnail, err := tx.UnsavedPostThumbnail.Query().
			Where(
				dbUnsavedPostThumbnail.HasUnsavedPostWith(
					dbUnsavedPost.And(
						dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID)),
						dbUnsavedPost.UUIDEQ(param.UUID)))).
			Select(
				dbUnsavedPostThumbnail.FieldID,
				dbUnsavedPostThumbnail.FieldValidity).
			First(context.Background())

		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, echo.NewHTTPError(http.StatusNotFound)
			}
			return nil, err
		}

		var deletionWaitGroup sync.WaitGroup

		if thumbnail.Validity == dbUnsavedPostThumbnail.ValidityValid {
			deletionWaitGroup.Add(1)

			go func() {
				defer deletionWaitGroup.Done()

				key := fmt.Sprintf("v1/%v/thumbnail", param.UUID)

				if _, err := ctx.AWSS3Client().DeleteObject(context.Background(), &awsS3.DeleteObjectInput{
					Bucket: &env.AWSS3Bucket,
					Key:    &key,
				}); err != nil {
					ctx.Logger().Warnf("unable to remove the thumbnail of the post %v from the AWS S3 due to: %v", param.UUID, err)
				}
			}()
		}

		deletionWaitGroup.Wait()

		if err := tx.UnsavedPostThumbnail.DeleteOneID(thumbnail.ID).Exec(context.Background()); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		return nil, nil
	})

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

// ListUnsavedPostImage godoc
// @router /admin/unsaved-posts/{uuid}/images [get]
// @summary List unsaved post images
// @description Lists all images for the given unsaved post.
// @description The images are sorted by the field 'created-at' in ascending order.
// @tags admin post management
// @produce json
// @success 200 {array} model.UnsavedPostImage
// @failure 400 {object} model.HTTPError400
// @failure 401 {object} model.HTTPError401
// @failure 500 {object} model.HTTPError500
func ListUnsavedPostImage(c echo.Context) error {
	param := new(model.UnsavedPostUUIDParam)

	if err := c.Bind(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	ctx := c.(*common.Context)

	post, err := ctx.Client().UnsavedPost.Query().
		Where(
			dbUnsavedPost.And(
				dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID)),
				dbUnsavedPost.UUIDEQ(param.UUID))).
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

	imagesJSON := make([]model.UnsavedPostImage, len(post.Edges.Images))

	for index, image := range post.Edges.Images {
		imagesJSON[index] = model.UnsavedPostImageFromModel(image)
	}

	return c.JSON(http.StatusOK, imagesJSON)
}

// GetUnsavedPostImage godoc
// @router /admin/unsaved-posts/{uuid}/images/{image} [get]
// @summary Get unsaved post image
// @description Gets the image of the unsaved post.
// @tags admin post management
// @param uuid path string true "An UUID of the unsaved post to be fetched"
// @param image path string true "An UUID of the image to be fetched"
// @produce json
// @success 200 {object} model.UnsavedPostImage
// @failure 400 {object} model.HTTPError400
// @failure 401 {object} model.HTTPError401
// @failure 404 {object} model.HTTPError404
// @failure 500 {object} model.HTTPError500
func GetUnsavedPostImage(c echo.Context) error {
	param := new(model.UnsavedPostUUIDWithImageParam)

	if err := c.Bind(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	ctx := c.(*common.Context)
	client := ctx.Client()

	image, err := client.UnsavedPostImage.Query().
		Where(
			dbUnsavedPostImage.And(
				dbUnsavedPostImage.HasUnsavedPostWith(
					dbUnsavedPost.And(
						dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID)),
						dbUnsavedPost.UUIDEQ(param.UUID))),
				dbUnsavedPostImage.UUIDEQ(param.Image))).
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

	return c.JSON(http.StatusCreated, model.UnsavedPostImageFromModel(image))
}

// NewUnsavedPostImage godoc
// @router /admin/unsaved-posts/{uuid}/images [post]
// @summary Create unsaved post image
// @description Creates a new image in pending state for the unsaved post.
// @description A subsequent request must be sent to fill its state and fields.
// @tags admin post management
// @param uuid path string true "An UUID of the unsaved post to be created"
// @produce json
// @success 201 "NoContent: when the image of the unsaved post has been created successfully"
// @failure 400 {object} model.HTTPError400
// @failure 401 {object} model.HTTPError401
// @failure 404 {object} model.HTTPError404
// @failure 500 {object} model.HTTPError500
func NewUnsavedPostImage(c echo.Context) error {
	param := new(model.UnsavedPostUUIDParam)

	if err := c.Bind(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	uuid, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		post, err := tx.UnsavedPost.Query().
			Where(
				dbUnsavedPost.And(
					dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID)),
					dbUnsavedPost.UUIDEQ(param.UUID))).
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

		if _, err := tx.UnsavedPostImage.Create().
			SetUUID(uuid).
			SetUnsavedPost(post).
			Save(context.Background()); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		return uuid, nil
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, model.UnsavedPostImageUUIDOnlyFromUUID(uuid.(string)))
}

// UpdateUnsavedPostImage godoc
// @router /admin/unsaved-posts/{uuid}/images/{image} [put]
// @summary Update unsaved post image
// @description Updates the image for the given unsaved post.
// @description The image can be either valid or invalid. The fields are required if it is valid. Ignored otherwise.
// @tags admin post management
// @accept json
// @param uuid path string true "An UUID of the unsaved post to be updated"
// @param image path string true "An UUID of the image to be fetched"
// @param fields body model.UnsavedPostImageParam true "The image of the unsaved post to be updated"
// @produce json
// @success 200 "NoContent: when the image of the unsaved post has been updated successfully"
// @failure 400 {object} model.HTTPError400
// @failure 401 {object} model.HTTPError401
// @failure 404 {object} model.HTTPError404
// @failure 500 {object} model.HTTPError500
func UpdateUnsavedPostImage(c echo.Context) error {
	param := new(model.UnsavedPostImageParam)

	if err := c.Bind(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		image, err := tx.UnsavedPostImage.Query().
			Where(dbUnsavedPostImage.And(
				dbUnsavedPostImage.HasUnsavedPostWith(dbUnsavedPost.UUIDEQ(param.UUID)),
				dbUnsavedPostImage.UUIDEQ(param.Image))).
			Select(dbUnsavedPostImage.FieldID, dbUnsavedPostImage.FieldValidity).
			First(context.Background())

		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, echo.NewHTTPError(http.StatusNotFound)
			}
			return nil, err
		}

		query := image.Update()

		if param.Validity == "valid" {
			if param.Width == nil ||
				param.Height == nil ||
				param.Hash == nil ||
				param.Title == nil ||
				param.URL == nil {
				return nil, echo.NewHTTPError(http.StatusBadRequest)
			}
			query.SetValidity(dbUnsavedPostImage.ValidityValid).
				SetWidth(*param.Width).
				SetHeight(*param.Height).
				SetHash(*param.Hash).
				SetTitle(*param.Title).
				SetURL(*param.URL)
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

// DeleteUnsavedPostImage godoc
// @router /admin/unsaved-posts/{uuid}/images/{image} [delete]
// @summary Delete unsaved post image
// @description Deletes the image for the given unsaved post.
// @tags admin post management
// @param uuid path string true "An UUID of the unsaved post to be updated"
// @param image path string true "An UUID of the image to be fetched"
// @produce json
// @success 200 "NoContent: when the image of the unsaved post has been deleted successfully"
// @failure 400 {object} model.HTTPError400
// @failure 401 {object} model.HTTPError401
// @failure 404 {object} model.HTTPError404
// @failure 500 {object} model.HTTPError500
func DeleteUnsavedPostImage(c echo.Context) error {
	param := new(model.UnsavedPostUUIDWithImageParam)

	if err := c.Bind(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := c.Validate(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	_, err := util.WithTx(c, func(ctx *common.Context, tx *ent.Tx) (interface{}, error) {
		image, err := tx.UnsavedPostImage.Query().
			Where(dbUnsavedPostImage.And(
				dbUnsavedPostImage.HasUnsavedPostWith(
					dbUnsavedPost.And(
						dbUnsavedPost.HasAuthorWith(dbAdmin.IDEQ(ctx.Admin().ID)),
						dbUnsavedPost.UUIDEQ(param.UUID))),
				dbUnsavedPostImage.UUIDEQ(param.Image))).
			Select(
				dbUnsavedPostImage.FieldID,
				dbUnsavedPostImage.FieldUUID,
				dbUnsavedPostImage.FieldValidity).
			First(context.Background())

		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, echo.NewHTTPError(http.StatusNotFound)
			}
			return nil, err
		}

		if image.Validity == dbUnsavedPostImage.ValidityValid {
			key := fmt.Sprintf("v1/%v/images/%v", param.UUID, image.UUID)

			if _, err := ctx.AWSS3Client().DeleteObject(context.Background(), &awsS3.DeleteObjectInput{
				Bucket: &env.AWSS3Bucket,
				Key:    &key,
			}); err != nil {
				ctx.Logger().Warnf("unable to remove the image %v of the post %v from the AWS S3 due to: %v", image.UUID, param.UUID, err)
			}
		}

		if err := tx.UnsavedPostImage.DeleteOneID(image.ID).Exec(context.Background()); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}

		return nil, nil
	})

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
