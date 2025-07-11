package dto

type SortProductRequest struct {
	Products []Product `json:"products"`
	SortBy   string    `json:"sort_by"`
}

type Product struct {
	Name    string  `json:"name"`
	ID      int     `json:"id"`
	Price   float64 `json:"price"`
	Created string  `json:"created"`
	Sales   int     `json:"sales"`
	Views   int     `json:"views"`
}
