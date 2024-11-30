package database

import (
	"fmt"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/internal/models"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/shared/crud"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/common/database/paginate"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/common/requests"
	"log"
)

func (s *service) CreateLocation(location *models.Locations) (uint, error) {
	log.Printf("Creating location %s", location.Name)
	id, err := crud.Create(s.db, location)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *service) UpdateLocation(id int, location requests.Update) error {
	log.Printf("Updating location with id %d", id)
	err := crud.Update(s.db, id, location, &models.Locations{})
	if err != nil {
		log.Printf("Failed to update location with id %d: %s", id, err)
		return fmt.Errorf("failed to update location with id %d: %w", id, err)
	}

	return nil
}

func (s *service) GetLocation(id int) (*models.Locations, error) {
	log.Printf("Getting location with id %d", id)
	location, err := crud.Read(s.db, uint(id), &models.Locations{})
	if err != nil {
		return nil, err
	}
	return location, nil
}

func (s *service) GetLocations(page int, amount int) (*paginate.Pagination, error) {
	log.Printf("Getting locations for page %d and amount %d", page, amount)

	paginated := paginate.Pagination{
		Page:  page,
		Limit: amount,
		Sort:  "Id asc",
	}

	locations, err := crud.ReadAll(s.db, paginated, &models.Locations{})
	if err != nil {
		return nil, err
	}

	return locations, nil
}

func (s *service) DeleteLocation(id int) error {
	log.Printf("Deleting location with id %d", id)

	err := crud.Delete(s.db, uint(id), &models.Locations{})
	if err != nil {
		log.Printf("Failed to delete location with id %d: %s", id, err)
		return fmt.Errorf("failed to delete location with id %d: %w", id, err)
	}

	return nil
}
