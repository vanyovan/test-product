package entity

// DTO from handler layer to usecase layer

type Product struct {
	ProductID          int64
	ProductName        string  `json:"product_name,omitempty"`
	ProductDescription string  `json:"product_description,omitempty"`
	ProductPrice       float64 `json:"product_price,omitempty"`
	ProductVariety     string  `json:"product_variety,omitempty"`
	ProductRating      float64 `json:"product_rating,omitempty"`
	ProductStock       int     `json:"product_stock,omitempty"`
}
