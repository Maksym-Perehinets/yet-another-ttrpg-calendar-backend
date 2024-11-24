package models

type Locations struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	City        string  `json:"city" gorm:"not null"`
	Street      string  `json:"street" gorm:"not null"`
	LinkToSite  string  `json:"link_to_site" gorm:"not null"`
	Price       uint    `json:"price" gorm:"not null"`
	PricingType string  `json:"pricing_type" gorm:"not null"`
	Games       []Games `json:"games" gorm:"foreignKey:LocationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
