package interfaces

import "auth/shared/request"

type UserInput interface {
	// ValidateEmail validates an email address returns nil if valid if not error
	ValidateEmail(Service) error

	// ValidateUsername validates a username returns nil if valid if not error
	ValidateUsername(Service) error

	// ValidatePassword validates a password returns nil if valid if not error
	ValidatePassword() error

	// GetStruct returns the struct
	GetStruct() *request.RegisterRequest
}
