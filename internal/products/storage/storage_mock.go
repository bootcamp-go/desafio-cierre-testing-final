package storage

// NewStorageProductMock returns a new instance of StorageProductMock.
func NewStorageProductMock() *StorageProductMock {
	return &StorageProductMock{}
}

// StorageProductMock is a mock implementation of StorageProduct.
type StorageProductMock struct {
	// FuncGetProducts is the function that is called when GetProducts is called.
	FuncGetProducts func(query *QueryParams) (ps []*Product, err error)

	// Calls
	Calls struct {
		// GetProducts is the number of times GetProducts has been called.
		GetProducts int
	}
}

// GetProducts returns a product by its id.
func (s *StorageProductMock) GetProducts(query *QueryParams) (ps []*Product, err error) {
	ps, err = s.FuncGetProducts(query)
	return
}