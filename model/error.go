package model

type HTTPError struct {
	Message string `json:"message" example:"500 - Internal Server Error"`
}
