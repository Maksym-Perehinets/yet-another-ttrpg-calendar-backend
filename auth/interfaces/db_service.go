package interfaces

import (
	"auth/internal/models"
	"gorm.io/gorm"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// AlreadyExists checks if a record exists in the database.
	// It returns an error if the record already exists.
	// The parameters are the colum name and value and table to check.
	AlreadyExists(c string, v string, m any) error

	// GetUser retrieves a user from the database based on the column and value provided.
	GetUser(c string, v string) (*models.User, error)

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error

	// DB exposes the GORM DB instance for application use.
	DB() *gorm.DB
}
