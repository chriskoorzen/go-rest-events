package models

import (
	"errors"

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

func (user *User) ValidateCredentials() error {
	// Check if user exists in database
	query := `
	SELECT id, password
	FROM users
	WHERE email = ?`

	row := db.DB.QueryRow(query, user.Email)

	var id int64
	var hashedPassword string

	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, hashedPassword)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	user.ID = id
	return nil
}
