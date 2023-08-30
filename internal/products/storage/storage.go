package storage

import "errors"

// Product is an struct that represents a product in the storage.
type Product struct {
	Id       	int
	Description	string
	Price    	float64
	SellerId 	int
}

// QueryParams is an struct that represents a query to the storage.
type QueryParams struct {
	Id	   	int
}

// StorageProduct is an interface that represents a storage.
type StorageProduct interface {
	// GetProduct returns a product by its id.
	GetProducts(query *QueryParams) (ps []*Product, err error)
}

var (
	// ErrStorageProductInternal is returned when an internal error occurs.
	ErrStorageProductInternal = errors.New("internal error")
)