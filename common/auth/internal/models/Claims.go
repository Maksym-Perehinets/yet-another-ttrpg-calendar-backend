package models

import "github.com/golang-jwt/jwt/v5"

// Claims is the JWT claims structure
type Claims struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
