package validate

import (
	"fmt"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/interfaces"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/internal/models"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/shared/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
)

var (
	emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

type RegisterRequest struct {
	request.RegisterRequest
}

func (r *RegisterRequest) ValidateEmail(service interfaces.Service) error {
	log.Printf("Validating email: %s", r.Email)
	matched, err := regexp.MatchString(emailRegex, r.Email)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("invalid email format")
	}

	if err := service.AlreadyExists("email", r.Email, models.User{}); err != nil {
		return err
	}

	return nil
}

func (r *RegisterRequest) ValidateUsername(service interfaces.Service) error {
	log.Printf("Validating username: %s", r.Username)
	if len(r.Username) < 3 {
		return fmt.Errorf("username must be at least 3 characters long")
	}

	if err := service.AlreadyExists("username", r.Username, models.User{}); err != nil {
		return err
	}

	return nil
}

func (r *RegisterRequest) ValidatePassword() error {
	log.Printf("Validating password for user: %s", r.Email)

	// Check if password meets length requirements
	if len(r.Password) < 8 || len(r.Password) > 64 {
		return fmt.Errorf("password must be between 8 and 64 characters")
	}

	// Define regex patterns and corresponding error messages
	rules := []struct {
		pattern string
		message string
	}{
		{`[a-z]`, "password must contain at least one lowercase letter"},
		{`[A-Z]`, "password must contain at least one uppercase letter"},
		{`\d`, "password must contain at least one digit"},
		{`[@$!%*?&]`, "password must contain at least one special character (@$!%*?&)"},
	}

	// Validate against each rule
	for _, rule := range rules {
		if !regexp.MustCompile(rule.pattern).MatchString(r.Password) {
			return fmt.Errorf(rule.message)
		}
	}

	return nil
}

func (r *RegisterRequest) GetStruct() *request.RegisterRequest {
	return &r.RegisterRequest
}

func NewRegisterRequest(c *gin.Context) interfaces.UserInput {
	var user RegisterRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error binding request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	return &user
}
