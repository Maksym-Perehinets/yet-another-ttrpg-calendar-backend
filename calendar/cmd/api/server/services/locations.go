package services

import (
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/internal/interfaces"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/internal/models"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/shared/transformers"
)

func GetLocationsService(db interfaces.Service, page, limit int) (interface{}, error) {
	data, err := db.GetLocations(page, limit)
	if err != nil {
		return nil, err
	}

	data.Entries = transformers.ToLocationsResponse(data.Entries.([]*models.Locations))

	return data, nil
}
