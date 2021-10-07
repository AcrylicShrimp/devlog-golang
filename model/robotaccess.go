package model

import (
	"devlog/ent"
	"time"
)

type RobotAccess struct {
	Memo         *string `json:"memo" validate:"max=1024" example:"for AWS Lambda"`
	CreatedAt    string  `json:"created-at" validate:"required" example:"2021-08-18T00:00:00Z00:00"`
	ExpiresAt    *string `json:"modified-at" example:"2021-08-18T00:00:00Z00:00"`
	LastAccessAt *string `json:"last-access-at" example:"2021-08-18T00:00:00Z00:00"`
}

func RobotAccessFromModel(access *ent.RobotAccess) RobotAccess {
	var expiresAt *string
	if access.ExpiresAt != nil {
		var time = access.ExpiresAt.UTC().Format(time.RFC3339)
		expiresAt = &time
	}

	var lastAccessAt *string
	if access.LastAccessAt != nil {
		var time = access.LastAccessAt.UTC().Format(time.RFC3339)
		lastAccessAt = &time
	}

	return RobotAccess{
		Memo:         access.Memo,
		CreatedAt:    access.CreatedAt.UTC().Format(time.RFC3339),
		ExpiresAt:    expiresAt,
		LastAccessAt: lastAccessAt,
	}
}
