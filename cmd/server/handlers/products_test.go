package handlers

import (
	"app/internal/products/storage"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for ControllerProduct
func TestControllerProduct_GetProducts(t *testing.T) {
	type input struct { setUpR func(r *http.Request); rr *httptest.ResponseRecorder }
	type output struct { code int; body string }
	type test struct {
		name string
		input input
		output output
		// set-up
		setUpDatabase func(db map[int]*storage.Product)
	}

	// test cases
	cases := []test{
		// success
		// -> no products - query non-set
		{
			name: "success: no products - query non-set",
			input: input{
				setUpR: func(r *http.Request) {},
				rr: httptest.NewRecorder(),
			},
			output: output{
				code: http.StatusOK,
				body: `{"message":"success","data":[],"error":false}`,
			},
			setUpDatabase: func(db map[int]*storage.Product) {},
		},
		// -> no products - query set
		{
			name: "success: no products - query set",
			input: input{
				setUpR: func(r *http.Request) {
					// add query params
					query := r.URL.Query()
					query.Add("id", "1")

					(*r).URL.RawQuery = query.Encode()
				},
				rr: httptest.NewRecorder(),
			},
			output: output{
				code: http.StatusOK,
				body: `{"message":"success","data":[],"error":false}`,
			},
			setUpDatabase: func(db map[int]*storage.Product) {},
		},
		// -> one product - query non-set
		{
			name: "success: one product - query non-set",
			input: input{
				setUpR: func(r *http.Request) {},
				rr: httptest.NewRecorder(),
			},
			output: output{
				code: http.StatusOK,
				body: `{"message":"success","data":[{"id":1,"description":"product 1","price":1.1,"seller_id":1}],"error":false}`,
			},
			setUpDatabase: func(db map[int]*storage.Product) {
				db[1] = &storage.Product{
					Id: 1,
					Description: "product 1",
					Price: 1.1,
					SellerId: 1,
				}
			},
		},
		// -> one product - query set
		{
			name: "success: one product - query set",
			input: input{
				setUpR: func(r *http.Request) {
					// add query params
					query := r.URL.Query()
					query.Add("id", "1")

					(*r).URL.RawQuery = query.Encode()
				},
				rr: httptest.NewRecorder(),
			},
			output: output{
				code: http.StatusOK,
				body: `{"message":"success","data":[{"id":1,"description":"product 1","price":1.1,"seller_id":1}],"error":false}`,
			},
			setUpDatabase: func(db map[int]*storage.Product) {
				db[1] = &storage.Product{
					Id: 1,
					Description: "product 1",
					Price: 1.1,
					SellerId: 1,
				}
				db[2] = &storage.Product{
					Id: 2,
					Description: "product 2",
					Price: 2.2,
					SellerId: 2,
				}
			},
		},
	}

	// run tests
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// arrange
			// -> database
			db := make(map[int]*storage.Product)
			c.setUpDatabase(db)

			// -> storage
			st := storage.NewStorageProductInMemory(db)

			// -> controller
			ct := NewControllerProduct(st)
			hd := ct.GetProducts()

			// act
			r := httptest.NewRequest(http.MethodGet, "/products", nil); c.input.setUpR(r)
			hd(c.input.rr, r)

			// assert
			require.Equal(t, c.output.code, c.input.rr.Code)
			require.JSONEq(t, c.output.body, c.input.rr.Body.String())
		})
	}
}