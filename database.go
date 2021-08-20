package main

import (
	"context"
	"database/sql"
	"devlog/ent"
	"devlog/env"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func InitDB() (*ent.Client, error) {
	db, err := sql.Open(env.DatabaseDriver, env.DatabaseDSN)

	if err != nil {
		return nil, err
	}

	var driver *entsql.Driver

	if env.DatabaseDriver == "pgx" {
		driver = entsql.OpenDB(dialect.Postgres, db)
	} else {
		driver = entsql.OpenDB(env.DatabaseDriver, db)
	}

	client := ent.NewClient(ent.Driver(driver))

	if err = client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}
