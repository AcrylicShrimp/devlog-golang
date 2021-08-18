package model

type AdminCategory struct {
	Name string `json:"name" example:"web"`
	Description string `json:"description" example:"web-related things"`
	CreatedAt string `json:"created-at" example:"2021-08-18 00:00:00Z"`
	ModifiedAt string `json:"modified-at" example:"2021-08-18 00:00:00Z"`
}
