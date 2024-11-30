package crud

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func Read[T ReadModel](db *gorm.DB, id uint, table T) (T, error) {
	rs := db.Where("id = ?", id).First(table)
	if rs.Error != nil {

		if errors.Is(rs.Error, gorm.ErrRecordNotFound) {
			log.Printf("Element with id %d not found", id)
			return table, fmt.Errorf("location with id %d not found", id)
		}

		return table, rs.Error
	}
	return table, nil
}
