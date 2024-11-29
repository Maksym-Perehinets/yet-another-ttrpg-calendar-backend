package interfaces

import "gorm.io/gorm"

type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error

	// DB exposes the GORM DB instance for application use.
	DB() *gorm.DB

	// Locations all interactions with locations table form db encapsulated here
	Locations
}
