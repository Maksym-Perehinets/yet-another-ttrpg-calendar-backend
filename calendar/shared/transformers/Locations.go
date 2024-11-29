package transformers

import (
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/internal/models"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/shared/response"
)

func ToLocationResponse(location *models.Locations) response.LocationResponse {
	return response.LocationResponse{
		ID:          location.ID,
		Name:        location.Name,
		Description: location.Description,
		City:        location.City,
		Street:      location.Street,
		LinkToSite:  location.LinkToSite,
		Price:       location.Price,
		PricingType: location.PricingType,
		OpenAt:      location.OpenAt,
		CloseAt:     location.CloseAt,
		Games:       ToGameResponses(location.Games),
	}
}

func ToLocationsResponse(locations []*models.Locations) []response.LocationResponse {
	responses := make([]response.LocationResponse, len(locations))
	for i, location := range locations {
		responses[i] = ToLocationResponse(location)
	}
	return responses
}

func ToGameResponses(games []models.Games) []response.GameResponse {
	responses := make([]response.GameResponse, len(games))
	for i, game := range games {
		responses[i] = response.GameResponse{
			ID: game.ID,
			//Name: game.Name,
		}
	}
	return responses
}
