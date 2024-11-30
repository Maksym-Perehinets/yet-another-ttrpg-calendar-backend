package crud

import (
	"gorm.io/gorm"
	"log"
)

func Create[T CreateModel](db *gorm.DB, table T) (uint, error) {
	rs := db.Create(&table)

	// errors.Is(result.Error, gorm.ErrRecordNotFound) use latter for error wrapping
	if rs.Error != nil {
		log.Printf("Failed to create element: %v", rs.Error)
		return 0, rs.Error
	}

	return table.GetID(), nil
}
