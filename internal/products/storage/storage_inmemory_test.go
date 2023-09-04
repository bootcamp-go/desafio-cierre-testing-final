package storage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for StorageProductInMemory.
func TestStorageProductInMemory_GetProducts(t *testing.T) {
	type input struct { query *QueryParams }
	type output struct { ps []*Product; err error; errMsg string }
	type test struct {
		name string
		input input
		output output
		// set-up
		setUpDataba func(db map[int]*Product)
	}

	// cases
	cases := []test{
		// case empty database
		{
			name: "empty database",
			input: input{ query: &QueryParams{} },
			output: output{ ps: []*Product{}, err: nil },
			setUpDataba: func(db map[int]*Product) {},
		},

		// case query nil
		{
			name: "query nil",
			input: input{ query: nil },
			output: output{ ps: []*Product{
				{ Id: 1, Description: "product 1", Price: 1.1, SellerId: 1 },
				{ Id: 2, Description: "product 2", Price: 2.2, SellerId: 2 },
				{ Id: 3, Description: "product 3", Price: 3.3, SellerId: 3 },
			}, err: nil },
			setUpDataba: func(db map[int]*Product) {
				db[1] = &Product{ Id: 1, Description: "product 1", Price: 1.1, SellerId: 1 }
				db[2] = &Product{ Id: 2, Description: "product 2", Price: 2.2, SellerId: 2 }
				db[3] = &Product{ Id: 3, Description: "product 3", Price: 3.3, SellerId: 3 }
			},
		},

		// case query none
		{
			name: "query none",
			input: input{ query: &QueryParams{} },
			output: output{ ps: []*Product{
				{ Id: 1, Description: "product 1", Price: 1.1, SellerId: 1 },
				{ Id: 2, Description: "product 2", Price: 2.2, SellerId: 2 },
				{ Id: 3, Description: "product 3", Price: 3.3, SellerId: 3 },
			}, err: nil },
			setUpDataba: func(db map[int]*Product) {
				db[1] = &Product{ Id: 1, Description: "product 1", Price: 1.1, SellerId: 1 }
				db[2] = &Product{ Id: 2, Description: "product 2", Price: 2.2, SellerId: 2 }
				db[3] = &Product{ Id: 3, Description: "product 3", Price: 3.3, SellerId: 3 }
			},
		},


		// case query id
		{
			name: "query id",
			input: input{ query: &QueryParams{ Id: 2 } },
			output: output{ ps: []*Product{
				{ Id: 2, Description: "product 2", Price: 2.2, SellerId: 2 },
			}, err: nil },
			setUpDataba: func(db map[int]*Product) {
				db[1] = &Product{ Id: 1, Description: "product 1", Price: 1.1, SellerId: 1 }
				db[2] = &Product{ Id: 2, Description: "product 2", Price: 2.2, SellerId: 2 }
				db[3] = &Product{ Id: 3, Description: "product 3", Price: 3.3, SellerId: 3 }
			},
		},
	}

	// run tests
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// arrange
			// -> database
			db := make(map[int]*Product)
			c.setUpDataba(db)

			// -> storage
			st := NewStorageProductInMemory(db)

			// act
			ps, err := st.GetProducts(c.input.query)

			// assert
			t.Log("db", db)
			require.Equal(t, c.output.ps, ps)
			require.ErrorIs(t, c.output.err, err)
			if c.output.err != nil {
				require.Equal(t, c.output.errMsg, err.Error())
			}
		})
	}
}