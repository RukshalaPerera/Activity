package utils

import (
	"Project2/app/models"
	"github.com/dgrijalva/jwt-go"
	"os"
)

func ParseToken(tokenString string) (models.Claims, error) {
	var claims models.Claims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		return models.Claims{}, err
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return *claims, nil
	} else {
		return models.Claims{}, jwt.NewValidationError("invalid token", jwt.ValidationErrorClaimsInvalid)
	}
}
