package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username       string `gorm:"unique;not null"`
	Email          string `gorm:"unique;not null"`
	Password       string `gorm:"not null"`
	Role           string `gorm:"default:user"`
	ProfilePicture string `gorm:"default:TODO.jpg"`
	SSO            bool   `gorm:"default:false"`
	TelegramLink   string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
