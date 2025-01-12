package models

import (
	"github.com/chriskoorzen/go-rest-demo/db"
	"github.com/chriskoorzen/go-rest-demo/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"    binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) Save() error {
	// Save user to database
	query := `
	INSERT INTO users (email, password)
	VALUES (?, ?)`

	stmnt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmnt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmnt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	user.ID = id
	return err
}
