package storage

// NewStorageProductInMemory returns a new instance of StorageProductInMemory.
func NewStorageProductInMemory() *StorageProductInMemory {
	return &StorageProductInMemory{products: make(map[int]*Product)}
}

// StorageProductInMemory is an in-memory implementation of StorageProduct.
type StorageProductInMemory struct {
	products map[int]*Product
}

// GetProducts returns a product by its id.
func (s *StorageProductInMemory) GetProducts(query *QueryParams) (ps []*Product, err error) {
	for _, p := range s.products {
		// check is query is not none and filter by id
		if query.Id != 0 {
			if p.Id == query.Id {
				ps = append(ps, p)
			}
		}
	}

	return
}