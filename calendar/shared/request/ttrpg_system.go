package request

import "github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/internal/models"

type CreateTTRPGSystem struct {
	Name        string `json:"name" binding:"required"`
	Genre       string `json:"genre" binding:"required"`
	Description string `json:"description" binding:"required"`
	LinkToSite  string `json:"link_to_site" binding:"required"`
}

func (c *CreateTTRPGSystem) ToModel() *models.TTRPGSystems {
	return &models.TTRPGSystems{
		Name:        c.Name,
		Genre:       c.Genre,
		Description: c.Description,
		LinkToSite:  c.LinkToSite,
	}
}
