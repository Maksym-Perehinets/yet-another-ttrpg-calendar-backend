package database

import (
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/internal/models"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/shared/crud"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/common/database/paginate"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/common/requests"
	"log"
)

func (s *service) CreateTTRPGSystem(system *models.TTRPGSystems) (uint, error) {
	log.Printf("Creating TTRPG system %s", system.Name)
	id, _ := crud.Create(s.db, system)

	return id, nil
}

func (s *service) UpdateTTRPGSystem(id int, system requests.Update) error {
	log.Printf("Updating TTRPG system with id %d", id)
	err := crud.Update(s.db, id, system, &models.TTRPGSystems{})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetTTRPGSystem(id int) (*models.TTRPGSystems, error) {
	log.Printf("Getting TTRPG system with id %d", id)
	system, err := crud.Read(s.db, uint(id), &models.TTRPGSystems{})
	if err != nil {
		return nil, err
	}

	return system, nil
}

func (s *service) GetTTRPGSystems(page int, amount int) (*paginate.Pagination, error) {
	log.Printf("Getting TTRPG systems for page %d and amount %d", page, amount)
	paginated := paginate.Pagination{
		Page:  page,
		Limit: amount,
		Sort:  "Id asc",
	}

	systems, err := crud.ReadAll(s.db, paginated, &models.TTRPGSystems{})
	if err != nil {
		return nil, err
	}

	return systems, nil
}

func (s *service) DeleteTTRPGSystem(id int) error {
	log.Printf("Deleting TTRPG system with id %d", id)
	err := crud.Delete(s.db, uint(id), &models.TTRPGSystems{})
	if err != nil {
		return err
	}

	return nil
}
