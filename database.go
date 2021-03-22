package main

import (
	"context"
	"devlog/ent"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func InitDB() (*ent.Client, error) {
	client, err := ent.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_DSN"))
	if err != nil {
		return nil, err
	}

	if err = client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}
