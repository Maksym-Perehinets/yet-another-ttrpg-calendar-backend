package interfaces

import (
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/internal/models"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/common/database/paginate"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/common/requests"
)

type TTRPGSystems interface {
	// CreateTTRPGSystem creates a TTRPG system in the database.
	CreateTTRPGSystem(system *models.TTRPGSystems) (uint, error)

	// UpdateTTRPGSystem updates a TTRPG system in the database.
	UpdateTTRPGSystem(id int, system requests.Update) error

	// GetTTRPGSystem returns a TTRPG system from the database.
	GetTTRPGSystem(id int) (*models.TTRPGSystems, error)

	// GetTTRPGSystems returns a list of TTRPG systems from the database.
	GetTTRPGSystems(page int, amount int) (*paginate.Pagination, error)

	// DeleteTTRPGSystem deletes a TTRPG system from the database.
	DeleteTTRPGSystem(id int) error
}
