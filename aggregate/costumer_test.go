package aggregate_test

import (
	"errors"
	"testing"

	"github.com/dkhaii/cofeeshop-be/aggregate"
)

func TestCostumer_NewCostumer(t *testing.T) {
	type testCase struct {
		test string
		name string
		expectedErr error
	}

	testCases := []testCase{
		{
			test: "Empty Name Validation",
			name: "",
			expectedErr: aggregate.ErrInvalidPerson,
		},
		{
			test: "Valid Name",
			name: "Mordekhai",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases{
		t.Run(tc.name, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}