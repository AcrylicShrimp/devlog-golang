package model

import (
	"devlog/ent"
	"time"
)

type NewCategoryParam struct {
	Name        string  `json:"name" validate:"required,max=32" example:"web"`
	Description *string `json:"description" validate:"min=1,max=255" example:"web-related things"`
}

type DeleteCategoryParam struct {
	Name string `param:"name" validate:"required,max=32" example:"web"`
}

type Category struct {
	Name        string `json:"name" validate:"required" example:"web"`
	Description string `json:"description" example:"web-related things"`
	CreatedAt   string `json:"created-at" validate:"required" example:"2021-08-18T00:00:00Z00:00"`
	ModifiedAt  string `json:"modified-at" validate:"required" example:"2021-08-18T00:00:00Z00:00"`
}

func CategoryFromModel(category *ent.Category) Category {
	return Category{
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt.UTC().Format(time.RFC3339),
		ModifiedAt:  category.ModifiedAt.UTC().Format(time.RFC3339),
	}
}
