package dto

type SortProductRequest struct {
	Products []Product `json:"product"`
	Strategy string    `json:"strategy"`
}

type Product struct {
	Name    string
	ID      int
	Price   float64
	Created string
	Sales   int
	Views   int
}
