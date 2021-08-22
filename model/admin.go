package model

type AdminCategory struct {
	Name        string `json:"name" validate:"required" example:"web"`
	Description string `json:"description" example:"web-related things"`
	CreatedAt   string `json:"created-at" validate:"required" example:"2021-08-18 00:00:00Z"`
	ModifiedAt  string `json:"modified-at" validate:"required" example:"2021-08-18 00:00:00Z"`
}

type AdminNewCategory struct {
	Name        string  `json:"name" validate:"required,max=32" example:"web"`
	Description *string `json:"description" validate:"min=1,max=255" example:"web-related things"`
}

type AdminDeleteCategory struct {
	Name string `json:"name" validate:"required,max=32" example:"web"`
}

type AdminNewCategoryCreated struct {
	Name string `json:"name" validate:"required" example:"web"`
}

type AdminUnsavedPostThumbnail struct {
	Validity  string  `json:"validity" validate:"required" example:"pending"`
	Width     *uint32 `json:"width" example:"256"`
	Height    *uint32 `json:"height" example:"128"`
	Hash      *string `json:"hash" example:"LEHV6nWB2yk8pyo0adR*.7kCMdnj"`
	URL       *string `json:"url" example:"https://image.example.com/example-image"`
	CreatedAt string  `json:"created-at" validate:"required" example:"2021-08-18 00:00:00Z"`
}

type AdminUnsavedPost struct {
	UUID        string                     `json:"uuid" validate:"required" example:"fd00000aa8660b5b010006acdc0100000101000100010000fd00000aa8660b5b"`
	Slug        *string                    `json:"slug" example:"my-first-post"`
	AccessLevel *string                    `json:"access-level" example:"public"`
	Title       *string                    `json:"title" example:"My first post"`
	CreatedAt   string                     `json:"created-at" validate:"required" example:"2021-08-18 00:00:00Z"`
	ModifiedAt  string                     `json:"modified-at" validate:"required" example:"2021-08-18 00:00:00Z"`
	Category    *string                    `json:"category" example:"web"`
	Thumbnail   *AdminUnsavedPostThumbnail `json:"thumbnail"`
}
