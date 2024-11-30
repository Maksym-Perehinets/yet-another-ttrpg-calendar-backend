package interfaces

import (
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/internal/models"
)

type TTRPGSystems interface {
	CreateTTRPGSystem(system *models.TTRPGSystems) (uint, error)
}
