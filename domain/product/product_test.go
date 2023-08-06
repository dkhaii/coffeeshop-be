package product

import (
	"errors"
	"testing"
)

func TestProduct_NewProduct(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		price       float64
		qty         int
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty name value",
			name:        "",
			price:       6000,
			qty:         4,
			expectedErr: ErrMissingValues,
		},
		{
			test:        "Empty price value",
			name:        "es teh",
			price:       0,
			qty:         2,
			expectedErr: ErrMissingValues,
		},
		{
			test:        "Empty quantity value",
			name:        "Bakso",
			price:       15000,
			qty:         0,
			expectedErr: ErrMissingValues,
		},
		{
			test:        "Empty quantity value",
			name:        "Bakso",
			price:       15000,
			qty:         2,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := NewProduct(tc.name, tc.price, tc.qty)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
