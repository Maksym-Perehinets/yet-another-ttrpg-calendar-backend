package database

import (
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/internal/models"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/shared/crud"
)

func (s *service) CreateTTRPGSystem(system *models.TTRPGSystems) (uint, error) {
	id, _ := crud.Create(s.db, system)

	return id, nil
}
