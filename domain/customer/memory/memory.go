package memory

import (
	"fmt"
	"sync"

	"github.com/dkhaii/cofeeshop-be/domain/customer"
	"github.com/google/uuid"
)

type MemoryRepository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

// factory
func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

func (mry *MemoryRepository) GetByID(id uuid.UUID) (customer.Customer, error) {
	if customer, isExist := mry.customers[id]; isExist {
		return customer, nil
	}

	return customer.Customer{}, customer.ErrCustomerNotFound
}

func (mry *MemoryRepository) Add(cust customer.Customer) error {
	if mry.customers == nil {
		mry.Lock()
		defer mry.Unlock()

		mry.customers = make(map[uuid.UUID]customer.Customer)
	}

	// check if customer is already in repo
	if _, isExist := mry.customers[cust.GetID()]; isExist {
		return fmt.Errorf("customer already exist: %w", customer.ErrFailedToAddCustomer)
	}

	mry.customers[cust.GetID()] = cust

	return nil
}

func (mry *MemoryRepository) Update(cust customer.Customer) error {
	// check if customer doesnt exist in repo
	if _, isExist := mry.customers[cust.GetID()]; !isExist {
		return fmt.Errorf("customer doesnt exist: %w", customer.ErrCustomerNotFound)
	}

	mry.Lock()
	defer mry.Unlock()

	mry.customers[cust.GetID()] = cust

	return nil
}
