package handler

// to store all request
type CreateProductRequest struct {
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	ProductPrice       float64 `json:"product_price"`
	ProductVariety     string  `json:"product_variety"`
	ProductRating      float64 `json:"product_rating"`
	ProductStock       int     `json:"product_stock"`
}

type UpdateProductRequest struct {
	ProductName        string  `json:"product_name,omitempty"`
	ProductDescription string  `json:"product_description,omitempty"`
	ProductPrice       float64 `json:"product_price,omitempty"`
	ProductVariety     string  `json:"product_variety,omitempty"`
	ProductRating      float64 `json:"product_rating,omitempty"`
	ProductStock       int     `json:"product_stock,omitempty"`
}
