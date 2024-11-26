package service

import (
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/internal/database"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/internal/models"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/shared/request"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/utils"
	"log"
	"strconv"
	"time"

	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/common/auth/jwt"
)

// RegisterService registers a user, returns jwt token if successful, error if not
func RegisterService(user *request.RegisterRequest) (string, time.Time, error) {
	log.Printf("Registering user: %s", user.Username)
	db := database.New()
	var createdUser models.User

	preparedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		log.Fatalf(utils.ToString(err))
	}

	result := db.DB().Create(&models.User{
		Username:       user.Username,
		Email:          user.Email,
		Password:       string(preparedPassword),
		ProfilePicture: user.ProfilePicture,
		TelegramLink:   user.TelegramLink,
	}).First(&createdUser)
	if result.Error != nil {
		log.Printf("Error inserting user: %v", result.Error)
		return "", time.Time{}, result.Error
	}

	token, cookieExpiresIn, err := jwt.GenerateToken(
		strconv.FormatUint(uint64(createdUser.ID), 10),
		user.Username,
		"user",
	)
	if err != nil {
		log.Fatalf("error generating token: %v", err)
		return "", time.Time{}, err
	}

	return token, cookieExpiresIn, nil
}
