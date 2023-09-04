package storage

// NewStorageProductInMemory returns a new instance of StorageProductInMemory.
func NewStorageProductInMemory(db map[int]*Product) *StorageProductInMemory {
	return &StorageProductInMemory{db: db}
}

// StorageProductInMemory is an in-memory implementation of StorageProduct.
type StorageProductInMemory struct {
	db map[int]*Product
}

// GetProducts returns a product by its id.
func (s *StorageProductInMemory) GetProducts(query *QueryParams) (ps []*Product, err error) {
	ps = []*Product{}

	// load and filter products
	for _, p := range s.db {
		// check if filter is set
		if query != nil && query.Id > 0 {
			// filter by id
			if query.Id != p.Id {
				continue
			}

			// others
			// ...
		}

		// add product
		ps = append(ps, p)
	}

	return
}