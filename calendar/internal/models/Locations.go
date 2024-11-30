package models

type Locations struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	City        string  `gorm:"not null"`
	Street      string  `gorm:"not null; unique"`
	LinkToSite  string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	PricingType string  `gorm:"not null"`
	OpenAt      string  `gorm:"not null"`
	CloseAt     string  `gorm:"not null"`
	Games       []Games `gorm:"foreignKey:LocationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (l *Locations) GetID() uint {
	return l.ID
}
