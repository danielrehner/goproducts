package dto

// Product is a struct to hold information about products
type Product struct {
	ID    string  `json:"id" example:"1"`
	Title string  `json:"title" example:"Cookies"`
	Price float64 `json:"price" example:"8.27"`
}

// ProductResponse is an API struct to hold and return a product.
type ProductResponse struct {
	Data Product `json:"data"`
}

// ProductSearchResult is an API struct to hold and return search results.
type ProductSearchResult struct {
	Data []Product `json:"data"`
}
