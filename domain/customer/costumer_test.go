package customer

import (
	"errors"
	"testing"
)

func TestCostumer_NewCustomer(t *testing.T) {
	// making the test case
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	// test cases that will be performed
	testCases := []testCase{
		{
			test:        "Empty Name Validation",
			name:        "",
			expectedErr: ErrInvalidPerson,
		},
		{
			test:        "Valid Name",
			name:        "Mordekhai",
			expectedErr: nil,
		},
	}

	// looping the test cases, and perform the test
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := NewCustomer(tc.name)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
