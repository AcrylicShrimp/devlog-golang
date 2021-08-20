package main

import (
	"context"
	"devlog/ent"
	"devlog/env"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4"
)

func InitDB() (*ent.Client, error) {
	client, err := ent.Open(env.DatabaseDriver, env.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	if err = client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}
