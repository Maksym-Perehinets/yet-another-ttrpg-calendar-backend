package database

import (
	"errors"
	"fmt"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/internal/models"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/shared/request"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/common/database/paginate"
	"gorm.io/gorm"
	"log"
)

func (s *service) CreateLocation(location *models.Locations) (uint, error) {
	rs := s.db.Create(&location)

	// errors.Is(result.Error, gorm.ErrRecordNotFound) use latter for error wrapping
	if rs.Error != nil {
		return 0, rs.Error
	}
	return location.ID, nil
}

func (s *service) GetLocations(page int, amount int) (*paginate.Pagination, error) {
	log.Printf("Getting locations for page %d and amount %d", page, amount)
	var locations []*models.Locations

	paginated := paginate.Pagination{
		Page:  page,
		Limit: amount,
		Sort:  "Id asc",
	}

	s.db.Scopes(paginate.Paginate(locations, &paginated, s.DB())).Find(&locations)

	paginated.Entries = locations

	return &paginated, nil
}

func (s *service) DeleteLocation(id int) error {
	rs := s.db.Where("id = ?", id).Delete(&models.Locations{})
	if rs.Error != nil {
		return fmt.Errorf("failed to delete location with id %d: %w", id, rs.Error)
	}

	if rs.RowsAffected == 0 {
		return fmt.Errorf("location with id %d not found", id)
	}

	return nil
}

func (s *service) GetLocation(id int) (*models.Locations, error) {
	var location models.Locations
	rs := s.db.Where("id = ?", id).First(&location)
	if rs.Error != nil {
		if errors.Is(rs.Error, gorm.ErrRecordNotFound) {
			log.Printf("Location with id %d not found", id)
			return nil, fmt.Errorf("location with id %d not found", id)
		}
		return &models.Locations{}, rs.Error
	}
	return &location, nil

}

func (s *service) UpdateLocation(id int, location request.Update) error {
	rs := s.db.Model(&models.Locations{}).Where("id = ?", id).Omit("id").Updates(location.ToMap())
	if rs.Error != nil {
		return fmt.Errorf("failed to update location with id %d: %w", id, rs.Error)
	}

	if rs.RowsAffected == 0 {
		return fmt.Errorf("location with id %d not found", id)
	}

	return nil
}
