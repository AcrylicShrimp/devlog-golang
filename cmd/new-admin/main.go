package main

import (
	"bufio"
	"context"
	"devlog/ent"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func main() {
	client, err := ent.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_DSN"))
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Print("Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Print("Password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Print("Password confirm: ")
	passwordConfirm, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	if password != passwordConfirm {
		panic("Password not match")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		panic(err)
	}

	admin := client.Admin.Create().SetEmail(email).SetUsername(username).SetPassword(hex.EncodeToString(passwordHash)).SaveX(context.Background())
	fmt.Print("A new admin has been successfully created: ", admin.ID)
}
