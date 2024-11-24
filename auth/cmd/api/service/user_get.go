package service

import (
	"auth/interfaces"
	"auth/internal/models"
	"auth/shared/response"
	"log"
)

func GetUsersService(service interfaces.Service) (interface{}, error) {
	log.Printf("Getting all users")
	var users []response.APIUsers
	err := service.DB().Model(&models.User{}).Find(&users)
	if err.Error != nil {
		return nil, err.Error
	}
	log.Printf("Users: %v", users)
	return users, nil
}
