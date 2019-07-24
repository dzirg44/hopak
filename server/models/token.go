package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
