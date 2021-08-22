package model

type AdminCategory struct {
	Name        string `json:"name" example:"web" validate:"required"`
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
