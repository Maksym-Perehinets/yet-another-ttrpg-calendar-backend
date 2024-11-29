package request

type CreateLocation struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	City        string  `json:"city" binding:"required"`
	Street      string  `json:"street" binding:"required"`
	LinkToSite  string  `json:"link_to_site" binding:"required"`
	Price       float64 `json:"price"`
	PricingType string  `json:"pricing_type" binding:"required"`
	OpenAt      string  `json:"open_at" binding:"required"`
	CloseAt     string  `json:"close_at" binding:"required"`
}

type Update []struct {
	Field string `json:"field" binding:"required"`
	Value string `json:"value" binding:"required"`
}

// TODO Rewrite the code bellow to remove boilerplate code

func (u *Update) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	for _, v := range *u {
		result[v.Field] = v.Value
	}
	return result
}
