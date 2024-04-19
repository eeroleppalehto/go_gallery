package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/eeroleppalehto/go_gallery/models"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("mysql", "root:Q2werty@/gollery?parseTime=true")
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
		return
	}

	queries := models.New(db)

	password := "Q2werty"
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		fmt.Println("Failed to hash to hash password")
		return
	}

	hash := string(bytes)

	result, err := queries.CreateUser(ctx, models.CreateUserParams{
		Username: "eevokki",
		Email:    "moi@gmail.com",
		Password: hash,
	})

	if err != nil {
		fmt.Println("Failed to create user")
		return
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Added user with id: ", insertedId)

	fetchUser, err := queries.GetUsers(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range fetchUser {
		fmt.Println(user.UserID, user.Username, user.Password, user.Email, user.CretedAt)
	}
}
