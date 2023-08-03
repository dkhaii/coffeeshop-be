package costumer

import (
	"errors"

	"github.com/dkhaii/cofeeshop-be/aggregate"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound       = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer    = errors.New("failed to add new customer")
	ErrFailedToUpdateCustomer = errors.New("failed to update the customer")
)

type CostumerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
