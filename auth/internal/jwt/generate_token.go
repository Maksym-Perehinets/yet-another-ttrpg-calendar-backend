package jwt

import (
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	jwtSecret    = []byte(os.Getenv("JWT_SECRET"))
	jwtExpiresIn = os.Getenv("JWT_EXPIRES_IN")
)

// GenerateToken generates a JWT token for the user
// returns the token string, expiration time for token and an error if any
func GenerateToken(userId string, username string, role string) (string, time.Time, error) {
	log.Printf("Generating token for user: %s", username)
	GetTokenExpiration, err := strconv.Atoi(jwtExpiresIn)
	if err != nil {
		log.Fatalf("error converting JWT_EXPIRES_IN to int: %v", err)
		return "", time.Time{}, err
	}

	tokenExpiresIn := time.Now().Add(time.Duration(GetTokenExpiration) * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, models.Claims{
		UserID:   userId,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "auth.service",
			ExpiresAt: jwt.NewNumericDate(tokenExpiresIn),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Fatalf("error while signing token: %v", err)
		return "", time.Time{}, err
	}

	return tokenString, tokenExpiresIn, nil
}
