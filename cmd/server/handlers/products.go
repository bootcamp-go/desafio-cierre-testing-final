package handlers

import (
	"app/internal/products/storage"
	"app/pkg/web/response"
	"net/http"
	"strconv"
)

// NewControllerProduct returns a new instance of ControllerProduct.
func NewControllerProduct(storage storage.StorageProduct) *ControllerProduct {
	return &ControllerProduct{storage: storage}
}

// ControllerProduct is a controller for products returning handlers.
type ControllerProduct struct {
	storage storage.StorageProduct
}

// GetProducts returns a handler for getting products.
type ProductResponseGetProducts struct {
	Id		   	int		`json:"id"`
	Description	string	`json:"description"`
	Price    	float64	`json:"price"`
	SellerId 	int		`json:"seller_id"`
}
type ResponseBodyGetProducts struct {
	Message string						  `json:"message"`
	Data    []*ProductResponseGetProducts `json:"data"`
	Error	bool						  `json:"error"`
}
func (c *ControllerProduct) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// -> get query params
		var query storage.QueryParams
		query.Id, _ = strconv.Atoi(r.URL.Query().Get("id"))

		// -> get products
		products, err := c.storage.GetProducts(&query)
		if err != nil {
			code := http.StatusInternalServerError
			body := &ResponseBodyGetProducts{Message: "internal server error", Data: nil, Error: true}

			response.JSON(w, code, body)
			return
		}

		// response
		code := http.StatusOK
		body := &ResponseBodyGetProducts{Message: "success", Data: make([]*ProductResponseGetProducts, len(products)), Error: false}
		for i, p := range products {
			body.Data[i] = &ProductResponseGetProducts{
				Id: p.Id,
				Description: p.Description,
				Price: p.Price,
				SellerId: p.SellerId,
			}
		}

		response.JSON(w, code, body)
	}
}