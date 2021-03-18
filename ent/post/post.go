// Code generated by entc, DO NOT EDIT.

package post

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the post type in the database.
	Label = "post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldAccessLevel holds the string denoting the access_level field in the database.
	FieldAccessLevel = "access_level"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldHTMLContent holds the string denoting the html_content field in the database.
	FieldHTMLContent = "html_content"
	// FieldPreviewContent holds the string denoting the preview_content field in the database.
	FieldPreviewContent = "preview_content"
	// FieldAccumulatedImageIndex holds the string denoting the accumulated_image_index field in the database.
	FieldAccumulatedImageIndex = "accumulated_image_index"
	// FieldAccumulatedVideoIndex holds the string denoting the accumulated_video_index field in the database.
	FieldAccumulatedVideoIndex = "accumulated_video_index"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldModifiedAt holds the string denoting the modified_at field in the database.
	FieldModifiedAt = "modified_at"

	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "category"
	// EdgeThumbnail holds the string denoting the thumbnail edge name in mutations.
	EdgeThumbnail = "thumbnail"
	// EdgeImages holds the string denoting the images edge name in mutations.
	EdgeImages = "images"
	// EdgeVideos holds the string denoting the videos edge name in mutations.
	EdgeVideos = "videos"

	// Table holds the table name of the post in the database.
	Table = "posts"
	// CategoryTable is the table the holds the category relation/edge.
	CategoryTable = "posts"
	// CategoryInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoryInverseTable = "categories"
	// CategoryColumn is the table column denoting the category relation/edge.
	CategoryColumn = "category_posts"
	// ThumbnailTable is the table the holds the thumbnail relation/edge.
	ThumbnailTable = "post_thumbnails"
	// ThumbnailInverseTable is the table name for the PostThumbnail entity.
	// It exists in this package in order to avoid circular dependency with the "postthumbnail" package.
	ThumbnailInverseTable = "post_thumbnails"
	// ThumbnailColumn is the table column denoting the thumbnail relation/edge.
	ThumbnailColumn = "post_thumbnail"
	// ImagesTable is the table the holds the images relation/edge.
	ImagesTable = "post_images"
	// ImagesInverseTable is the table name for the PostImage entity.
	// It exists in this package in order to avoid circular dependency with the "postimage" package.
	ImagesInverseTable = "post_images"
	// ImagesColumn is the table column denoting the images relation/edge.
	ImagesColumn = "post_images"
	// VideosTable is the table the holds the videos relation/edge.
	VideosTable = "post_videos"
	// VideosInverseTable is the table name for the PostVideo entity.
	// It exists in this package in order to avoid circular dependency with the "postvideo" package.
	VideosInverseTable = "post_videos"
	// VideosColumn is the table column denoting the videos relation/edge.
	VideosColumn = "post_videos"
)

// Columns holds all SQL columns for post fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldSlug,
	FieldAccessLevel,
	FieldTitle,
	FieldContent,
	FieldHTMLContent,
	FieldPreviewContent,
	FieldAccumulatedImageIndex,
	FieldAccumulatedVideoIndex,
	FieldCreatedAt,
	FieldModifiedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Post type.
var ForeignKeys = []string{
	"category_posts",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// UUIDValidator is a validator for the "uuid" field. It is called by the builders before save.
	UUIDValidator func(string) error
	// SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	SlugValidator func(string) error
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// PreviewContentValidator is a validator for the "preview_content" field. It is called by the builders before save.
	PreviewContentValidator func(string) error
	// DefaultAccumulatedImageIndex holds the default value on creation for the "accumulated_image_index" field.
	DefaultAccumulatedImageIndex uint64
	// DefaultAccumulatedVideoIndex holds the default value on creation for the "accumulated_video_index" field.
	DefaultAccumulatedVideoIndex uint64
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// UpdateDefaultModifiedAt holds the default value on update for the "modified_at" field.
	UpdateDefaultModifiedAt func() time.Time
)

// AccessLevel defines the type for the "access_level" enum field.
type AccessLevel string

// AccessLevel values.
const (
	AccessLevelPublic   AccessLevel = "public"
	AccessLevelUnlisted AccessLevel = "unlisted"
	AccessLevelPrivate  AccessLevel = "private"
)

func (al AccessLevel) String() string {
	return string(al)
}

// AccessLevelValidator is a validator for the "access_level" field enum values. It is called by the builders before save.
func AccessLevelValidator(al AccessLevel) error {
	switch al {
	case AccessLevelPublic, AccessLevelUnlisted, AccessLevelPrivate:
		return nil
	default:
		return fmt.Errorf("post: invalid enum value for access_level field: %q", al)
	}
}
