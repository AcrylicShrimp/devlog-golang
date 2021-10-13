package model

import (
	"devlog/ent"
	"time"
)

type RobotAccessParam struct {
	Memo *string `json:"slug" validate:"max=1024" example:"for AWS Lambda"`
	TTL  *int64  `json:"ttl" validate:"min=0" example:"86400"` // Time-to-live in seconds
}

type RobotAccessTokenParam struct {
	Token string `json:"-" param:"token" validate:"required,hexadecimal,len=256" example:"175d2572edf5e47826021b8e4ed5ea0497bd58e45d2c69b79517ee53b10ebcd35ed035b96eee4fffa595e2777da45d27c5f4692891572adb29c3a904de38f457a1c89a4b9cf97c4365368207916388ab78f598f4c9cde71df54c613cc4eca23e48acc6e719291c0c3311354e0c009465382fb457244a64814891eda82d211626"`
}

type RobotAccess struct {
	IssuedFor    string  `json:"issued-for" validate:"required" example:"issuer"`
	Memo         *string `json:"memo" example:"for AWS Lambda"`
	CreatedAt    string  `json:"created-at" validate:"required" example:"2021-08-18T00:00:00Z00:00"`
	ExpiresAt    *string `json:"modified-at" example:"2021-08-18T00:00:00Z00:00"`
	LastAccessAt *string `json:"last-access-at" example:"2021-08-18T00:00:00Z00:00"`
}

func RobotAccessFromModel(access *ent.RobotAccess) RobotAccess {
	var expiresAt *string

	if access.ExpiresAt != nil {
		var expires = access.ExpiresAt.UTC().Format(time.RFC3339)
		expiresAt = &expires
	}

	var lastAccessAt *string

	if access.LastAccessAt != nil {
		var lastAccess = access.LastAccessAt.UTC().Format(time.RFC3339)
		lastAccessAt = &lastAccess
	}

	return RobotAccess{
		IssuedFor:    access.Edges.User.Username,
		Memo:         access.Memo,
		CreatedAt:    access.CreatedAt.UTC().Format(time.RFC3339),
		ExpiresAt:    expiresAt,
		LastAccessAt: lastAccessAt,
	}
}

type RobotAccessTokenOnly struct {
	Token string `json:"token" validate:"required" example:"175d2572edf5e47826021b8e4ed5ea0497bd58e45d2c69b79517ee53b10ebcd35ed035b96eee4fffa595e2777da45d27c5f4692891572adb29c3a904de38f457a1c89a4b9cf97c4365368207916388ab78f598f4c9cde71df54c613cc4eca23e48acc6e719291c0c3311354e0c009465382fb457244a64814891eda82d211626"`
}

func RobotAccessTokenOnlyFromToken(token string) RobotAccessTokenOnly {
	return RobotAccessTokenOnly{
		Token: token,
	}
}
