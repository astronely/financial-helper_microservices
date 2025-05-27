package model

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	jwt.RegisteredClaims
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type BoardClaims struct {
	jwt.RegisteredClaims
	ID      int64 `json:"id"`
	OwnerID int64 `json:"ownerId"`
}
