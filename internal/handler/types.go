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
