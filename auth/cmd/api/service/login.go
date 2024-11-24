package service

import (
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/cmd/api/service/validate"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/interfaces"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// LoginService is a service for logging in a user
// returns a JWT token, expiration time and an error if any
func LoginService(s interfaces.Service, c *gin.Context) (string, time.Time, error) {
	l := validate.NewLoginRequest(c)

	token, cookieExpiresIn, err := l.GetUserJWT(s)
	if err != nil {
		log.Printf("Error logging in: %s", err.Error())
		return "", time.Time{}, err
	}

	return token, cookieExpiresIn, nil
}
