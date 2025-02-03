package model

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	jwt.StandardClaims
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
