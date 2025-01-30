package model

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	jwt.StandardClaims
	Email    string `json:"email"`
	Username string `json:"username"`
}
