package response

type LocationResponse struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	City        string         `json:"city"`
	Street      string         `json:"street"`
	LinkToSite  string         `json:"link_to_site"`
	Price       float64        `json:"price"`
	PricingType string         `json:"pricing_type"`
	OpenAt      string         `json:"open_at"`
	CloseAt     string         `json:"close_at"`
	Games       []GameResponse `json:"games,omitempty"`
}
