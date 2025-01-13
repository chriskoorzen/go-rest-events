package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "reallyhardtoguesssupersecretkey"

func GenerateJWToken(userID int64, userEmail string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  userEmail,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	token, err := claims.SignedString([]byte(secretKey))
	return token, err
}

func VerifyJWToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token") // technically invalid signature, but we don't want to expose that info
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	// _, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token")
	}

	// extract information from token
	// email := claims["email"].(string)
	userID := int64(claims["userID"].(float64))

	return userID, nil
}
