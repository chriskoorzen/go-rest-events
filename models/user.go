package models

import (
	"errors"

	"github.com/chriskoorzen/go-rest-events/db"
	"github.com/chriskoorzen/go-rest-events/utils"
)

// "User" is the struct that must be used to interact with the database internally
// "UserExposedJSON" is the struct that must be returned to the client publicly
//
// It is necessary to maintain a separation between the two structs to prevent sensitive information from being exposed
// because there exists a conflict between the struct tags "binding:"required" and "json".
// If the json tag is turned off with "-", the binding tag will not work, which is a problem when validating the incoming request body.
// If not turned off, sensitive fields and information will be exposed to the client.
type UserExposedJSON struct {
	Email string `json:"email"`
}

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
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

	// SECURITY: Keep in mind that the time taken to check for a valid username (email) is
	// quick, whereas a password hash check is slow. This discrepancy can be used to determine
	// if a user exists in the database without revealing if the password is correct.

	var id int64
	var hashedPassword string

	err := row.Scan(&id, &hashedPassword)
	if err != nil { // username not found
		return errors.New("invalid credentials")
	}

	if !utils.IsPasswordHash(user.Password, hashedPassword) { // password does not match
		return errors.New("invalid credentials")
	}

	// If we reach this point, the user is valid
	user.ID = id
	return nil
}
