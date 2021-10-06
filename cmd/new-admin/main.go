package main

import (
	"bufio"
	"context"
	"database/sql"
	"devlog/ent"
	"devlog/env"
	"devlog/util"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

func main() {
	env.InitDBEnvVars()

	db, err := sql.Open(env.DatabaseDriver, env.DatabaseDSN)

	if err != nil {
		panic(err)
	}

	var driver *entsql.Driver

	if env.DatabaseDriver == "pgx" {
		driver = entsql.OpenDB(dialect.Postgres, db)
	} else {
		driver = entsql.OpenDB(env.DatabaseDriver, db)
	}

	client := ent.NewClient(ent.Driver(driver))

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Email: ")
	if !scanner.Scan() {
		panic("Email required")
	}
	email := scanner.Text()

	fmt.Print("Username: ")
	if !scanner.Scan() {
		panic("Username required")
	}
	username := scanner.Text()

	fmt.Print("Password: ")
	if !scanner.Scan() {
		panic("Password required")
	}
	password := scanner.Text()

	fmt.Print("Password confirm: ")
	if !scanner.Scan() {
		panic("Password confirm required")
	}
	passwordConfirm := scanner.Text()

	if password != passwordConfirm {
		panic("Password not match")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		panic(err)
	}

	key, err := util.GenerateToken64()
	if err != nil {
		panic(err)
	}

	admin := client.Admin.Create().
		SetEmail(strings.TrimSpace(email)).
		SetUsername(strings.TrimSpace(username)).
		SetPassword(string(passwordHash)).
		SetKey(key).
		SaveX(context.Background())
	fmt.Println("A new admin has been successfully created: ", admin.ID)
}
