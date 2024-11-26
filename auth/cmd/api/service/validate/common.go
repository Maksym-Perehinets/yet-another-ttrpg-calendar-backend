package validate

import (
	"fmt"
	"regexp"
)

var (
	emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

func validateEmail(email string) error {
	matched, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("invalid email format")
	}
	return nil
}
