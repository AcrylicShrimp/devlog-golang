// Code generated by entc, DO NOT EDIT.

package ent

import (
	"devlog/ent/admin"
	"devlog/ent/adminsession"
	"devlog/ent/category"
	"devlog/ent/post"
	"devlog/ent/postattachment"
	"devlog/ent/postimage"
	"devlog/ent/postthumbnail"
	"devlog/ent/postvideo"
	"devlog/ent/schema"
	"devlog/ent/unsavedpost"
	"devlog/ent/unsavedpostattachment"
	"devlog/ent/unsavedpostimage"
	"devlog/ent/unsavedpostthumbnail"
	"devlog/ent/unsavedpostvideo"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	adminFields := schema.Admin{}.Fields()
	_ = adminFields
	// adminDescEmail is the schema descriptor for email field.
	adminDescEmail := adminFields[0].Descriptor()
	// admin.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	admin.EmailValidator = adminDescEmail.Validators[0].(func(string) error)
	// adminDescUsername is the schema descriptor for username field.
	adminDescUsername := adminFields[1].Descriptor()
	// admin.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	admin.UsernameValidator = adminDescUsername.Validators[0].(func(string) error)
	// adminDescPassword is the schema descriptor for password field.
	adminDescPassword := adminFields[2].Descriptor()
	// admin.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	admin.PasswordValidator = adminDescPassword.Validators[0].(func(string) error)
	// adminDescJoinedAt is the schema descriptor for joined_at field.
	adminDescJoinedAt := adminFields[3].Descriptor()
	// admin.DefaultJoinedAt holds the default value on creation for the joined_at field.
	admin.DefaultJoinedAt = adminDescJoinedAt.Default.(func() time.Time)
	adminsessionFields := schema.AdminSession{}.Fields()
	_ = adminsessionFields
	// adminsessionDescToken is the schema descriptor for token field.
	adminsessionDescToken := adminsessionFields[0].Descriptor()
	// adminsession.TokenValidator is a validator for the "token" field. It is called by the builders before save.
	adminsession.TokenValidator = adminsessionDescToken.Validators[0].(func(string) error)
	// adminsessionDescCreatedAt is the schema descriptor for created_at field.
	adminsessionDescCreatedAt := adminsessionFields[2].Descriptor()
	// adminsession.DefaultCreatedAt holds the default value on creation for the created_at field.
	adminsession.DefaultCreatedAt = adminsessionDescCreatedAt.Default.(func() time.Time)
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescName is the schema descriptor for name field.
	categoryDescName := categoryFields[0].Descriptor()
	// category.NameValidator is a validator for the "name" field. It is called by the builders before save.
	category.NameValidator = categoryDescName.Validators[0].(func(string) error)
	// categoryDescDescription is the schema descriptor for description field.
	categoryDescDescription := categoryFields[1].Descriptor()
	// category.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	category.DescriptionValidator = categoryDescDescription.Validators[0].(func(string) error)
	// categoryDescCreatedAt is the schema descriptor for created_at field.
	categoryDescCreatedAt := categoryFields[2].Descriptor()
	// category.DefaultCreatedAt holds the default value on creation for the created_at field.
	category.DefaultCreatedAt = categoryDescCreatedAt.Default.(func() time.Time)
	// categoryDescModifiedAt is the schema descriptor for modified_at field.
	categoryDescModifiedAt := categoryFields[3].Descriptor()
	// category.DefaultModifiedAt holds the default value on creation for the modified_at field.
	category.DefaultModifiedAt = categoryDescModifiedAt.Default.(func() time.Time)
	// category.UpdateDefaultModifiedAt holds the default value on update for the modified_at field.
	category.UpdateDefaultModifiedAt = categoryDescModifiedAt.UpdateDefault.(func() time.Time)
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescUUID is the schema descriptor for uuid field.
	postDescUUID := postFields[0].Descriptor()
	// post.UUIDValidator is a validator for the "uuid" field. It is called by the builders before save.
	post.UUIDValidator = postDescUUID.Validators[0].(func(string) error)
	// postDescSlug is the schema descriptor for slug field.
	postDescSlug := postFields[1].Descriptor()
	// post.SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	post.SlugValidator = postDescSlug.Validators[0].(func(string) error)
	// postDescTitle is the schema descriptor for title field.
	postDescTitle := postFields[3].Descriptor()
	// post.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	post.TitleValidator = postDescTitle.Validators[0].(func(string) error)
	// postDescPreviewContent is the schema descriptor for preview_content field.
	postDescPreviewContent := postFields[6].Descriptor()
	// post.PreviewContentValidator is a validator for the "preview_content" field. It is called by the builders before save.
	post.PreviewContentValidator = postDescPreviewContent.Validators[0].(func(string) error)
	// postDescCreatedAt is the schema descriptor for created_at field.
	postDescCreatedAt := postFields[7].Descriptor()
	// post.DefaultCreatedAt holds the default value on creation for the created_at field.
	post.DefaultCreatedAt = postDescCreatedAt.Default.(func() time.Time)
	// postDescModifiedAt is the schema descriptor for modified_at field.
	postDescModifiedAt := postFields[8].Descriptor()
	// post.DefaultModifiedAt holds the default value on creation for the modified_at field.
	post.DefaultModifiedAt = postDescModifiedAt.Default.(func() time.Time)
	// post.UpdateDefaultModifiedAt holds the default value on update for the modified_at field.
	post.UpdateDefaultModifiedAt = postDescModifiedAt.UpdateDefault.(func() time.Time)
	postattachmentFields := schema.PostAttachment{}.Fields()
	_ = postattachmentFields
	// postattachmentDescUUID is the schema descriptor for uuid field.
	postattachmentDescUUID := postattachmentFields[0].Descriptor()
	// postattachment.UUIDValidator is a validator for the "uuid" field. It is called by the builders before save.
	postattachment.UUIDValidator = postattachmentDescUUID.Validators[0].(func(string) error)
	// postattachmentDescName is the schema descriptor for name field.
	postattachmentDescName := postattachmentFields[2].Descriptor()
	// postattachment.NameValidator is a validator for the "name" field. It is called by the builders before save.
	postattachment.NameValidator = postattachmentDescName.Validators[0].(func(string) error)
	// postattachmentDescMime is the schema descriptor for mime field.
	postattachmentDescMime := postattachmentFields[3].Descriptor()
	// postattachment.MimeValidator is a validator for the "mime" field. It is called by the builders before save.
	postattachment.MimeValidator = postattachmentDescMime.Validators[0].(func(string) error)
	// postattachmentDescURL is the schema descriptor for url field.
	postattachmentDescURL := postattachmentFields[4].Descriptor()
	// postattachment.URLValidator is a validator for the "url" field. It is called by the builders before save.
	postattachment.URLValidator = postattachmentDescURL.Validators[0].(func(string) error)
	// postattachmentDescCreatedAt is the schema descriptor for created_at field.
	postattachmentDescCreatedAt := postattachmentFields[5].Descriptor()
	// postattachment.DefaultCreatedAt holds the default value on creation for the created_at field.
	postattachment.DefaultCreatedAt = postattachmentDescCreatedAt.Default.(func() time.Time)
	postimageFields := schema.PostImage{}.Fields()
	_ = postimageFields
	// postimageDescUUID is the schema descriptor for uuid field.
	postimageDescUUID := postimageFields[0].Descriptor()
	// postimage.UUIDValidator is a validator for the "uuid" field. It is called by the builders before save.
	postimage.UUIDValidator = postimageDescUUID.Validators[0].(func(string) error)
	// postimageDescHash is the schema descriptor for hash field.
	postimageDescHash := postimageFields[3].Descriptor()
	// postimage.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	postimage.HashValidator = postimageDescHash.Validators[0].(func(string) error)
	// postimageDescTitle is the schema descriptor for title field.
	postimageDescTitle := postimageFields[4].Descriptor()
	// postimage.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	postimage.TitleValidator = postimageDescTitle.Validators[0].(func(string) error)
	// postimageDescURL is the schema descriptor for url field.
	postimageDescURL := postimageFields[5].Descriptor()
	// postimage.URLValidator is a validator for the "url" field. It is called by the builders before save.
	postimage.URLValidator = postimageDescURL.Validators[0].(func(string) error)
	// postimageDescCreatedAt is the schema descriptor for created_at field.
	postimageDescCreatedAt := postimageFields[6].Descriptor()
	// postimage.DefaultCreatedAt holds the default value on creation for the created_at field.
	postimage.DefaultCreatedAt = postimageDescCreatedAt.Default.(func() time.Time)
	postthumbnailFields := schema.PostThumbnail{}.Fields()
	_ = postthumbnailFields
	// postthumbnailDescHash is the schema descriptor for hash field.
	postthumbnailDescHash := postthumbnailFields[2].Descriptor()
	// postthumbnail.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	postthumbnail.HashValidator = postthumbnailDescHash.Validators[0].(func(string) error)
	// postthumbnailDescURL is the schema descriptor for url field.
	postthumbnailDescURL := postthumbnailFields[3].Descriptor()
	// postthumbnail.URLValidator is a validator for the "url" field. It is called by the builders before save.
	postthumbnail.URLValidator = postthumbnailDescURL.Validators[0].(func(string) error)
	// postthumbnailDescCreatedAt is the schema descriptor for created_at field.
	postthumbnailDescCreatedAt := postthumbnailFields[4].Descriptor()
	// postthumbnail.DefaultCreatedAt holds the default value on creation for the created_at field.
	postthumbnail.DefaultCreatedAt = postthumbnailDescCreatedAt.Default.(func() time.Time)
	postvideoFields := schema.PostVideo{}.Fields()
	_ = postvideoFields
	// postvideoDescUUID is the schema descriptor for uuid field.
	postvideoDescUUID := postvideoFields[0].Descriptor()
	// postvideo.UUIDValidator is a validator for the "uuid" field. It is called by the builders before save.
	postvideo.UUIDValidator = postvideoDescUUID.Validators[0].(func(string) error)
	// postvideoDescTitle is the schema descriptor for title field.
	postvideoDescTitle := postvideoFields[1].Descriptor()
	// postvideo.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	postvideo.TitleValidator = postvideoDescTitle.Validators[0].(func(string) error)
	// postvideoDescURL is the schema descriptor for url field.
	postvideoDescURL := postvideoFields[2].Descriptor()
	// postvideo.URLValidator is a validator for the "url" field. It is called by the builders before save.
	postvideo.URLValidator = postvideoDescURL.Validators[0].(func(string) error)
	// postvideoDescCreatedAt is the schema descriptor for created_at field.
	postvideoDescCreatedAt := postvideoFields[3].Descriptor()
	// postvideo.DefaultCreatedAt holds the default value on creation for the created_at field.
	postvideo.DefaultCreatedAt = postvideoDescCreatedAt.Default.(func() time.Time)
	unsavedpostFields := schema.UnsavedPost{}.Fields()
	_ = unsavedpostFields
	// unsavedpostDescUUID is the schema descriptor for uuid field.
	unsavedpostDescUUID := unsavedpostFields[0].Descriptor()
	// unsavedpost.UUIDValidator is a validator for the "uuid" field. It is called by the builders before save.
	unsavedpost.UUIDValidator = unsavedpostDescUUID.Validators[0].(func(string) error)
	// unsavedpostDescSlug is the schema descriptor for slug field.
	unsavedpostDescSlug := unsavedpostFields[1].Descriptor()
	// unsavedpost.SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	unsavedpost.SlugValidator = unsavedpostDescSlug.Validators[0].(func(string) error)
	// unsavedpostDescTitle is the schema descriptor for title field.
	unsavedpostDescTitle := unsavedpostFields[3].Descriptor()
	// unsavedpost.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	unsavedpost.TitleValidator = unsavedpostDescTitle.Validators[0].(func(string) error)
	// unsavedpostDescCreatedAt is the schema descriptor for created_at field.
	unsavedpostDescCreatedAt := unsavedpostFields[5].Descriptor()
	// unsavedpost.DefaultCreatedAt holds the default value on creation for the created_at field.
	unsavedpost.DefaultCreatedAt = unsavedpostDescCreatedAt.Default.(func() time.Time)
	// unsavedpostDescModifiedAt is the schema descriptor for modified_at field.
	unsavedpostDescModifiedAt := unsavedpostFields[6].Descriptor()
	// unsavedpost.DefaultModifiedAt holds the default value on creation for the modified_at field.
	unsavedpost.DefaultModifiedAt = unsavedpostDescModifiedAt.Default.(func() time.Time)
	// unsavedpost.UpdateDefaultModifiedAt holds the default value on update for the modified_at field.
	unsavedpost.UpdateDefaultModifiedAt = unsavedpostDescModifiedAt.UpdateDefault.(func() time.Time)
	unsavedpostattachmentFields := schema.UnsavedPostAttachment{}.Fields()
	_ = unsavedpostattachmentFields
	// unsavedpostattachmentDescUUID is the schema descriptor for uuid field.
	unsavedpostattachmentDescUUID := unsavedpostattachmentFields[0].Descriptor()
	// unsavedpostattachment.UUIDValidator is a validator for the "uuid" field. It is called by the builders before save.
	unsavedpostattachment.UUIDValidator = unsavedpostattachmentDescUUID.Validators[0].(func(string) error)
	// unsavedpostattachmentDescName is the schema descriptor for name field.
	unsavedpostattachmentDescName := unsavedpostattachmentFields[2].Descriptor()
	// unsavedpostattachment.NameValidator is a validator for the "name" field. It is called by the builders before save.
	unsavedpostattachment.NameValidator = unsavedpostattachmentDescName.Validators[0].(func(string) error)
	// unsavedpostattachmentDescMime is the schema descriptor for mime field.
	unsavedpostattachmentDescMime := unsavedpostattachmentFields[3].Descriptor()
	// unsavedpostattachment.MimeValidator is a validator for the "mime" field. It is called by the builders before save.
	unsavedpostattachment.MimeValidator = unsavedpostattachmentDescMime.Validators[0].(func(string) error)
	// unsavedpostattachmentDescURL is the schema descriptor for url field.
	unsavedpostattachmentDescURL := unsavedpostattachmentFields[4].Descriptor()
	// unsavedpostattachment.URLValidator is a validator for the "url" field. It is called by the builders before save.
	unsavedpostattachment.URLValidator = unsavedpostattachmentDescURL.Validators[0].(func(string) error)
	// unsavedpostattachmentDescCreatedAt is the schema descriptor for created_at field.
	unsavedpostattachmentDescCreatedAt := unsavedpostattachmentFields[5].Descriptor()
	// unsavedpostattachment.DefaultCreatedAt holds the default value on creation for the created_at field.
	unsavedpostattachment.DefaultCreatedAt = unsavedpostattachmentDescCreatedAt.Default.(func() time.Time)
	unsavedpostimageFields := schema.UnsavedPostImage{}.Fields()
	_ = unsavedpostimageFields
	// unsavedpostimageDescUUID is the schema descriptor for uuid field.
	unsavedpostimageDescUUID := unsavedpostimageFields[0].Descriptor()
	// unsavedpostimage.UUIDValidator is a validator for the "uuid" field. It is called by the builders before save.
	unsavedpostimage.UUIDValidator = unsavedpostimageDescUUID.Validators[0].(func(string) error)
	// unsavedpostimageDescHash is the schema descriptor for hash field.
	unsavedpostimageDescHash := unsavedpostimageFields[3].Descriptor()
	// unsavedpostimage.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	unsavedpostimage.HashValidator = unsavedpostimageDescHash.Validators[0].(func(string) error)
	// unsavedpostimageDescTitle is the schema descriptor for title field.
	unsavedpostimageDescTitle := unsavedpostimageFields[4].Descriptor()
	// unsavedpostimage.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	unsavedpostimage.TitleValidator = unsavedpostimageDescTitle.Validators[0].(func(string) error)
	// unsavedpostimageDescURL is the schema descriptor for url field.
	unsavedpostimageDescURL := unsavedpostimageFields[5].Descriptor()
	// unsavedpostimage.URLValidator is a validator for the "url" field. It is called by the builders before save.
	unsavedpostimage.URLValidator = unsavedpostimageDescURL.Validators[0].(func(string) error)
	// unsavedpostimageDescCreatedAt is the schema descriptor for created_at field.
	unsavedpostimageDescCreatedAt := unsavedpostimageFields[6].Descriptor()
	// unsavedpostimage.DefaultCreatedAt holds the default value on creation for the created_at field.
	unsavedpostimage.DefaultCreatedAt = unsavedpostimageDescCreatedAt.Default.(func() time.Time)
	unsavedpostthumbnailFields := schema.UnsavedPostThumbnail{}.Fields()
	_ = unsavedpostthumbnailFields
	// unsavedpostthumbnailDescHash is the schema descriptor for hash field.
	unsavedpostthumbnailDescHash := unsavedpostthumbnailFields[2].Descriptor()
	// unsavedpostthumbnail.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	unsavedpostthumbnail.HashValidator = unsavedpostthumbnailDescHash.Validators[0].(func(string) error)
	// unsavedpostthumbnailDescURL is the schema descriptor for url field.
	unsavedpostthumbnailDescURL := unsavedpostthumbnailFields[3].Descriptor()
	// unsavedpostthumbnail.URLValidator is a validator for the "url" field. It is called by the builders before save.
	unsavedpostthumbnail.URLValidator = unsavedpostthumbnailDescURL.Validators[0].(func(string) error)
	// unsavedpostthumbnailDescCreatedAt is the schema descriptor for created_at field.
	unsavedpostthumbnailDescCreatedAt := unsavedpostthumbnailFields[4].Descriptor()
	// unsavedpostthumbnail.DefaultCreatedAt holds the default value on creation for the created_at field.
	unsavedpostthumbnail.DefaultCreatedAt = unsavedpostthumbnailDescCreatedAt.Default.(func() time.Time)
	unsavedpostvideoFields := schema.UnsavedPostVideo{}.Fields()
	_ = unsavedpostvideoFields
	// unsavedpostvideoDescUUID is the schema descriptor for uuid field.
	unsavedpostvideoDescUUID := unsavedpostvideoFields[0].Descriptor()
	// unsavedpostvideo.UUIDValidator is a validator for the "uuid" field. It is called by the builders before save.
	unsavedpostvideo.UUIDValidator = unsavedpostvideoDescUUID.Validators[0].(func(string) error)
	// unsavedpostvideoDescTitle is the schema descriptor for title field.
	unsavedpostvideoDescTitle := unsavedpostvideoFields[1].Descriptor()
	// unsavedpostvideo.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	unsavedpostvideo.TitleValidator = unsavedpostvideoDescTitle.Validators[0].(func(string) error)
	// unsavedpostvideoDescURL is the schema descriptor for url field.
	unsavedpostvideoDescURL := unsavedpostvideoFields[2].Descriptor()
	// unsavedpostvideo.URLValidator is a validator for the "url" field. It is called by the builders before save.
	unsavedpostvideo.URLValidator = unsavedpostvideoDescURL.Validators[0].(func(string) error)
	// unsavedpostvideoDescCreatedAt is the schema descriptor for created_at field.
	unsavedpostvideoDescCreatedAt := unsavedpostvideoFields[3].Descriptor()
	// unsavedpostvideo.DefaultCreatedAt holds the default value on creation for the created_at field.
	unsavedpostvideo.DefaultCreatedAt = unsavedpostvideoDescCreatedAt.Default.(func() time.Time)
}
