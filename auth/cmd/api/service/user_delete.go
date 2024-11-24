package service

import (
	"fmt"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/interfaces"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/internal/models"
	"log"
	"strconv"
)

func DeleteUserService(service interfaces.Service, param string) (string, error) {
	log.Printf("Deleting user: %s", param)
	id, err := strconv.Atoi(param)
	if err != nil {
		return "", err
	}

	rs := service.DB().Where("id = ?", id).Delete(&models.User{})
	if rs.Error != nil {
		return "", rs.Error
	}

	if rs.RowsAffected == 0 {
		return "", fmt.Errorf("user with id %s not found", param)
	}

	return param, nil

}
