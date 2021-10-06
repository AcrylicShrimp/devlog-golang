package model

import "github.com/golang-jwt/jwt"

type SessionParam struct {
	Username string `json:"username" validate:"required" example:"username"`
	Password string `json:"password" validate:"required" example:"password"`
}

type SessionClaim struct {
	Key string `json:"key"`
	jwt.StandardClaims
}
