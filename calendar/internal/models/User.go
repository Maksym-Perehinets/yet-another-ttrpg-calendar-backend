package models

type Users struct {
	ID    uint    `json:"id" gorm:"primaryKey"`
	Name  string  `json:"name" gorm:"not null"`
	Games []Games `json:"games" gorm:"many2many:game_users"`
}
