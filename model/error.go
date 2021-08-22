package model

type HTTPError400 struct {
	Message string `json:"message" validate:"required" example:"400 - Bad request"`
}

type HTTPError401 struct {
	Message string `json:"message" validate:"required" example:"401 - Unauthorized"`
}

type HTTPError404 struct {
	Message string `json:"message" validate:"required" example:"404 - Not found"`
}

type HTTPError409 struct {
	Message string `json:"message" validate:"required" example:"409 - Conflict"`
}

type HTTPError500 struct {
	Message string `json:"message" validate:"required" example:"500 - Internal Server Error"`
}
