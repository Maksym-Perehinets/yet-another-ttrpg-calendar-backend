package models

type TTRPGSystems struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"not null; unique"`
	Genre       string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	LinkToSite  string  `gorm:"not null"`
	Games       []Games `gorm:"foreignKey:TTRPGSystemID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (t *TTRPGSystems) GetID() uint {
	return t.ID
}
