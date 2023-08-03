package aggregate_test

import (
	"errors"
	"testing"

	"github.com/dkhaii/cofeeshop-be/aggregate"
)

func TestCostumer_NewCustomer(t *testing.T) {
	// making the test case
	type testCase struct {
		test string
		name string
		expectedErr error
	}

	// test cases that will be performed
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

	// looping the test cases, and perform the test
	for _, tc := range testCases{
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}