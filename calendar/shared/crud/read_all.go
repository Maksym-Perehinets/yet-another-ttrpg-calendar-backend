package crud

import (
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/common/database/paginate"
	"gorm.io/gorm"
	"log"
)

// ReadAll reads all elements from the table, return type *paginate.Pagination because it includes all the data,
// as well as pagination information
func ReadAll[T ReadModel](db *gorm.DB, paginated paginate.Pagination, table T) (*paginate.Pagination, error) {
	log.Printf("Getting all elements for page %d and amount %d", paginated.Page, paginated.Limit)
	var elements []T

	db.Scopes(paginate.Paginate(table, &paginated, db)).Find(&elements)

	paginated.Entries = elements

	return &paginated, nil
}
