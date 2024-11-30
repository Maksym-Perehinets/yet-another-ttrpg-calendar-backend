package crud

import (
	"fmt"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/common/requests"
	"gorm.io/gorm"
)

func Update[T UpdateModel](db *gorm.DB, id int, request requests.Update, table T) error {
	rs := db.Model(&table).Where("id = ?", id).Omit("id").Updates(request.ToMap())
	if rs.Error != nil {
		return fmt.Errorf("failed to update v with id %d: %w", id, rs.Error)
	}

	if rs.RowsAffected == 0 {
		return fmt.Errorf("element with id %d not found", id)
	}

	return nil
}
