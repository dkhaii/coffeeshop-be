package customer

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add new customer")
)

type Repository interface {
	GetByID(id uuid.UUID) (Customer, error)
	Add(cust Customer) error
	Update(cust Customer) error
}
