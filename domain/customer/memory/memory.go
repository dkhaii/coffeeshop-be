package memory

import (
	"fmt"
	"sync"

	"github.com/dkhaii/cofeeshop-be/aggregate"
	customer "github.com/dkhaii/cofeeshop-be/domain/customer"
	"github.com/google/uuid"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mry *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, isExist := mry.customers[id]; isExist {
		return customer, nil
	}

	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (mry *MemoryRepository) Add(c aggregate.Customer) error {
	if mry.customers == nil {
		mry.Lock()
		mry.customers = make(map[uuid.UUID]aggregate.Customer)
		mry.Unlock()
	}

	// check if customer is already in repo
	if _, isExist := mry.customers[c.GetID()]; isExist {
		return fmt.Errorf("customer already exist: %w", customer.ErrFailedToAddCustomer)
	}

	mry.Lock()
	mry.customers[c.GetID()] = c
	mry.Unlock()

	return nil
}

func (mry *MemoryRepository) Update(c aggregate.Customer) error {
	// check if customer doesnt exist in repo
	if _, isExist := mry.customers[c.GetID()]; !isExist {
		return fmt.Errorf("customer doesnt exist: %w", customer.ErrFailedToUpdateCustomer)
	}

	mry.Lock()
	mry.customers[c.GetID()] = c
	mry.Unlock()

	return nil
}
