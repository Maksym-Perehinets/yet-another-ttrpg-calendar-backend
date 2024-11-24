package models

type Games struct {
	ID            uint         `json:"id" gorm:"primaryKey"`
	Name          string       `json:"name" gorm:"not null"`
	Description   string       `json:"description"`
	PictureLink   string       `json:"picture_link"`
	GameMasterID  uint         `json:"game_master_id" gorm:"not null"`
	Players       []Users      `json:"players" gorm:"many2many:game_users"`
	TTRPGSystemID uint         `json:"ttrpg_system_id" gorm:"not null"`
	TTRPGSystem   TTRPGSystems `json:"ttrpg_system" gorm:"foreignKey:TTRPGSystemID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	LocationID    uint         `json:"location_id" gorm:"not null"`
	Location      Locations    `json:"location" gorm:"foreignKey:LocationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
