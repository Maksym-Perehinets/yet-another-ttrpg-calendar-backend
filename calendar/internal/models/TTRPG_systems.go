package models

type TTRPGSystems struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Name       string  `json:"name" gorm:"not null"`
	Genre      string  `json:"genre" gorm:"not null"`
	LinkToSite string  `json:"link_to_site" gorm:"not null"`
	Games      []Games `json:"games" gorm:"foreignKey:TTRPGSystemID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
