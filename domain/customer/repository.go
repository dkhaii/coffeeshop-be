package costumer

import (
	"errors"

	"github.com/dkhaii/cofeeshop-be/aggregate"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound       = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer    = errors.New("failed to add new customer")
)

type CostumerRepository interface {
	Get(id uuid.UUID) (aggregate.Customer, error)
	Add(cust aggregate.Customer) error
	Update(cust aggregate.Customer) error
}
