package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "reallyhardtoguesssupersecretkey"

func GenerateToken(userID int64, userEmail string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  userEmail,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	token, err := claims.SignedString([]byte(secretKey))
	return token, err
}
