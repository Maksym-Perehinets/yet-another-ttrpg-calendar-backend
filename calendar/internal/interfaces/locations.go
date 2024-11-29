package interfaces

import (
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/internal/models"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/shared/request"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/common/database/paginate"
)

type Locations interface {
	// CreateLocation creates location if no such exists
	CreateLocation(locations *models.Locations) (uint, error)

	// GetLocation returns array of all available locations
	GetLocation(id int) (*models.Locations, error)

	// GetLocations returns array of locations
	GetLocations(page int, amount int) (*paginate.Pagination, error)

	// UpdateLocation updates location by its id
	UpdateLocation(id int, location request.Update) error

	// DeleteLocation return location by its id
	DeleteLocation(id int) error
}
