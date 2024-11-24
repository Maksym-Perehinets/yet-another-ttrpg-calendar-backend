package interfaces

import (
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/internal/models"
	"time"
)

type UserLoginInput interface {
	// EmailOrUsername returns the email or username depending on the input,
	// and returns an error if both are empty or value and column name to get full user info
	EmailOrUsername() (string, string, error)

	// ValidatePassword validates the password
	ValidatePassword(Service) (*models.User, error)

	// GetUserJWT returns the user JWT token
	GetUserJWT(Service) (string, time.Time, error)
}
