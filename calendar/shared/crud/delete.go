package crud

import (
	"fmt"
	"gorm.io/gorm"
)

func Delete[T DeleteModel](db *gorm.DB, id uint, table T) error {
	rs := db.Where("id = ?", id).Delete(&table)
	if rs.Error != nil {
		return fmt.Errorf("failed to delete element with id %d: %w", id, rs.Error)
	}

	if rs.RowsAffected == 0 {
		return fmt.Errorf("element with id %d not found", id)
	}

	return nil
}
