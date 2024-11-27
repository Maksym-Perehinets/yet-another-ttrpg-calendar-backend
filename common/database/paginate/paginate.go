package paginate

import (
	"gorm.io/gorm"
	"math"
)

type Pagination struct {
	Page      int         `json:"page"`
	Limit     int         `json:"limit"`
	Sort      string      `json:"sort"`
	TotalRows int         `json:"total_rows"`
	TotalPage int         `json:"total_page"`
	Entries   interface{} `json:"entries"`
}

func (p *Pagination) GetOffset() int {
	return (p.Page - 1) * p.Limit
}

func (p *Pagination) GetLimit() int {
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		return 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}
func Paginate(table interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(table).Count(&totalRows)

	pagination.TotalRows = int(totalRows)
	pagination.TotalPage = int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
