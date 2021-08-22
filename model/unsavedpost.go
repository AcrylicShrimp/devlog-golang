package model

import (
	"devlog/ent"
	"time"
)

type GetUnsavedPostParam struct {
	UUID string `param:"uuid" validate:"required,hexadecimal,len=64" example:"fd00000aa8660b5b010006acdc0100000101000100010000fd00000aa8660b5b"`
}

type UnsavedPostThumbnail struct {
	Validity  string  `json:"validity" validate:"required" example:"pending"`
	Width     *uint32 `json:"width" example:"256"`
	Height    *uint32 `json:"height" example:"128"`
	Hash      *string `json:"hash" example:"LEHV6nWB2yk8pyo0adR*.7kCMdnj"`
	URL       *string `json:"url" example:"https://image.example.com/example-image"`
	CreatedAt string  `json:"created-at" validate:"required" example:"2021-08-18 00:00:00Z"`
}

func UnsavedPostThumbnailFromModel(thumbnail *ent.UnsavedPostThumbnail) UnsavedPostThumbnail {
	return UnsavedPostThumbnail{
		Validity:  (string)(thumbnail.Validity),
		Width:     thumbnail.Width,
		Height:    thumbnail.Height,
		Hash:      thumbnail.Hash,
		URL:       thumbnail.URL,
		CreatedAt: thumbnail.CreatedAt.UTC().Format(time.RFC3339),
	}
}

type UnsavedPostImage struct {
	UUID      string  `json:"uuid" validate:"required" example:"fd00000aa8660b5b010006acdc0100000101000100010000fd00000aa8660b5b"`
	Validity  string  `json:"validity" validate:"required" example:"pending"`
	Width     *uint32 `json:"width" example:"256"`
	Height    *uint32 `json:"height" example:"128"`
	Hash      *string `json:"hash" example:"LEHV6nWB2yk8pyo0adR*.7kCMdnj"`
	Title     *string `json:"title" example:"My image"`
	URL       *string `json:"url" example:"https://image.example.com/example-image"`
	CreatedAt string  `json:"created-at" validate:"required" example:"2021-08-18 00:00:00Z"`
}

func UnsavedPostImageFromModel(image *ent.UnsavedPostImage) UnsavedPostImage {
	return UnsavedPostImage{
		UUID:      image.UUID,
		Validity:  (string)(image.Validity),
		Width:     image.Width,
		Height:    image.Height,
		Hash:      image.Hash,
		Title:     image.Title,
		URL:       image.URL,
		CreatedAt: image.CreatedAt.UTC().Format(time.RFC3339),
	}
}

type UnsavedPostWithoutImage struct {
	UUID        string                `json:"uuid" validate:"required" example:"fd00000aa8660b5b010006acdc0100000101000100010000fd00000aa8660b5b"`
	Slug        *string               `json:"slug" example:"my-first-post"`
	AccessLevel *string               `json:"access-level" example:"public"`
	Title       *string               `json:"title" example:"My first post"`
	CreatedAt   string                `json:"created-at" validate:"required" example:"2021-08-18 00:00:00Z"`
	ModifiedAt  string                `json:"modified-at" validate:"required" example:"2021-08-18 00:00:00Z"`
	Category    *string               `json:"category" example:"web"`
	Thumbnail   *UnsavedPostThumbnail `json:"thumbnail"`
}

func UnsavedPostWithoutImageFromModel(post *ent.UnsavedPost) UnsavedPostWithoutImage {
	var category *string

	if post.Edges.Category != nil {
		category = &post.Edges.Category.Name
	}

	var thumbnail *UnsavedPostThumbnail

	if post.Edges.Thumbnail != nil {
		model := UnsavedPostThumbnailFromModel(post.Edges.Thumbnail)
		thumbnail = &model
	}

	return UnsavedPostWithoutImage{
		UUID:        post.UUID,
		Slug:        post.Slug,
		AccessLevel: (*string)(post.AccessLevel),
		Title:       post.Title,
		CreatedAt:   post.CreatedAt.UTC().Format(time.RFC3339),
		ModifiedAt:  post.ModifiedAt.UTC().Format(time.RFC3339),
		Category:    category,
		Thumbnail:   thumbnail,
	}
}

type UnsavedPost struct {
	UUID        string                `json:"uuid" validate:"required" example:"fd00000aa8660b5b010006acdc0100000101000100010000fd00000aa8660b5b"`
	Slug        *string               `json:"slug" example:"my-first-post"`
	AccessLevel *string               `json:"access-level" example:"public"`
	Title       *string               `json:"title" example:"My first post"`
	CreatedAt   string                `json:"created-at" validate:"required" example:"2021-08-18 00:00:00Z"`
	ModifiedAt  string                `json:"modified-at" validate:"required" example:"2021-08-18 00:00:00Z"`
	Category    *string               `json:"category" example:"web"`
	Thumbnail   *UnsavedPostThumbnail `json:"thumbnail"`
	Images      []UnsavedPostImage    `json:"images"`
}

func UnsavedPostFromModel(post *ent.UnsavedPost) UnsavedPost {
	var category *string

	if post.Edges.Category != nil {
		category = &post.Edges.Category.Name
	}

	var thumbnail *UnsavedPostThumbnail

	if post.Edges.Thumbnail != nil {
		model := UnsavedPostThumbnailFromModel(post.Edges.Thumbnail)
		thumbnail = &model
	}

	images := make([]UnsavedPostImage, len(post.Edges.Images))

	for index, image := range post.Edges.Images {
		images[index] = UnsavedPostImageFromModel(image)
	}

	return UnsavedPost{
		UUID:        post.UUID,
		Slug:        post.Slug,
		AccessLevel: (*string)(post.AccessLevel),
		Title:       post.Title,
		CreatedAt:   post.CreatedAt.UTC().Format(time.RFC3339),
		ModifiedAt:  post.ModifiedAt.UTC().Format(time.RFC3339),
		Category:    category,
		Thumbnail:   thumbnail,
		Images:      images,
	}
}
