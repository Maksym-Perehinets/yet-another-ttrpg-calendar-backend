package service

import (
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/interfaces"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/internal/models"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/auth/shared/request"
	"log"
)

// UpdateUserService updates user info
func UpdateUserService(db interfaces.Service, userID string, updateFields request.UserUpdate) error {
	log.Printf("Updating user: %s with fields %s", userID, updateFields.ToMap())
	// Update user info
	rs := db.DB().Model(models.User{}).
		Omit("id", "password", "role", "sso").
		Where("id = ?", userID).Updates(updateFields.ToMap())
	if rs.Error != nil {
		return rs.Error
	}

	return nil
}
