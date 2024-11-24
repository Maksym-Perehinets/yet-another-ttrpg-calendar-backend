package service

import (
	"fmt"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/interfaces"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/internal/models"
	"log"
)

func ChangeRoleService(service interfaces.Service, user string, role string) (string, error) {
	log.Printf("Changing permission for user: %s", user)

	if role != "admin" && role != "user" {
		return "", fmt.Errorf("role %s is invalid", role)
	}

	rs := service.DB().Model(&models.User{}).Where("id = ? AND role != ?", user, role).Update("role", role)
	if rs.Error != nil {
		return "", rs.Error
	}

	if rs.RowsAffected == 0 {
		return "", fmt.Errorf("user with id %s already has an %s role", user, role)
	}

	return user, nil
}
