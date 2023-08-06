package memory

import (
	"errors"
	"testing"

	"github.com/dkhaii/cofeeshop-be/domain/product"
	"github.com/google/uuid"
)

func TestMemory_GetProduct(t *testing.T) {
	type testCase struct {
		test        string
		id          uuid.UUID
		expectedErr error
	}

	esTehJumbo, err := product.NewProduct("Es Teh Jumbo", 9.000, 2)
	if err != nil {
		t.Fatal(err)
	}

	id := esTehJumbo.GetID()

	repo := MemoryRepository{
		products: map[uuid.UUID]product.Product{
			id: esTehJumbo,
		},
	}

	testCases := []testCase{
		{
			test:        "no product id",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: product.ErrProductNotFound,
		},
		{
			test:        "product id exist",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := repo.GetByID(tc.id)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
