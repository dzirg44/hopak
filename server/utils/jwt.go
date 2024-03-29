package utils

import (
	"fmt"

	"github.com/dzirg44/hopak/server/models"
	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT generates new JWT token for the user
func GenerateJWT(user models.User) (string, error) {

	tokenObj := &models.Token{
		UserID:   user.UserID,
		Username: user.Username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenObj)

	tokenString, err := token.SignedString([]byte(MySecretKey))

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil

}
