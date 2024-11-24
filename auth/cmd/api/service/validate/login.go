package validate

import (
	"auth/interfaces"
	"auth/internal/jwt"
	"auth/internal/models"
	"auth/shared/request"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type LoginRequest struct {
	request.LoginRequest
}

func (l *LoginRequest) EmailOrUsername() (string, string, error) {
	log.Printf("Validating login request")
	if l.Username == "" && l.Email == "" {
		return "", "", errors.New("username or email is required")
	}

	if l.Username != "" {
		return l.Username, "username", nil
	}

	return l.Email, "email", nil
}

func (l *LoginRequest) ValidatePassword(service interfaces.Service) (*models.User, error) {
	if l.Password == "" {
		return nil, errors.New("password is required")
	}

	v, f, err := l.EmailOrUsername()

	if err != nil {
		return nil, err
	}
	log.Printf("Validating password for user: %s", v)

	user, err := service.GetUser(f, v)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(l.Password))
	if err != nil {
		log.Printf("Password mismatch %s", v)
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (l *LoginRequest) GetUserJWT(service interfaces.Service) (string, time.Time, error) {

	user, err := l.ValidatePassword(service)
	if err != nil {
		return "", time.Time{}, err
	}

	jwtToken, expireIn, err := jwt.GenerateToken(strconv.FormatUint(uint64(user.ID), 10), user.Username, user.Role)
	if err != nil {
		return "", time.Time{}, err
	}

	return jwtToken, expireIn, nil
}

// NewLoginRequest creates a new login request
func NewLoginRequest(c *gin.Context) interfaces.UserLoginInput {
	var l LoginRequest
	if err := c.ShouldBindJSON(&l); err != nil {
		log.Printf("Error binding request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}

	return &l
}
